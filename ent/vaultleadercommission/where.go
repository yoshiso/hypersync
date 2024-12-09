// Code generated by ent, DO NOT EDIT.

package vaultleadercommission

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLTE(FieldID, id))
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldUser, v))
}

// Usdc applies equality check predicate on the "usdc" field. It's identical to UsdcEQ.
func Usdc(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldUsdc, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldTime, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldAddress, v))
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldUser, v))
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNEQ(FieldUser, v))
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldIn(FieldUser, vs...))
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNotIn(FieldUser, vs...))
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGT(FieldUser, v))
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGTE(FieldUser, v))
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLT(FieldUser, v))
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLTE(FieldUser, v))
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldContains(FieldUser, v))
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldHasPrefix(FieldUser, v))
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldHasSuffix(FieldUser, v))
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEqualFold(FieldUser, v))
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldContainsFold(FieldUser, v))
}

// UsdcEQ applies the EQ predicate on the "usdc" field.
func UsdcEQ(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldUsdc, v))
}

// UsdcNEQ applies the NEQ predicate on the "usdc" field.
func UsdcNEQ(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNEQ(FieldUsdc, v))
}

// UsdcIn applies the In predicate on the "usdc" field.
func UsdcIn(vs ...string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldIn(FieldUsdc, vs...))
}

// UsdcNotIn applies the NotIn predicate on the "usdc" field.
func UsdcNotIn(vs ...string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNotIn(FieldUsdc, vs...))
}

// UsdcGT applies the GT predicate on the "usdc" field.
func UsdcGT(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGT(FieldUsdc, v))
}

// UsdcGTE applies the GTE predicate on the "usdc" field.
func UsdcGTE(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGTE(FieldUsdc, v))
}

// UsdcLT applies the LT predicate on the "usdc" field.
func UsdcLT(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLT(FieldUsdc, v))
}

// UsdcLTE applies the LTE predicate on the "usdc" field.
func UsdcLTE(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLTE(FieldUsdc, v))
}

// UsdcContains applies the Contains predicate on the "usdc" field.
func UsdcContains(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldContains(FieldUsdc, v))
}

// UsdcHasPrefix applies the HasPrefix predicate on the "usdc" field.
func UsdcHasPrefix(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldHasPrefix(FieldUsdc, v))
}

// UsdcHasSuffix applies the HasSuffix predicate on the "usdc" field.
func UsdcHasSuffix(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldHasSuffix(FieldUsdc, v))
}

// UsdcEqualFold applies the EqualFold predicate on the "usdc" field.
func UsdcEqualFold(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEqualFold(FieldUsdc, v))
}

// UsdcContainsFold applies the ContainsFold predicate on the "usdc" field.
func UsdcContainsFold(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldContainsFold(FieldUsdc, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLTE(FieldTime, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.FieldContainsFold(FieldAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VaultLeaderCommission) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VaultLeaderCommission) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VaultLeaderCommission) predicate.VaultLeaderCommission {
	return predicate.VaultLeaderCommission(sql.NotPredicates(p))
}
