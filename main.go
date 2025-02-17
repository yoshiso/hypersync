package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli/v3"
	"github.com/yoshiso/hypersync/ent"
	_ "modernc.org/sqlite"
)

type WSRequest struct {
	Method       string           `json:"method"`
	Subscription UserSubscription `json:"subscription"`
}

type UserSubscription struct {
	Type string `json:"type"`
	User string `json:"user"`
}

type FillLiquidation struct {
	LiquidatedUser *string `json:"liquidatedUser"`
	MarkPx         string  `json:"markPx"`
	Method         string  `json:"method"`
}

type WsUserFunding struct {
	Time        int64  `json:"time"`
	Coin        string `json:"coin"`
	Usdc        string `json:"usdc"`
	Szi         string `json:"szi"`
	FundingRate string `json:"fundingRate"`
}

type WsFill struct {
	Coin          string           `json:"coin"`
	Px            string           `json:"px"`
	Sz            string           `json:"sz"`
	Side          string           `json:"side"`
	Time          int64            `json:"time"`
	StartPosition string           `json:"startPosition"`
	Dir           string           `json:"dir"`
	ClosedPnl     string           `json:"closedPnl"`
	Hash          string           `json:"hash"`
	Oid           int64            `json:"oid"`
	Crossed       bool             `json:"crossed"`
	Fee           string           `json:"fee"`
	Tid           int64            `json:"tid"`
	Liquidation   *FillLiquidation `json:"liquidation"`
	FeeToken      string           `json:"feeToken"`
	BuilderFee    *string          `json:"builderFee"`
}

type WsUserFundings struct {
	IsSnapshot bool            `json:"isSnapshot"`
	User       string          `json:"user"`
	Fundings   []WsUserFunding `json:"fundings"`
}

type WsUserFills struct {
	IsSnapshot bool     `json:"isSnapshot"`
	User       string   `json:"user"`
	Fills      []WsFill `json:"fills"`
}

type WsUserFillsResponse struct {
	Channel string      `json:"channel"`
	Data    WsUserFills `json:"data"`
}

type WsUserFundingsResponse struct {
	Channel string         `json:"channel"`
	Data    WsUserFundings `json:"data"`
}

type HTTPUserFundingDelta struct {
	Coin        string `json:"coin"`
	Usdc        string `json:"usdc"`
	Szi         string `json:"szi"`
	FundingRate string `json:"fundingRate"`
}

type HTTPUserFunding struct {
	Time  int64                `json:"time"`
	Hash  string               `json:"hash"`
	Delta HTTPUserFundingDelta `json:"delta"`
}

type WSUserNonFundingLedgerUpdate struct {
	Time  int64                             `json:"time"`
	Hash  string                            `json:"hash"`
	Delta WSUserNonFundingLedgerUpdateDelta `json:"delta"`
}

type WSUserNonFundingLedgerUpdateDelta struct {
	Type            string `json:"type"`
	Amount          string `json:"amount"`
	Token           string `json: "token"`
	Usdc            string `json:"usdc"`
	User            string `json:"user"`
	Vault           string `json:"vault"`
	RequestedUsd    string `json:"requestedUsd"`
	Commission      string `json:"commission"`
	ClosingCost     string `json:"closingCost"`
	Basis           string `json:"basis"`
	NetWithdrawnUsd string `json:"netWithdrawnUsd"`
	Destination     string `json:"destination"`
	Fee             string `json:"fee"`
	Nonce           int64  `json:"nonce"`
}

type WSUserNonFundingLedgerUpdates struct {
	IsSnapshot              bool                           `json:"isSnapshot"`
	User                    string                         `json:"user"`
	NonFundingLedgerUpdates []WSUserNonFundingLedgerUpdate `json:"nonFundingLedgerUpdates"`
}

type WSUserNonFundingLedgerUpdatesResponse struct {
	Channel string                        `json:"channel"`
	Data    WSUserNonFundingLedgerUpdates `json:"data"`
}

type DelegatorHistoryDelta struct {
	Delegate *DelegatorHistoryDelegateDelta `json:"delegate"`
}

type DelegatorHistoryDelegateDelta struct {
	Validator    string `json:"validator"`
	Amount       string `json:"amount"`
	IsUndelegate bool   `json:"isUndelegate"`
}

type DelegatorHisotry struct {
	Time  int64                 `json:"time"`
	Hash  string                `json:"hash"`
	Delta DelegatorHistoryDelta `json:"delta"`
}

type DelegatorReward struct {
	Time        int64  `json:"time"`
	Source      string `json:"source"`
	TotalAmount string `json:"totalAmount"`
}

type UserTwapSliceFill struct {
	Fill   WsFill `json:"fill"`
	TwapID int64  `json:"twapId"`
}

type WSUserTwapSliceFills struct {
	IsSnapshot     bool                `json:"isSnapshot"`
	User           string              `json:"user"`
	TwapSliceFills []UserTwapSliceFill `json:"twapSliceFills"`
}
type WSUserTwapSliceFillsResponse struct {
	Channel string               `json:"channel"`
	Data    WSUserTwapSliceFills `json:"data"`
}

type HyperunitOperation struct {
    OpCreatedAt         time.Time `json:"opCreatedAt"`
    OperationID         string    `json:"operationId"`
    ProtocolAddress     string    `json:"protocolAddress"`
    SourceAddress       string    `json:"sourceAddress"`
    DestinationAddress  string    `json:"destinationAddress"`
    SourceChain         string    `json:"sourceChain"`
    DestinationChain    string    `json:"destinationChain"`
    SourceAmount        string    `json:"sourceAmount"`
    DestinationFeeAmount string   `json:"destinationFeeAmount"`
    SweepFeeAmount      string    `json:"sweepFeeAmount"`
    StateStartedAt      time.Time `json:"stateStartedAt"`
    StateUpdatedAt      time.Time `json:"stateUpdatedAt"`
    StateNextAttemptAt  time.Time `json:"stateNextAttemptAt"`
    SourceTxHash        string    `json:"sourceTxHash"`
    DestinationTxHash   string    `json:"destinationTxHash"`
    BroadcastAt         time.Time `json:"broadcastAt"`
    State               string    `json:"state"`
}

type HyperunitOperationResponse struct {
	Channel string                        `json:"channel"`
	Operations    []HyperunitOperation `json:"operations"`
	Address string
}


type Msg struct {
	Channel string `json:"channel"`
}

type Ping struct {
	Method string `json:"method"`
}

func Conn(userAddress string) chan []byte {
	address := "wss://api.hyperliquid.xyz/ws"

	var c *websocket.Conn
	var err error

	for {
		c, _, err = websocket.DefaultDialer.Dial(address, nil)

		if err != nil {
			fmt.Println("Causing error and trying to re-connect websocket...")
			time.Sleep(time.Second * 10)
			continue
		}

		break
	}

	var msgC = make(chan []byte)

	go func() {
		for {
			_, msg, err := c.ReadMessage()
			// When error happened then forsibly reconnect socket and notify error.
			if err != nil {
				// TODO: what should I do to notify?
				fmt.Println(fmt.Errorf("ws.ReadMessage error %v", err))
				close(msgC)
				break
			}
			msgC <- msg
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 50)

		for {
			<-ticker.C
			if err := c.WriteJSON(Ping{Method: "ping"}); err != nil {
				fmt.Println(fmt.Errorf("ws.WriteJSON ping error %v", err))
				break
			}
		}
	}()

	if err := c.WriteJSON(WSRequest{Method: "subscribe", Subscription: UserSubscription{User: userAddress, Type: "userFills"}}); err != nil {
		fmt.Println(fmt.Errorf("ws.WriteJSON error %v", err))
	}

	if err := c.WriteJSON(WSRequest{Method: "subscribe", Subscription: UserSubscription{User: userAddress, Type: "userFundings"}}); err != nil {
		fmt.Println(fmt.Errorf("ws.WriteJSON error %v", err))
	}

	if err := c.WriteJSON(WSRequest{Method: "subscribe", Subscription: UserSubscription{User: userAddress, Type: "userNonFundingLedgerUpdates"}}); err != nil {
		fmt.Println(fmt.Errorf("ws.WriteJSON error %v", err))
	}

	if err := c.WriteJSON(WSRequest{Method: "subscribe", Subscription: UserSubscription{User: userAddress, Type: "userTwapSliceFills"}}); err != nil {
		fmt.Println(fmt.Errorf("ws.WriteJSON error %v", err))
	}

	return msgC
}

type HTTPUserRequest struct {
	User string `json:"user"`
	Type string `json:"type"`
}

func fetchUserFills(userAddress string) ([]WsFill, error) {
	jsonBody, _ := json.Marshal(HTTPUserRequest{User: userAddress, Type: "userFills"})

	res, err := http.Post("https://api.hyperliquid.xyz/info", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fills := []WsFill{}

	json.Unmarshal(body, &fills)

	return fills, nil
}

func fetchUserFundings(userAddress string) ([]WsUserFunding, error) {
	jsonBody, _ := json.Marshal(HTTPUserRequest{User: userAddress, Type: "userFunding"})

	res, err := http.Post("https://api.hyperliquid.xyz/info", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	httpFundings := []HTTPUserFunding{}

	json.Unmarshal(body, &httpFundings)

	fundings := []WsUserFunding{}

	for _, funding := range httpFundings {
		fundings = append(fundings, WsUserFunding{
			Coin:        funding.Delta.Coin,
			Time:        funding.Time,
			Usdc:        funding.Delta.Usdc,
			Szi:         funding.Delta.Szi,
			FundingRate: funding.Delta.FundingRate,
		})
	}

	return fundings, nil
}

func fetchUserNonFundingLedgerUpdates(userAddress string) ([]WSUserNonFundingLedgerUpdate, error) {
	jsonBody, _ := json.Marshal(HTTPUserRequest{User: userAddress, Type: "userNonFundingLedgerUpdates"})

	res, err := http.Post("https://api.hyperliquid.xyz/info", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	updates := []WSUserNonFundingLedgerUpdate{}

	json.Unmarshal(body, &updates)

	return updates, nil
}

func fetchDelegatorHistory(userAddress string) ([]DelegatorHisotry, error) {
	jsonBody, _ := json.Marshal(HTTPUserRequest{User: userAddress, Type: "delegatorHistory"})

	res, err := http.Post("https://api.hyperliquid.xyz/info", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	updates := []DelegatorHisotry{}

	json.Unmarshal(body, &updates)

	return updates, nil
}

func fetchDelegatorRewards(userAddress string) ([]DelegatorReward, error) {
	jsonBody, _ := json.Marshal(HTTPUserRequest{User: userAddress, Type: "delegatorRewards"})

	res, err := http.Post("https://api.hyperliquid.xyz/info", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	updates := []DelegatorReward{}

	json.Unmarshal(body, &updates)

	return updates, nil
}

func fetchUserTwapSliceFills(userAddress string) ([]UserTwapSliceFill, error) {
	jsonBody, _ := json.Marshal(HTTPUserRequest{User: userAddress, Type: "userTwapSliceFills"})

	res, err := http.Post("https://api.hyperliquid.xyz/info", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	updates := []UserTwapSliceFill{}

	json.Unmarshal(body, &updates)

	return updates, nil
}


func fetchHyperunitOperations(userAddress string) ([]HyperunitOperation, error) {
	res, err := http.Get("https://api.hyperunit.xyz/operations/" + userAddress)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	updates := HyperunitOperationResponse{}

	json.Unmarshal(body, &updates)

	return updates.Operations, nil
}


func uploadToS3(filePath string, bucket string, key string, awsRegion string) error {
	imageFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(newSession, &aws.Config{
		Region: aws.String(awsRegion),
	})

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   imageFile,
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return err
	}

	return nil
}

func runBackup(db_path string, backupPath string, awsS3Region string) {

	if backupPath == "" {
		return
	}

	if backupPath[:5] == "s3://" {

		s3FullPath := backupPath[5:]

		paths := strings.SplitN(s3FullPath, "/", 2)

		if len(paths) != 2 {
			fmt.Println(fmt.Sprintf("Invalid backup path format: ", backupPath))
			return
		}

		if err := uploadToS3(db_path, paths[0], paths[1], awsS3Region); err != nil {
			fmt.Println(fmt.Sprintf("Failed to upload to S3: ", err))
		}

		fmt.Println(fmt.Sprintf("[Backup] Successfully bacaked up to %s", backupPath))

		return
	}

	if backupPath[:7] == "file://" {

		filePath := backupPath[7:]

		copyFileContents(db_path, filePath)

		fmt.Println(fmt.Sprintf("[Backup] Successfully bacaked up to %s", backupPath))

		return
	}
	fmt.Println("Unknown backupPath = %s", backupPath)
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
				Value: "",
				Usage: "hyperliquid wallet address to sync fills",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Value: false,
				Usage: "verbose.",
			},
			&cli.StringFlag{
				Name:  "out",
				Value: "db.sqlite3",
				Usage: "database file output destination",
			},
			&cli.StringFlag{
				Name:  "backup",
				Value: "",
				Usage: "backup file output destination",
			},
			&cli.IntFlag{
				Name:  "backup-interval-seconds",
				Value: 60 * 60 * 8,
				Usage: "seconds interval to run backup",
			},
			&cli.StringFlag{
				Name:  "aws-s3-region",
				Value: "ap-northeast-1",
				Usage: "aws s3 region used to store backup file",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			userAddress := cmd.String("address")

			if userAddress == "" {
				fmt.Println("Add --address flag with your wallet")
				os.Exit(1)
			}

			verbose := cmd.Bool("verbose")

			outputFile := cmd.String("out")

			backup := cmd.String("backup")

			awsS3Region := cmd.String("aws-s3-region")

			backupIntervalSeconds := cmd.Int("backup-interval-seconds")

			run(userAddress, verbose, outputFile, backup, awsS3Region, backupIntervalSeconds)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runFetchAndStoreDelegatorHistory(userAddress string, client *ent.Client, ctx context.Context, verbose bool) error {
	fmt.Println("Loading delegatorHisotry via REST API")
	initDelegates, err := fetchDelegatorHistory(userAddress)

	if err != nil {
		fmt.Println(fmt.Sprintf("failed to load delegatorHistory via API, and retrying... (%v)"))
		return err
	}

	for _, item := range initDelegates {
		if item.Delta.Delegate != nil {
			delegate := client.Delegate.Create().
				SetAddress(userAddress).
				SetTime(item.Time).
				SetAmount(item.Delta.Delegate.Amount).
				SetValidator(item.Delta.Delegate.Validator).
				SetIsUndelegate(item.Delta.Delegate.IsUndelegate)

			delegate.OnConflict().UpdateNewValues().IDX(ctx)

			if verbose {
				fmt.Println(fmt.Sprintf(
					"[Delegate] Validator: %s, Amount: %s, IsUndelegate: %v",
					item.Delta.Delegate.Validator, item.Delta.Delegate.Amount, item.Delta.Delegate.IsUndelegate,
				))
			}
		}
	}

	fmt.Println("Complete delegatorHisotry via REST API")

	return nil
}

func runFetchAndStoreDelegatorReward(userAddress string, client *ent.Client, ctx context.Context, verbose bool) error {
	fmt.Println("Loading delegatorReward via REST API")
	initRewards, err := fetchDelegatorRewards(userAddress)

	if err != nil {
		fmt.Println(fmt.Sprintf("failed to load delegatorReward via API, and retrying... (%v)"))
		return err
	}

	for _, item := range initRewards {
		reward := client.DelegatorReward.Create().
			SetAddress(userAddress).
			SetTime(item.Time).
			SetTotalAmount(item.TotalAmount).
			SetSource(item.Source)

		reward.OnConflict().UpdateNewValues().IDX(ctx)

		if verbose {
			fmt.Println(fmt.Sprintf(
				"[DelegatorReward] Source: %s, TotalAmount: %s",
				item.Source, item.TotalAmount,
			))
		}
	}

	fmt.Println("Complete delegatorReward via REST API")

	return nil
}


func runFetchAndStoreHyperunitOperation(userAddress string, client *ent.Client, ctx context.Context, verbose bool) error {
	fmt.Println("Loading hyperunitOperation via REST API")
	initOps, err := fetchHyperunitOperations(userAddress)

	if err != nil {
		fmt.Println(fmt.Sprintf("failed to load hyperunitOperation via API, and retrying... (%v)"))
		return err
	}

	for _, item := range initOps {
		op := client.HyperunitOperation.Create().
			SetAddress(userAddress).
			SetOperationID(item.OperationID).
			SetSourceChain(item.SourceChain).
			SetSourceAmount(item.SourceAmount).
			SetSourceAddress(item.SourceAddress).
			SetSourceTxHash(item.SourceTxHash).
			SetDestinationTxHash(item.DestinationTxHash).
			SetDestinationFeeAmount(item.DestinationFeeAmount).
			SetDestinationChain(item.DestinationChain).
			SetDestinationAddress(item.DestinationAddress).
			SetSweepFeeAmount(item.SweepFeeAmount).
			SetOpCreatedAt(item.OpCreatedAt).
			SetBroadcastAt(item.BroadcastAt).
			SetStateUpdatedAt(item.StateUpdatedAt)

			op.OnConflict().UpdateNewValues().IDX(ctx)

		if verbose {
			fmt.Println(fmt.Sprintf(
				"[HyperunitOperation] SourceAmount: %s, DestinationFeeAmount: %s, SweepFeeAmount: %v, StateUpdatedAt: %v",
				item.SourceAmount, item.DestinationFeeAmount, item.SweepFeeAmount, item.StateUpdatedAt,
			))
		}
	}

	fmt.Println("Complete delegatorHisotry via REST API")

	return nil
}


func run(userAddress string, verbose bool, outputFile string, backupFile string, awsS3Region string, backupIntervalSeconds int64) {

	fmt.Println("start running server", userAddress, outputFile)

	client, err := ent.Open("sqlite3", "file:"+outputFile+"?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ticker := time.NewTicker(time.Second * time.Duration(backupIntervalSeconds))

	periodicSyncTicker := time.NewTicker(time.Second * time.Duration(60*60*4))

	for {
		msgC := Conn(userAddress)

		ctx := context.TODO()

		if err := runFetchAndStoreDelegatorHistory(userAddress, client, ctx, verbose); err != nil {
			time.Sleep(time.Second * 10)
			continue
		}

		if err := runFetchAndStoreDelegatorReward(userAddress, client, ctx, verbose); err != nil {
			time.Sleep(time.Second * 10)
			continue
		}

		if err := runFetchAndStoreHyperunitOperation(userAddress, client, ctx, verbose); err != nil {
			time.Sleep(time.Second * 10)
			continue
		}

		fmt.Println("Loading fills via REST API")
		initFills, err := fetchUserFills(userAddress)

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to load fills via API, and retrying... (%v)"))
			time.Sleep(time.Second * 10)
			continue
		}

		for _, fill := range initFills {
			user := client.Fill.Create().
				SetCoin(fill.Coin).
				SetPx(fill.Px).
				SetSz(fill.Sz).
				SetSide(fill.Side).
				SetTime(fill.Time).
				SetStartPosition(fill.StartPosition).
				SetDir(fill.Dir).
				SetHash(fill.Hash).
				SetCrossed(fill.Crossed).
				SetClosedPnl(fill.ClosedPnl).
				SetFee(fill.Fee).
				SetTid(fill.Tid).
				SetOid(fill.Oid).
				SetAddress(userAddress).
				SetFeeToken(fill.FeeToken)

			if fill.BuilderFee != nil {
				user = user.SetBuilderFee(*fill.BuilderFee)
			}

			user.OnConflict().UpdateNewValues().IDX(ctx)

			if verbose {
				fmt.Println(fmt.Sprintf(
					"[Fill] Coin: %s, Px: %s, Sz: %s, Side: %s, Dir: %s, Fee: %s, FeeToken: %s",
					fill.Coin, fill.Px, fill.Sz, fill.Side, fill.Dir, fill.Fee, fill.FeeToken,
				))
			}
		}

		fmt.Println("Complete loading fills via API")

		fmt.Println("Loading twapSliceFills via REST API")
		initUserTwapSliceFills, err := fetchUserTwapSliceFills(userAddress)

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to load fills via API, and retrying... (%v)"))
			time.Sleep(time.Second * 10)
			continue
		}

		for _, twap := range initUserTwapSliceFills {
			fill := twap.Fill

			user := client.TwapSliceFill.Create().
				SetCoin(fill.Coin).
				SetPx(fill.Px).
				SetSz(fill.Sz).
				SetSide(fill.Side).
				SetTime(fill.Time).
				SetStartPosition(fill.StartPosition).
				SetClosedPnl(fill.ClosedPnl).
				SetDir(fill.Dir).
				SetHash(fill.Hash).
				SetCrossed(fill.Crossed).
				SetFee(fill.Fee).
				SetTid(fill.Tid).
				SetOid(fill.Oid).
				SetFeeToken(fill.FeeToken).
				SetTwapID(twap.TwapID).
				SetAddress(userAddress)

			if fill.BuilderFee != nil {
				user = user.SetBuilderFee(*fill.BuilderFee)
			}

			user.OnConflict().UpdateNewValues().IDX(ctx)

			if verbose {
				fmt.Println(fmt.Sprintf(
					"[TwapSliceFill] Coin: %s, Px: %s, Sz: %s, Side: %s, Dir: %s, Fee: %s, FeeToken: %s",
					fill.Coin, fill.Px, fill.Sz, fill.Side, fill.Dir, fill.Fee, fill.FeeToken,
				))
			}
		}

		fmt.Println("Complete loading twapSliceFills via API")

		fmt.Println("Loading fundings via REST API")
		initFundings, err := fetchUserFundings(userAddress)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to load fundings via API, and retrying... (%v)", err))
			time.Sleep(time.Second * 10)
			continue
		}

		for _, funding := range initFundings {

			wip := client.Funding.Create().
				SetCoin(funding.Coin).
				SetFundingRate(funding.FundingRate).
				SetSzi(funding.Szi).
				SetUsdc(funding.Usdc).
				SetTime(funding.Time).
				SetAddress(userAddress)

			wip.OnConflict().UpdateNewValues().IDX(ctx)

			if verbose {
				fmt.Println(fmt.Sprintf(
					"[Funding] Coin: %s, USDC: %s, Time: %v",
					funding.Coin, funding.Usdc, funding.Time,
				))
			}
		}
		fmt.Println("Complete loading funding via API")

		fmt.Println("Loading UserNonFundingLedgerUpdates via REST API")
		initUpdates, err := fetchUserNonFundingLedgerUpdates(userAddress)
		if err != nil {
			fmt.Println(fmt.Sprintf("failed to load UserNonFundingLedgerUpdates via API, and retrying... (%v)", err))
			time.Sleep(time.Second * 10)
			continue
		}

		for _, update := range initUpdates {
			switch update.Delta.Type {
			case "spotGenesis":
				wip := client.SpotGenesis.Create().
					SetAmount(update.Delta.Amount).
					SetCoin(update.Delta.Token).
					SetTime(update.Time).
					SetAddress(userAddress)

				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[SpotGenesis] Coin: %s, Amount: %s, Time: %v",
						update.Delta.Token, update.Delta.Amount, update.Time,
					))
				}
			case "rewardsClaim":
				wip := client.RewardsClaim.Create().
					SetAmount(update.Delta.Amount).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[RewardsClaim] Amount: %s, Time: %v",
						update.Delta.Amount, update.Time,
					))
				}
			case "vaultLeaderCommission":
				wip := client.VaultLeaderCommission.Create().
					SetUsdc(update.Delta.Usdc).
					SetUser(update.Delta.User).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[VaultLeaderCommission] USDC: %s, User: %s, Time: %v",
						update.Delta.Usdc, update.Delta.User, update.Time,
					))
				}
			case "vaultWithdraw":
				wip := client.VaultWithdrawal.Create().
					SetVault(update.Delta.Vault).
					SetUser(update.Delta.User).
					SetRequestedUsd(update.Delta.RequestedUsd).
					SetCommission(update.Delta.Commission).
					SetClosingCost(update.Delta.ClosingCost).
					SetBasis(update.Delta.Basis).
					SetNetWithdrawnUsd(update.Delta.NetWithdrawnUsd).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[vaultWithdraw] USDC: %s, Vault: %s, Time: %v",
						update.Delta.NetWithdrawnUsd, update.Delta.Vault, update.Time,
					))
				}
			case "vaultCreate", "vaultDeposit", "vaultDistribution":
				wip := client.VaultDelta.Create().
					SetType(update.Delta.Type).
					SetUsdc(update.Delta.Usdc).
					SetVault(update.Delta.Vault).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[VaultDelta] Type: %s, USDC: %s, Vault: %s, Time: %v",
						update.Delta.Type, update.Delta.Usdc, update.Delta.Vault, update.Time,
					))
				}
			case "internalTransfer":
				wip := client.InternalTransfer.Create().
					SetUsdc(update.Delta.Usdc).
					SetUser(update.Delta.User).
					SetDestination(update.Delta.Destination).
					SetFee(update.Delta.Fee).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[InternalTransfer] User: %s, Dest: %s, USDC: %s, Fee: %s, Time: %v",
						update.Delta.User, update.Delta.Destination, update.Delta.Usdc, update.Delta.Fee, update.Time,
					))
				}
			case "spotTransfer":
				wip := client.SpotTransfer.Create().
					SetToken(update.Delta.Token).
					SetAmount(update.Delta.Amount).
					SetUser(update.Delta.User).
					SetDestination(update.Delta.Destination).
					SetFee(update.Delta.Fee).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[SpotTransfer] User: %s, Dest: %s, Token:%s, Amount: %s, Fee: %s, Time: %v",
						update.Delta.User, update.Delta.Destination, update.Delta.Token, update.Delta.Amount, update.Delta.Fee, update.Time,
					))
				}
			case "withdraw":
				wip := client.Withdraw.Create().
					SetUsdc(update.Delta.Usdc).
					SetNonce(update.Delta.Nonce).
					SetFee(update.Delta.Fee).
					SetTime(update.Time).
					SetAddress(userAddress)
				wip.OnConflict().UpdateNewValues().IDX(ctx)
				if verbose {
					fmt.Println(fmt.Sprintf(
						"[Withdraw] Usdc:%s, Fee: %s, Time: %v",
						update.Delta.User, update.Delta.Usdc, update.Delta.Fee, update.Time,
					))
				}
			}
		}

		fmt.Println("Complete loading UserNonFundingLedgerUpdates via API")

	Loop:

		for {
			select {
			case <-ticker.C:
				runBackup(outputFile, backupFile, awsS3Region)
			case <-periodicSyncTicker.C:

				// ignore error and hope it will be handled at next iter.
				runFetchAndStoreDelegatorHistory(userAddress, client, ctx, verbose)
				runFetchAndStoreDelegatorReward(userAddress, client, ctx, verbose)
				runFetchAndStoreHyperunitOperation(userAddress, client, ctx, verbose)

			case msg, ok := <-msgC:
				if !ok {
					fmt.Println("websocket connection disconnected and restarting...")
					break Loop
				}

				ch := Msg{}
				json.Unmarshal(msg, &ch)

				switch ch.Channel {
				case "userFills":
					data := &WsUserFillsResponse{}

					if err := json.Unmarshal(msg, &data); err != nil {
						fmt.Println("failed to unmarshal JSON: %v", err)
						break
					}

					for _, fill := range data.Data.Fills {

						user := client.Fill.Create().
							SetCoin(fill.Coin).
							SetPx(fill.Px).
							SetSz(fill.Sz).
							SetSide(fill.Side).
							SetTime(fill.Time).
							SetStartPosition(fill.StartPosition).
							SetClosedPnl(fill.ClosedPnl).
							SetDir(fill.Dir).
							SetHash(fill.Hash).
							SetCrossed(fill.Crossed).
							SetFee(fill.Fee).
							SetTid(fill.Tid).
							SetOid(fill.Oid).
							SetFeeToken(fill.FeeToken).
							SetAddress(userAddress)

						if fill.BuilderFee != nil {
							user = user.SetBuilderFee(*fill.BuilderFee)
						}

						user.OnConflict().UpdateNewValues().IDX(ctx)

						if verbose {
							fmt.Println(fmt.Sprintf(
								"[Fill] Coin: %s, Px: %s, Sz: %s, Side: %s, Dir: %s, Fee: %s, FeeToken: %s",
								fill.Coin, fill.Px, fill.Sz, fill.Side, fill.Dir, fill.Fee, fill.FeeToken,
							))
						}
					}
				case "userFundings":
					data := &WsUserFundingsResponse{}
					if err := json.Unmarshal(msg, &data); err != nil {
						fmt.Println("failed to unmarshal JSON: %v", err)
						break
					}

					for _, funding := range data.Data.Fundings {

						wip := client.Funding.Create().
							SetCoin(funding.Coin).
							SetFundingRate(funding.FundingRate).
							SetSzi(funding.Szi).
							SetUsdc(funding.Usdc).
							SetTime(funding.Time).
							SetAddress(userAddress)

						wip.OnConflict().UpdateNewValues().IDX(ctx)

						if verbose {
							fmt.Println(fmt.Sprintf(
								"[Funding] Coin: %s, USDC: %s, Time: %v",
								funding.Coin, funding.Usdc, funding.Time,
							))
						}
					}
				case "userTwapSliceFills":
					data := &WSUserTwapSliceFillsResponse{}
					if err := json.Unmarshal(msg, &data); err != nil {
						fmt.Println("failed to unmarshal JSON: %v", err)
						break
					}

					for _, twap := range data.Data.TwapSliceFills {
						fill := twap.Fill

						user := client.TwapSliceFill.Create().
							SetCoin(fill.Coin).
							SetPx(fill.Px).
							SetSz(fill.Sz).
							SetSide(fill.Side).
							SetTime(fill.Time).
							SetStartPosition(fill.StartPosition).
							SetClosedPnl(fill.ClosedPnl).
							SetDir(fill.Dir).
							SetHash(fill.Hash).
							SetCrossed(fill.Crossed).
							SetFee(fill.Fee).
							SetTid(fill.Tid).
							SetOid(fill.Oid).
							SetFeeToken(fill.FeeToken).
							SetTwapID(twap.TwapID).
							SetAddress(userAddress)

						if fill.BuilderFee != nil {
							user = user.SetBuilderFee(*fill.BuilderFee)
						}

						user.OnConflict().UpdateNewValues().IDX(ctx)

						if verbose {
							fmt.Println(fmt.Sprintf(
								"[TwapSliceFill] Coin: %s, Px: %s, Sz: %s, Side: %s, Dir: %s, Fee: %s, FeeToken: %s",
								fill.Coin, fill.Px, fill.Sz, fill.Side, fill.Dir, fill.Fee, fill.FeeToken,
							))
						}

					}

				case "userNonFundingLedgerUpdates":
					data := &WSUserNonFundingLedgerUpdatesResponse{}
					if err := json.Unmarshal(msg, &data); err != nil {
						fmt.Println("failed to unmarshal JSON: %v", err)
						break
					}
					for _, update := range data.Data.NonFundingLedgerUpdates {
						switch update.Delta.Type {
						case "spotGenesis":
							wip := client.SpotGenesis.Create().
								SetAmount(update.Delta.Amount).
								SetCoin(update.Delta.Token).
								SetTime(update.Time).
								SetAddress(userAddress)

							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[SpotGenesis] Coin: %s, Amount: %s, Time: %v",
									update.Delta.Token, update.Delta.Amount, update.Time,
								))
							}
						case "rewardsClaim":
							wip := client.RewardsClaim.Create().
								SetAmount(update.Delta.Amount).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[RewardsClaim] Amount: %s, Time: %v",
									update.Delta.Amount, update.Time,
								))
							}
						case "vaultLeaderCommission":
							wip := client.VaultLeaderCommission.Create().
								SetUsdc(update.Delta.Usdc).
								SetUser(update.Delta.User).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[VaultLeaderCommission] USDC: %s, User: %s, Time: %v",
									update.Delta.Usdc, update.Delta.User, update.Time,
								))
							}
						case "vaultWithdraw":
							wip := client.VaultWithdrawal.Create().
								SetVault(update.Delta.Vault).
								SetUser(update.Delta.User).
								SetRequestedUsd(update.Delta.RequestedUsd).
								SetCommission(update.Delta.Commission).
								SetClosingCost(update.Delta.ClosingCost).
								SetBasis(update.Delta.Basis).
								SetNetWithdrawnUsd(update.Delta.NetWithdrawnUsd).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[vaultWithdraw] USDC: %s, Vault: %s, Time: %v",
									update.Delta.NetWithdrawnUsd, update.Delta.Vault, update.Time,
								))
							}
						case "vaultCreate", "vaultDeposit", "vaultDistribution":
							wip := client.VaultDelta.Create().
								SetType(update.Delta.Type).
								SetUsdc(update.Delta.Usdc).
								SetVault(update.Delta.Vault).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[VaultDelta] Type: %s, USDC: %s, Vault: %s, Time: %v",
									update.Delta.Type, update.Delta.Usdc, update.Delta.Vault, update.Time,
								))
							}
						case "internalTransfer":
							wip := client.InternalTransfer.Create().
								SetUsdc(update.Delta.Usdc).
								SetUser(update.Delta.User).
								SetDestination(update.Delta.Destination).
								SetFee(update.Delta.Fee).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[InternalTransfer] User: %s, Dest: %s, USDC: %s, Fee: %s, Time: %v",
									update.Delta.User, update.Delta.Destination, update.Delta.Usdc, update.Delta.Fee, update.Time,
								))
							}
						case "spotTransfer":
							wip := client.SpotTransfer.Create().
								SetToken(update.Delta.Token).
								SetAmount(update.Delta.Amount).
								SetUser(update.Delta.User).
								SetDestination(update.Delta.Destination).
								SetFee(update.Delta.Fee).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[SpotTransfer] User: %s, Dest: %s, Token:%s, Amount: %s, Fee: %s, Time: %v",
									update.Delta.User, update.Delta.Destination, update.Delta.Token, update.Delta.Amount, update.Delta.Fee, update.Time,
								))
							}
						case "withdraw":
							wip := client.Withdraw.Create().
								SetUsdc(update.Delta.Usdc).
								SetNonce(update.Delta.Nonce).
								SetFee(update.Delta.Fee).
								SetTime(update.Time).
								SetAddress(userAddress)
							wip.OnConflict().UpdateNewValues().IDX(ctx)
							if verbose {
								fmt.Println(fmt.Sprintf(
									"[Withdraw] User: %s, Usdc:%s, Fee: %s, Time: %v",
									update.Delta.User, update.Delta.Usdc, update.Delta.Fee, update.Time,
								))
							}
						}
					}
				case "subscriptionResponse":
					continue
				case "pong":
					continue
				default:
					fmt.Println("Unhandled message: ", string(msg))
				}

			}
		}

	}
}
