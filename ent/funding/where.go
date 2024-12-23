// Code generated by ent, DO NOT EDIT.

package funding

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldID, id))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldTime, v))
}

// Coin applies equality check predicate on the "coin" field. It's identical to CoinEQ.
func Coin(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldCoin, v))
}

// Usdc applies equality check predicate on the "usdc" field. It's identical to UsdcEQ.
func Usdc(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldUsdc, v))
}

// Szi applies equality check predicate on the "szi" field. It's identical to SziEQ.
func Szi(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldSzi, v))
}

// FundingRate applies equality check predicate on the "funding_rate" field. It's identical to FundingRateEQ.
func FundingRate(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldFundingRate, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldAddress, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldTime, v))
}

// CoinEQ applies the EQ predicate on the "coin" field.
func CoinEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldCoin, v))
}

// CoinNEQ applies the NEQ predicate on the "coin" field.
func CoinNEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldCoin, v))
}

// CoinIn applies the In predicate on the "coin" field.
func CoinIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldCoin, vs...))
}

// CoinNotIn applies the NotIn predicate on the "coin" field.
func CoinNotIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldCoin, vs...))
}

// CoinGT applies the GT predicate on the "coin" field.
func CoinGT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldCoin, v))
}

// CoinGTE applies the GTE predicate on the "coin" field.
func CoinGTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldCoin, v))
}

// CoinLT applies the LT predicate on the "coin" field.
func CoinLT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldCoin, v))
}

// CoinLTE applies the LTE predicate on the "coin" field.
func CoinLTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldCoin, v))
}

// CoinContains applies the Contains predicate on the "coin" field.
func CoinContains(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContains(FieldCoin, v))
}

// CoinHasPrefix applies the HasPrefix predicate on the "coin" field.
func CoinHasPrefix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasPrefix(FieldCoin, v))
}

// CoinHasSuffix applies the HasSuffix predicate on the "coin" field.
func CoinHasSuffix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasSuffix(FieldCoin, v))
}

// CoinEqualFold applies the EqualFold predicate on the "coin" field.
func CoinEqualFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEqualFold(FieldCoin, v))
}

// CoinContainsFold applies the ContainsFold predicate on the "coin" field.
func CoinContainsFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContainsFold(FieldCoin, v))
}

// UsdcEQ applies the EQ predicate on the "usdc" field.
func UsdcEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldUsdc, v))
}

// UsdcNEQ applies the NEQ predicate on the "usdc" field.
func UsdcNEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldUsdc, v))
}

// UsdcIn applies the In predicate on the "usdc" field.
func UsdcIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldUsdc, vs...))
}

// UsdcNotIn applies the NotIn predicate on the "usdc" field.
func UsdcNotIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldUsdc, vs...))
}

// UsdcGT applies the GT predicate on the "usdc" field.
func UsdcGT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldUsdc, v))
}

// UsdcGTE applies the GTE predicate on the "usdc" field.
func UsdcGTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldUsdc, v))
}

// UsdcLT applies the LT predicate on the "usdc" field.
func UsdcLT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldUsdc, v))
}

// UsdcLTE applies the LTE predicate on the "usdc" field.
func UsdcLTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldUsdc, v))
}

// UsdcContains applies the Contains predicate on the "usdc" field.
func UsdcContains(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContains(FieldUsdc, v))
}

// UsdcHasPrefix applies the HasPrefix predicate on the "usdc" field.
func UsdcHasPrefix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasPrefix(FieldUsdc, v))
}

// UsdcHasSuffix applies the HasSuffix predicate on the "usdc" field.
func UsdcHasSuffix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasSuffix(FieldUsdc, v))
}

// UsdcEqualFold applies the EqualFold predicate on the "usdc" field.
func UsdcEqualFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEqualFold(FieldUsdc, v))
}

// UsdcContainsFold applies the ContainsFold predicate on the "usdc" field.
func UsdcContainsFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContainsFold(FieldUsdc, v))
}

// SziEQ applies the EQ predicate on the "szi" field.
func SziEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldSzi, v))
}

// SziNEQ applies the NEQ predicate on the "szi" field.
func SziNEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldSzi, v))
}

// SziIn applies the In predicate on the "szi" field.
func SziIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldSzi, vs...))
}

// SziNotIn applies the NotIn predicate on the "szi" field.
func SziNotIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldSzi, vs...))
}

// SziGT applies the GT predicate on the "szi" field.
func SziGT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldSzi, v))
}

// SziGTE applies the GTE predicate on the "szi" field.
func SziGTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldSzi, v))
}

// SziLT applies the LT predicate on the "szi" field.
func SziLT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldSzi, v))
}

// SziLTE applies the LTE predicate on the "szi" field.
func SziLTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldSzi, v))
}

// SziContains applies the Contains predicate on the "szi" field.
func SziContains(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContains(FieldSzi, v))
}

// SziHasPrefix applies the HasPrefix predicate on the "szi" field.
func SziHasPrefix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasPrefix(FieldSzi, v))
}

// SziHasSuffix applies the HasSuffix predicate on the "szi" field.
func SziHasSuffix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasSuffix(FieldSzi, v))
}

// SziEqualFold applies the EqualFold predicate on the "szi" field.
func SziEqualFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEqualFold(FieldSzi, v))
}

// SziContainsFold applies the ContainsFold predicate on the "szi" field.
func SziContainsFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContainsFold(FieldSzi, v))
}

// FundingRateEQ applies the EQ predicate on the "funding_rate" field.
func FundingRateEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldFundingRate, v))
}

// FundingRateNEQ applies the NEQ predicate on the "funding_rate" field.
func FundingRateNEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldFundingRate, v))
}

// FundingRateIn applies the In predicate on the "funding_rate" field.
func FundingRateIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldFundingRate, vs...))
}

// FundingRateNotIn applies the NotIn predicate on the "funding_rate" field.
func FundingRateNotIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldFundingRate, vs...))
}

// FundingRateGT applies the GT predicate on the "funding_rate" field.
func FundingRateGT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldFundingRate, v))
}

// FundingRateGTE applies the GTE predicate on the "funding_rate" field.
func FundingRateGTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldFundingRate, v))
}

// FundingRateLT applies the LT predicate on the "funding_rate" field.
func FundingRateLT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldFundingRate, v))
}

// FundingRateLTE applies the LTE predicate on the "funding_rate" field.
func FundingRateLTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldFundingRate, v))
}

// FundingRateContains applies the Contains predicate on the "funding_rate" field.
func FundingRateContains(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContains(FieldFundingRate, v))
}

// FundingRateHasPrefix applies the HasPrefix predicate on the "funding_rate" field.
func FundingRateHasPrefix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasPrefix(FieldFundingRate, v))
}

// FundingRateHasSuffix applies the HasSuffix predicate on the "funding_rate" field.
func FundingRateHasSuffix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasSuffix(FieldFundingRate, v))
}

// FundingRateEqualFold applies the EqualFold predicate on the "funding_rate" field.
func FundingRateEqualFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEqualFold(FieldFundingRate, v))
}

// FundingRateContainsFold applies the ContainsFold predicate on the "funding_rate" field.
func FundingRateContainsFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContainsFold(FieldFundingRate, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Funding {
	return predicate.Funding(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Funding {
	return predicate.Funding(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Funding {
	return predicate.Funding(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Funding {
	return predicate.Funding(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Funding {
	return predicate.Funding(sql.FieldContainsFold(FieldAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Funding) predicate.Funding {
	return predicate.Funding(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Funding) predicate.Funding {
	return predicate.Funding(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Funding) predicate.Funding {
	return predicate.Funding(sql.NotPredicates(p))
}
