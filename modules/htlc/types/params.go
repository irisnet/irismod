package types

import (
	"gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyAssetParams = []byte("AssetParams") // asset params key
)

// NewParams is the HTLC params constructor
func NewParams(assetParams []AssetParam) Params {
	return Params{
		AssetParams: assetParams,
	}
}

// ParamKeyTable returns the TypeTable for coinswap module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAssetParams, &p.AssetParams, validateAssetParams),
	}
}

// DefaultParams returns the default coinswap module parameters
func DefaultParams() Params {
	return Params{
		// TODO
		AssetParams: []AssetParam{},
	}
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// String returns a human readable string representation of the parameters.
func (p AssetParam) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// String returns a human readable string representation of the parameters.
func (p SupplyLimit) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate returns err if Params is invalid
func (p Params) Validate() error {
	// TODO
	return nil
}

func validateAssetParams(i interface{}) error {
	// TODO
	return nil
}
