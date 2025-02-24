// Code generated by ent, DO NOT EDIT.

package vaultdelta

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the vaultdelta type in the database.
	Label = "vault_delta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldVault holds the string denoting the vault field in the database.
	FieldVault = "vault"
	// FieldUsdc holds the string denoting the usdc field in the database.
	FieldUsdc = "usdc"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// Table holds the table name of the vaultdelta in the database.
	Table = "vault_delta"
)

// Columns holds all SQL columns for vaultdelta fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldVault,
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

// OrderOption defines the ordering options for the VaultDelta queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByVault orders the results by the vault field.
func ByVault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVault, opts...).ToFunc()
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
