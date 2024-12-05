// Code generated by ent, DO NOT EDIT.

package fill

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the fill type in the database.
	Label = "fill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCoin holds the string denoting the coin field in the database.
	FieldCoin = "coin"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPx holds the string denoting the px field in the database.
	FieldPx = "px"
	// FieldSz holds the string denoting the sz field in the database.
	FieldSz = "sz"
	// FieldSide holds the string denoting the side field in the database.
	FieldSide = "side"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldStartPosition holds the string denoting the start_position field in the database.
	FieldStartPosition = "start_position"
	// FieldDir holds the string denoting the dir field in the database.
	FieldDir = "dir"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldCrossed holds the string denoting the crossed field in the database.
	FieldCrossed = "crossed"
	// FieldFee holds the string denoting the fee field in the database.
	FieldFee = "fee"
	// FieldOid holds the string denoting the oid field in the database.
	FieldOid = "oid"
	// FieldTid holds the string denoting the tid field in the database.
	FieldTid = "tid"
	// FieldFeeToken holds the string denoting the fee_token field in the database.
	FieldFeeToken = "fee_token"
	// FieldBuilderFee holds the string denoting the builder_fee field in the database.
	FieldBuilderFee = "builder_fee"
	// Table holds the table name of the fill in the database.
	Table = "fills"
)

// Columns holds all SQL columns for fill fields.
var Columns = []string{
	FieldID,
	FieldCoin,
	FieldAddress,
	FieldPx,
	FieldSz,
	FieldSide,
	FieldTime,
	FieldStartPosition,
	FieldDir,
	FieldHash,
	FieldCrossed,
	FieldFee,
	FieldOid,
	FieldTid,
	FieldFeeToken,
	FieldBuilderFee,
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

// OrderOption defines the ordering options for the Fill queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCoin orders the results by the coin field.
func ByCoin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoin, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByPx orders the results by the px field.
func ByPx(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPx, opts...).ToFunc()
}

// BySz orders the results by the sz field.
func BySz(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSz, opts...).ToFunc()
}

// BySide orders the results by the side field.
func BySide(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSide, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByStartPosition orders the results by the start_position field.
func ByStartPosition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartPosition, opts...).ToFunc()
}

// ByDir orders the results by the dir field.
func ByDir(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDir, opts...).ToFunc()
}

// ByHash orders the results by the hash field.
func ByHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHash, opts...).ToFunc()
}

// ByCrossed orders the results by the crossed field.
func ByCrossed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCrossed, opts...).ToFunc()
}

// ByFee orders the results by the fee field.
func ByFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFee, opts...).ToFunc()
}

// ByOid orders the results by the oid field.
func ByOid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOid, opts...).ToFunc()
}

// ByTid orders the results by the tid field.
func ByTid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTid, opts...).ToFunc()
}

// ByFeeToken orders the results by the fee_token field.
func ByFeeToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeeToken, opts...).ToFunc()
}

// ByBuilderFee orders the results by the builder_fee field.
func ByBuilderFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuilderFee, opts...).ToFunc()
}
