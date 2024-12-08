// Code generated by ent, DO NOT EDIT.

package funding

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the funding type in the database.
	Label = "funding"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldCoin holds the string denoting the coin field in the database.
	FieldCoin = "coin"
	// FieldUsdc holds the string denoting the usdc field in the database.
	FieldUsdc = "usdc"
	// FieldSzi holds the string denoting the szi field in the database.
	FieldSzi = "szi"
	// FieldFundingRate holds the string denoting the funding_rate field in the database.
	FieldFundingRate = "funding_rate"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// Table holds the table name of the funding in the database.
	Table = "fundings"
)

// Columns holds all SQL columns for funding fields.
var Columns = []string{
	FieldID,
	FieldTime,
	FieldCoin,
	FieldUsdc,
	FieldSzi,
	FieldFundingRate,
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

// OrderOption defines the ordering options for the Funding queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByCoin orders the results by the coin field.
func ByCoin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoin, opts...).ToFunc()
}

// ByUsdc orders the results by the usdc field.
func ByUsdc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsdc, opts...).ToFunc()
}

// BySzi orders the results by the szi field.
func BySzi(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSzi, opts...).ToFunc()
}

// ByFundingRate orders the results by the funding_rate field.
func ByFundingRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFundingRate, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}
