// Code generated by ent, DO NOT EDIT.

package vaultleadercommission

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the vaultleadercommission type in the database.
	Label = "vault_leader_commission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldUsdc holds the string denoting the usdc field in the database.
	FieldUsdc = "usdc"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// Table holds the table name of the vaultleadercommission in the database.
	Table = "vault_leader_commissions"
)

// Columns holds all SQL columns for vaultleadercommission fields.
var Columns = []string{
	FieldID,
	FieldUser,
	FieldUsdc,
	FieldTime,
	FieldAddress,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the VaultLeaderCommission queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUser orders the results by the user field.
func ByUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUser, opts...).ToFunc()
}

// ByUsdc orders the results by the usdc field.
func ByUsdc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsdc, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}
