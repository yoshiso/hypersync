// Code generated by ent, DO NOT EDIT.

package withdraw

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLTE(FieldID, id))
}

// Usdc applies equality check predicate on the "usdc" field. It's identical to UsdcEQ.
func Usdc(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldUsdc, v))
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldNonce, v))
}

// Fee applies equality check predicate on the "fee" field. It's identical to FeeEQ.
func Fee(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldFee, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldTime, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldAddress, v))
}

// UsdcEQ applies the EQ predicate on the "usdc" field.
func UsdcEQ(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldUsdc, v))
}

// UsdcNEQ applies the NEQ predicate on the "usdc" field.
func UsdcNEQ(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNEQ(FieldUsdc, v))
}

// UsdcIn applies the In predicate on the "usdc" field.
func UsdcIn(vs ...string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldIn(FieldUsdc, vs...))
}

// UsdcNotIn applies the NotIn predicate on the "usdc" field.
func UsdcNotIn(vs ...string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNotIn(FieldUsdc, vs...))
}

// UsdcGT applies the GT predicate on the "usdc" field.
func UsdcGT(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGT(FieldUsdc, v))
}

// UsdcGTE applies the GTE predicate on the "usdc" field.
func UsdcGTE(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGTE(FieldUsdc, v))
}

// UsdcLT applies the LT predicate on the "usdc" field.
func UsdcLT(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLT(FieldUsdc, v))
}

// UsdcLTE applies the LTE predicate on the "usdc" field.
func UsdcLTE(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLTE(FieldUsdc, v))
}

// UsdcContains applies the Contains predicate on the "usdc" field.
func UsdcContains(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldContains(FieldUsdc, v))
}

// UsdcHasPrefix applies the HasPrefix predicate on the "usdc" field.
func UsdcHasPrefix(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldHasPrefix(FieldUsdc, v))
}

// UsdcHasSuffix applies the HasSuffix predicate on the "usdc" field.
func UsdcHasSuffix(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldHasSuffix(FieldUsdc, v))
}

// UsdcEqualFold applies the EqualFold predicate on the "usdc" field.
func UsdcEqualFold(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEqualFold(FieldUsdc, v))
}

// UsdcContainsFold applies the ContainsFold predicate on the "usdc" field.
func UsdcContainsFold(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldContainsFold(FieldUsdc, v))
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldNonce, v))
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNEQ(FieldNonce, v))
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldIn(FieldNonce, vs...))
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNotIn(FieldNonce, vs...))
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGT(FieldNonce, v))
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGTE(FieldNonce, v))
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLT(FieldNonce, v))
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLTE(FieldNonce, v))
}

// FeeEQ applies the EQ predicate on the "fee" field.
func FeeEQ(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldFee, v))
}

// FeeNEQ applies the NEQ predicate on the "fee" field.
func FeeNEQ(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNEQ(FieldFee, v))
}

// FeeIn applies the In predicate on the "fee" field.
func FeeIn(vs ...string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldIn(FieldFee, vs...))
}

// FeeNotIn applies the NotIn predicate on the "fee" field.
func FeeNotIn(vs ...string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNotIn(FieldFee, vs...))
}

// FeeGT applies the GT predicate on the "fee" field.
func FeeGT(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGT(FieldFee, v))
}

// FeeGTE applies the GTE predicate on the "fee" field.
func FeeGTE(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGTE(FieldFee, v))
}

// FeeLT applies the LT predicate on the "fee" field.
func FeeLT(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLT(FieldFee, v))
}

// FeeLTE applies the LTE predicate on the "fee" field.
func FeeLTE(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLTE(FieldFee, v))
}

// FeeContains applies the Contains predicate on the "fee" field.
func FeeContains(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldContains(FieldFee, v))
}

// FeeHasPrefix applies the HasPrefix predicate on the "fee" field.
func FeeHasPrefix(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldHasPrefix(FieldFee, v))
}

// FeeHasSuffix applies the HasSuffix predicate on the "fee" field.
func FeeHasSuffix(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldHasSuffix(FieldFee, v))
}

// FeeEqualFold applies the EqualFold predicate on the "fee" field.
func FeeEqualFold(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEqualFold(FieldFee, v))
}

// FeeContainsFold applies the ContainsFold predicate on the "fee" field.
func FeeContainsFold(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldContainsFold(FieldFee, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLTE(FieldTime, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Withdraw {
	return predicate.Withdraw(sql.FieldContainsFold(FieldAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Withdraw) predicate.Withdraw {
	return predicate.Withdraw(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Withdraw) predicate.Withdraw {
	return predicate.Withdraw(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Withdraw) predicate.Withdraw {
	return predicate.Withdraw(sql.NotPredicates(p))
}
