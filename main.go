package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/urfave/cli/v3"
	"github.com/yoshiso/hypersync/ent"
	_ "modernc.org/sqlite"
)

type WSRequest struct {
	Method       string               `json:"method"`
	Subscription UserSubscription `json:"subscription"`
}

type UserSubscription struct {
	Type string `json:"type"`
	User string `json:"user"`
}

type FillLiquidation struct {
	LiquidatedUser *string `json:"liquidatedUser"`
	MarkPx string `json:"markPx"`
	Method string `json:"method"`	
}

type WsUserFunding struct {
	Time int64 `json:"time"`
	Coin string `json:"coin"`
	Usdc string `json:"usdc"`
	Szi string `json:"szi"`
	FundingRate string `json:"fundingRate"`
}

type WsFill struct {
	Coin string `json:"coin"`
	Px string `json:"px"`
	Sz string `json:"sz"`
	Side string `json:"side"`
	Time int64 `json:"time"`
	StartPosition string `json:"startPosition"`
	Dir string `json:"dir"`
	ClosedPnl string `json:"closedPnl"`
	Hash string `json:"hash"`
	Oid int64 `json:"oid"`
	Crossed bool `json:"crossed"`
	Fee string `json:"fee"`
	Tid int64 `json:"tid"`
	Liquidation *FillLiquidation `json:"liquidation"`
	FeeToken string `json:"feeToken"`
	BuilderFee *string `json:"builderFee"`
}

type WsUserFundings struct {
	IsSnapshot bool `json:"isSnapshot"`
	User string `json:"user"`
	Fundings []WsUserFunding `json:"fundings"`
}

type WsUserFills struct {
	IsSnapshot bool `json:"isSnapshot"`
	User string `json:"user"`
	Fills []WsFill `json:"fills"`
}

type WsUserFillsResponse struct {
	Channel string	`json:"channel"`
	Data    WsUserFills `json:"data"`
}

type WsUserFundingsResponse struct {
	Channel string	`json:"channel"`
	Data    WsUserFundings `json:"data"`
}

type HTTPUserFundingDelta struct {
	Coin string `json:"coin"`
	Usdc string `json:"usdc"`
	Szi string `json:"szi"`
	FundingRate string `json:"fundingRate"`
}

type HTTPUserFunding struct {
	Time int64 `json:"time"`
	Hash string `json:"hash"`
	Delta HTTPUserFundingDelta `json:"delta"`
}

type Msg struct {
	Channel string `json:"channel"`
}

type Ping struct {
	Method string `json:"method"`
}

func Conn(userAddress string) (chan []byte) {
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

	if err := c.WriteJSON(WSRequest{Method: "subscribe", Subscription: UserSubscription{User: userAddress, Type: "userFills"} }); err != nil {
		fmt.Println(fmt.Errorf("ws.WriteJSON error %v", err))
	}

	if err := c.WriteJSON(WSRequest{Method: "subscribe", Subscription: UserSubscription{User: userAddress, Type: "userFundings"} }); err != nil {
		fmt.Println(fmt.Errorf("ws.WriteJSON error %v", err))
	}

	return msgC
}


type HTTPUserRequest struct {
	User   string    `json:"user"`
	Type   string    `json:"type"`
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
			Coin: funding.Delta.Coin,
			Time: funding.Time,
			Usdc: funding.Delta.Usdc,
			Szi: funding.Delta.Szi,
			FundingRate: funding.Delta.FundingRate,
		})
	}

	return fundings, nil
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
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {
            
			userAddress := cmd.String("address")
			
			if userAddress == "" {
				fmt.Println("Add --address flag with your wallet")
				os.Exit(1)
			}

			verbose := cmd.Bool("verbose")


			output_file := cmd.String("out")

			run(userAddress, verbose, output_file)

            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}


func run(userAddress string, verbose bool, output_file string) {

	fmt.Println("start running server", userAddress, output_file)

	client, err := ent.Open("sqlite3", "file:" + output_file + "?cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()

	// Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
	
	for {
		msgC := Conn(userAddress)

		ctx := context.TODO()

		fmt.Println("Loading fills via REST API")
		initFills, err := fetchUserFills(userAddress)

		if err != nil {
			fmt.Println("failed to load fills via API, and retrying... (%v)", err)
			time.Sleep(time.Second * 5)
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
					"Coin: %s, Px: %s, Sz: %s, Side: %s, Dir: %s, Fee: %s, FeeToken: %s",
					fill.Coin, fill.Px, fill.Sz, fill.Side, fill.Dir, fill.Fee, fill.FeeToken,
				))
			}
		}

		fmt.Println("Complete loading fills via API")

		fmt.Println("Loading fundings via REST API")
		initFundings, err := fetchUserFundings(userAddress)
		if err != nil {
			fmt.Println("failed to load fundings via API, and retrying... (%v)", err)
			time.Sleep(time.Second * 5)
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
					"Coin: %s, USDC: %s, Time: %v",
					funding.Coin, funding.Usdc, funding.Time,
				))
			}			
		}
		fmt.Println("Complete loading funding via API")
		
		Loop:

			for {
				select {
				case msg, ok := <- msgC:
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
									"Coin: %s, Px: %s, Sz: %s, Side: %s, Dir: %s, Fee: %s, FeeToken: %s",
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
									"Coin: %s, USDC: %s, Time: %v",
									funding.Coin, funding.Usdc, funding.Time,
								))
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

