package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sigs.k8s.io/yaml"
)

// NewParams creates a new Params instance
func NewParams(createPoolFee sdk.Coin, maxRewardCategories uint32, taxRate math.LegacyDec) Params {
	return Params{
		PoolCreationFee:     createPoolFee,
		TaxRate:             taxRate,
		MaxRewardCategories: maxRewardCategories,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(5000)),
		2,
		math.LegacyNewDecWithPrec(4, 1),
	)
}

// Validate validates a set of params
func (p Params) Validate() error {
	return validatePoolCreationFee(p.PoolCreationFee)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validatePoolCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid minimum deposit: %s", v)
	}
	return nil
}

func validateTaxRate(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.GT(math.LegacyZeroDec()) || !v.LT(math.LegacyOneDec()) {
		return fmt.Errorf("tax rate must be positive and less than 1: %s", v.String())
	}
	return nil
}

func validateMaxRewardCategories(i interface{}) error { return nil }
