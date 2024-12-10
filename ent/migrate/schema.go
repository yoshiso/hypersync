// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FillsColumns holds the columns for the "fills" table.
	FillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "coin", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "px", Type: field.TypeString},
		{Name: "sz", Type: field.TypeString},
		{Name: "side", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "start_position", Type: field.TypeString},
		{Name: "closed_pnl", Type: field.TypeString},
		{Name: "dir", Type: field.TypeString},
		{Name: "hash", Type: field.TypeString},
		{Name: "crossed", Type: field.TypeBool},
		{Name: "fee", Type: field.TypeString},
		{Name: "oid", Type: field.TypeInt64},
		{Name: "tid", Type: field.TypeInt64, Unique: true},
		{Name: "fee_token", Type: field.TypeString},
		{Name: "builder_fee", Type: field.TypeString, Nullable: true},
	}
	// FillsTable holds the schema information for the "fills" table.
	FillsTable = &schema.Table{
		Name:       "fills",
		Columns:    FillsColumns,
		PrimaryKey: []*schema.Column{FillsColumns[0]},
	}
	// FundingsColumns holds the columns for the "fundings" table.
	FundingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "time", Type: field.TypeInt64},
		{Name: "coin", Type: field.TypeString},
		{Name: "usdc", Type: field.TypeString},
		{Name: "szi", Type: field.TypeString},
		{Name: "funding_rate", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
	}
	// FundingsTable holds the schema information for the "fundings" table.
	FundingsTable = &schema.Table{
		Name:       "fundings",
		Columns:    FundingsColumns,
		PrimaryKey: []*schema.Column{FundingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "funding_address_time_coin",
				Unique:  true,
				Columns: []*schema.Column{FundingsColumns[6], FundingsColumns[1], FundingsColumns[2]},
			},
		},
	}
	// InternalTransfersColumns holds the columns for the "internal_transfers" table.
	InternalTransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeString},
		{Name: "destination", Type: field.TypeString},
		{Name: "usdc", Type: field.TypeString},
		{Name: "fee", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// InternalTransfersTable holds the schema information for the "internal_transfers" table.
	InternalTransfersTable = &schema.Table{
		Name:       "internal_transfers",
		Columns:    InternalTransfersColumns,
		PrimaryKey: []*schema.Column{InternalTransfersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "internaltransfer_address_time_user_destination",
				Unique:  true,
				Columns: []*schema.Column{InternalTransfersColumns[6], InternalTransfersColumns[5], InternalTransfersColumns[1], InternalTransfersColumns[2]},
			},
		},
	}
	// RewardsClaimsColumns holds the columns for the "rewards_claims" table.
	RewardsClaimsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// RewardsClaimsTable holds the schema information for the "rewards_claims" table.
	RewardsClaimsTable = &schema.Table{
		Name:       "rewards_claims",
		Columns:    RewardsClaimsColumns,
		PrimaryKey: []*schema.Column{RewardsClaimsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "rewardsclaim_address_time",
				Unique:  true,
				Columns: []*schema.Column{RewardsClaimsColumns[3], RewardsClaimsColumns[2]},
			},
		},
	}
	// SpotGenesesColumns holds the columns for the "spot_geneses" table.
	SpotGenesesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "coin", Type: field.TypeString},
		{Name: "amount", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// SpotGenesesTable holds the schema information for the "spot_geneses" table.
	SpotGenesesTable = &schema.Table{
		Name:       "spot_geneses",
		Columns:    SpotGenesesColumns,
		PrimaryKey: []*schema.Column{SpotGenesesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "spotgenesis_address_time_coin",
				Unique:  true,
				Columns: []*schema.Column{SpotGenesesColumns[4], SpotGenesesColumns[3], SpotGenesesColumns[1]},
			},
		},
	}
	// SpotTransfersColumns holds the columns for the "spot_transfers" table.
	SpotTransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeString},
		{Name: "destination", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "amount", Type: field.TypeString},
		{Name: "fee", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// SpotTransfersTable holds the schema information for the "spot_transfers" table.
	SpotTransfersTable = &schema.Table{
		Name:       "spot_transfers",
		Columns:    SpotTransfersColumns,
		PrimaryKey: []*schema.Column{SpotTransfersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "spottransfer_address_time_token_user_destination",
				Unique:  true,
				Columns: []*schema.Column{SpotTransfersColumns[7], SpotTransfersColumns[6], SpotTransfersColumns[3], SpotTransfersColumns[1], SpotTransfersColumns[2]},
			},
		},
	}
	// VaultDeltaColumns holds the columns for the "vault_delta" table.
	VaultDeltaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "vault", Type: field.TypeString},
		{Name: "usdc", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// VaultDeltaTable holds the schema information for the "vault_delta" table.
	VaultDeltaTable = &schema.Table{
		Name:       "vault_delta",
		Columns:    VaultDeltaColumns,
		PrimaryKey: []*schema.Column{VaultDeltaColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vaultdelta_address_time_vault",
				Unique:  true,
				Columns: []*schema.Column{VaultDeltaColumns[5], VaultDeltaColumns[4], VaultDeltaColumns[2]},
			},
		},
	}
	// VaultLeaderCommissionsColumns holds the columns for the "vault_leader_commissions" table.
	VaultLeaderCommissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeString},
		{Name: "usdc", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// VaultLeaderCommissionsTable holds the schema information for the "vault_leader_commissions" table.
	VaultLeaderCommissionsTable = &schema.Table{
		Name:       "vault_leader_commissions",
		Columns:    VaultLeaderCommissionsColumns,
		PrimaryKey: []*schema.Column{VaultLeaderCommissionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vaultleadercommission_address_time_user",
				Unique:  true,
				Columns: []*schema.Column{VaultLeaderCommissionsColumns[4], VaultLeaderCommissionsColumns[3], VaultLeaderCommissionsColumns[1]},
			},
		},
	}
	// VaultWithdrawalsColumns holds the columns for the "vault_withdrawals" table.
	VaultWithdrawalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "vault", Type: field.TypeString},
		{Name: "user", Type: field.TypeString},
		{Name: "requested_usd", Type: field.TypeString},
		{Name: "commission", Type: field.TypeString},
		{Name: "closing_cost", Type: field.TypeString},
		{Name: "basis", Type: field.TypeString},
		{Name: "net_withdrawn_usd", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// VaultWithdrawalsTable holds the schema information for the "vault_withdrawals" table.
	VaultWithdrawalsTable = &schema.Table{
		Name:       "vault_withdrawals",
		Columns:    VaultWithdrawalsColumns,
		PrimaryKey: []*schema.Column{VaultWithdrawalsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vaultwithdrawal_address_time_vault",
				Unique:  true,
				Columns: []*schema.Column{VaultWithdrawalsColumns[9], VaultWithdrawalsColumns[8], VaultWithdrawalsColumns[1]},
			},
		},
	}
	// WithdrawsColumns holds the columns for the "withdraws" table.
	WithdrawsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "usdc", Type: field.TypeString},
		{Name: "nonce", Type: field.TypeInt64},
		{Name: "fee", Type: field.TypeString},
		{Name: "time", Type: field.TypeInt64},
		{Name: "address", Type: field.TypeString},
	}
	// WithdrawsTable holds the schema information for the "withdraws" table.
	WithdrawsTable = &schema.Table{
		Name:       "withdraws",
		Columns:    WithdrawsColumns,
		PrimaryKey: []*schema.Column{WithdrawsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "withdraw_address_time_nonce",
				Unique:  true,
				Columns: []*schema.Column{WithdrawsColumns[5], WithdrawsColumns[4], WithdrawsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FillsTable,
		FundingsTable,
		InternalTransfersTable,
		RewardsClaimsTable,
		SpotGenesesTable,
		SpotTransfersTable,
		VaultDeltaTable,
		VaultLeaderCommissionsTable,
		VaultWithdrawalsTable,
		WithdrawsTable,
	}
)

func init() {
}
