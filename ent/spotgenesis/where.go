// Code generated by ent, DO NOT EDIT.

package spotgenesis

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLTE(FieldID, id))
}

// Coin applies equality check predicate on the "coin" field. It's identical to CoinEQ.
func Coin(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldCoin, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldAmount, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldTime, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldAddress, v))
}

// CoinEQ applies the EQ predicate on the "coin" field.
func CoinEQ(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldCoin, v))
}

// CoinNEQ applies the NEQ predicate on the "coin" field.
func CoinNEQ(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNEQ(FieldCoin, v))
}

// CoinIn applies the In predicate on the "coin" field.
func CoinIn(vs ...string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldIn(FieldCoin, vs...))
}

// CoinNotIn applies the NotIn predicate on the "coin" field.
func CoinNotIn(vs ...string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNotIn(FieldCoin, vs...))
}

// CoinGT applies the GT predicate on the "coin" field.
func CoinGT(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGT(FieldCoin, v))
}

// CoinGTE applies the GTE predicate on the "coin" field.
func CoinGTE(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGTE(FieldCoin, v))
}

// CoinLT applies the LT predicate on the "coin" field.
func CoinLT(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLT(FieldCoin, v))
}

// CoinLTE applies the LTE predicate on the "coin" field.
func CoinLTE(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLTE(FieldCoin, v))
}

// CoinContains applies the Contains predicate on the "coin" field.
func CoinContains(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldContains(FieldCoin, v))
}

// CoinHasPrefix applies the HasPrefix predicate on the "coin" field.
func CoinHasPrefix(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldHasPrefix(FieldCoin, v))
}

// CoinHasSuffix applies the HasSuffix predicate on the "coin" field.
func CoinHasSuffix(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldHasSuffix(FieldCoin, v))
}

// CoinEqualFold applies the EqualFold predicate on the "coin" field.
func CoinEqualFold(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEqualFold(FieldCoin, v))
}

// CoinContainsFold applies the ContainsFold predicate on the "coin" field.
func CoinContainsFold(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldContainsFold(FieldCoin, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLTE(FieldAmount, v))
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldContains(FieldAmount, v))
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldHasPrefix(FieldAmount, v))
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldHasSuffix(FieldAmount, v))
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEqualFold(FieldAmount, v))
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldContainsFold(FieldAmount, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLTE(FieldTime, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.FieldContainsFold(FieldAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SpotGenesis) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SpotGenesis) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SpotGenesis) predicate.SpotGenesis {
	return predicate.SpotGenesis(sql.NotPredicates(p))
}
