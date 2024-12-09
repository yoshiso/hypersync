// Code generated by ent, DO NOT EDIT.

package withdraw

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the withdraw type in the database.
	Label = "withdraw"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsdc holds the string denoting the usdc field in the database.
	FieldUsdc = "usdc"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldFee holds the string denoting the fee field in the database.
	FieldFee = "fee"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// Table holds the table name of the withdraw in the database.
	Table = "withdraws"
)

// Columns holds all SQL columns for withdraw fields.
var Columns = []string{
	FieldID,
	FieldUsdc,
	FieldNonce,
	FieldFee,
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

// OrderOption defines the ordering options for the Withdraw queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsdc orders the results by the usdc field.
func ByUsdc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsdc, opts...).ToFunc()
}

// ByNonce orders the results by the nonce field.
func ByNonce(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNonce, opts...).ToFunc()
}

// ByFee orders the results by the fee field.
func ByFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFee, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}