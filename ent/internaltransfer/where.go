// Code generated by ent, DO NOT EDIT.

package internaltransfer

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldID, id))
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldUser, v))
}

// Destination applies equality check predicate on the "destination" field. It's identical to DestinationEQ.
func Destination(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldDestination, v))
}

// Usdc applies equality check predicate on the "usdc" field. It's identical to UsdcEQ.
func Usdc(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldUsdc, v))
}

// Fee applies equality check predicate on the "fee" field. It's identical to FeeEQ.
func Fee(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldFee, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldTime, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldAddress, v))
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldUser, v))
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldUser, v))
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldUser, vs...))
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldUser, vs...))
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldUser, v))
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldUser, v))
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldUser, v))
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldUser, v))
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContains(FieldUser, v))
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasPrefix(FieldUser, v))
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasSuffix(FieldUser, v))
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEqualFold(FieldUser, v))
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContainsFold(FieldUser, v))
}

// DestinationEQ applies the EQ predicate on the "destination" field.
func DestinationEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldDestination, v))
}

// DestinationNEQ applies the NEQ predicate on the "destination" field.
func DestinationNEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldDestination, v))
}

// DestinationIn applies the In predicate on the "destination" field.
func DestinationIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldDestination, vs...))
}

// DestinationNotIn applies the NotIn predicate on the "destination" field.
func DestinationNotIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldDestination, vs...))
}

// DestinationGT applies the GT predicate on the "destination" field.
func DestinationGT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldDestination, v))
}

// DestinationGTE applies the GTE predicate on the "destination" field.
func DestinationGTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldDestination, v))
}

// DestinationLT applies the LT predicate on the "destination" field.
func DestinationLT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldDestination, v))
}

// DestinationLTE applies the LTE predicate on the "destination" field.
func DestinationLTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldDestination, v))
}

// DestinationContains applies the Contains predicate on the "destination" field.
func DestinationContains(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContains(FieldDestination, v))
}

// DestinationHasPrefix applies the HasPrefix predicate on the "destination" field.
func DestinationHasPrefix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasPrefix(FieldDestination, v))
}

// DestinationHasSuffix applies the HasSuffix predicate on the "destination" field.
func DestinationHasSuffix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasSuffix(FieldDestination, v))
}

// DestinationEqualFold applies the EqualFold predicate on the "destination" field.
func DestinationEqualFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEqualFold(FieldDestination, v))
}

// DestinationContainsFold applies the ContainsFold predicate on the "destination" field.
func DestinationContainsFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContainsFold(FieldDestination, v))
}

// UsdcEQ applies the EQ predicate on the "usdc" field.
func UsdcEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldUsdc, v))
}

// UsdcNEQ applies the NEQ predicate on the "usdc" field.
func UsdcNEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldUsdc, v))
}

// UsdcIn applies the In predicate on the "usdc" field.
func UsdcIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldUsdc, vs...))
}

// UsdcNotIn applies the NotIn predicate on the "usdc" field.
func UsdcNotIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldUsdc, vs...))
}

// UsdcGT applies the GT predicate on the "usdc" field.
func UsdcGT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldUsdc, v))
}

// UsdcGTE applies the GTE predicate on the "usdc" field.
func UsdcGTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldUsdc, v))
}

// UsdcLT applies the LT predicate on the "usdc" field.
func UsdcLT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldUsdc, v))
}

// UsdcLTE applies the LTE predicate on the "usdc" field.
func UsdcLTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldUsdc, v))
}

// UsdcContains applies the Contains predicate on the "usdc" field.
func UsdcContains(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContains(FieldUsdc, v))
}

// UsdcHasPrefix applies the HasPrefix predicate on the "usdc" field.
func UsdcHasPrefix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasPrefix(FieldUsdc, v))
}

// UsdcHasSuffix applies the HasSuffix predicate on the "usdc" field.
func UsdcHasSuffix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasSuffix(FieldUsdc, v))
}

// UsdcEqualFold applies the EqualFold predicate on the "usdc" field.
func UsdcEqualFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEqualFold(FieldUsdc, v))
}

// UsdcContainsFold applies the ContainsFold predicate on the "usdc" field.
func UsdcContainsFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContainsFold(FieldUsdc, v))
}

// FeeEQ applies the EQ predicate on the "fee" field.
func FeeEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldFee, v))
}

// FeeNEQ applies the NEQ predicate on the "fee" field.
func FeeNEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldFee, v))
}

// FeeIn applies the In predicate on the "fee" field.
func FeeIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldFee, vs...))
}

// FeeNotIn applies the NotIn predicate on the "fee" field.
func FeeNotIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldFee, vs...))
}

// FeeGT applies the GT predicate on the "fee" field.
func FeeGT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldFee, v))
}

// FeeGTE applies the GTE predicate on the "fee" field.
func FeeGTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldFee, v))
}

// FeeLT applies the LT predicate on the "fee" field.
func FeeLT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldFee, v))
}

// FeeLTE applies the LTE predicate on the "fee" field.
func FeeLTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldFee, v))
}

// FeeContains applies the Contains predicate on the "fee" field.
func FeeContains(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContains(FieldFee, v))
}

// FeeHasPrefix applies the HasPrefix predicate on the "fee" field.
func FeeHasPrefix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasPrefix(FieldFee, v))
}

// FeeHasSuffix applies the HasSuffix predicate on the "fee" field.
func FeeHasSuffix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasSuffix(FieldFee, v))
}

// FeeEqualFold applies the EqualFold predicate on the "fee" field.
func FeeEqualFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEqualFold(FieldFee, v))
}

// FeeContainsFold applies the ContainsFold predicate on the "fee" field.
func FeeContainsFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContainsFold(FieldFee, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldTime, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.FieldContainsFold(FieldAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InternalTransfer) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InternalTransfer) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InternalTransfer) predicate.InternalTransfer {
	return predicate.InternalTransfer(sql.NotPredicates(p))
}
