package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Farm params default values
var (
	DefaultCreatePoolFee = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(5000)) // 5000stake
)

// Keys for parameter access
// nolint
var (
	KeyCreatePoolFee = []byte("CreatePoolFee")
)

// NewParams creates a new Params instance
func NewParams(createPoolFee sdk.Coin) Params {
	return Params{
		CreatePoolFee: createPoolFee,
	}
}

// ParamSetPairs implements paramstypes.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyCreatePoolFee, &p.CreatePoolFee, validateCreatePoolFee),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(DefaultCreatePoolFee)
}

// Validate validates a set of params
func (p Params) Validate() error {
	return validateCreatePoolFee(p.CreatePoolFee)
}

func validateCreatePoolFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return fmt.Errorf("invalid minimum deposit: %s", v)
	}
	return nil
}
