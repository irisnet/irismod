package v1beta1

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// parameter keys
var (
	KeyTokenTaxRate      = []byte("TokenTaxRate")
	KeyIssueTokenBaseFee = []byte("IssueTokenBaseFee")
	KeyMintTokenFeeRatio = []byte("MintTokenFeeRatio")
)

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyTokenTaxRate, &p.TokenTaxRate, validateTaxRate),
		paramtypes.NewParamSetPair(KeyIssueTokenBaseFee, &p.IssueTokenBaseFee, validateIssueTokenBaseFee),
		paramtypes.NewParamSetPair(KeyMintTokenFeeRatio, &p.MintTokenFeeRatio, validateMintTokenFeeRatio),
	}
}

// NewParams constructs a new Params instance
func NewParams(tokenTaxRate math.LegacyDec, issueTokenBaseFee sdk.Coin,
	mintTokenFeeRatio math.LegacyDec,
) Params {
	return Params{
		TokenTaxRate:      tokenTaxRate,
		IssueTokenBaseFee: issueTokenBaseFee,
		MintTokenFeeRatio: mintTokenFeeRatio,
	}
}

// ParamKeyTable returns the TypeTable for the token module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams return the default params
func DefaultParams() Params {
	defaultToken := GetNativeToken()
	return Params{
		TokenTaxRate:      math.LegacyNewDecWithPrec(4, 1), // 0.4 (40%)
		IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, math.NewInt(60000)),
		MintTokenFeeRatio: math.LegacyNewDecWithPrec(1, 1), // 0.1 (10%)
	}
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ValidateParams validates the given params
func ValidateParams(p Params) error {
	if err := validateTaxRate(p.TokenTaxRate); err != nil {
		return err
	}
	if err := validateMintTokenFeeRatio(p.MintTokenFeeRatio); err != nil {
		return err
	}
	if err := validateIssueTokenBaseFee(p.IssueTokenBaseFee); err != nil {
		return err
	}

	return nil
}

func validateTaxRate(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.GT(math.LegacyNewDec(1)) || v.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("token tax rate [%s] should be between [0, 1]", v.String())
	}
	return nil
}

func validateMintTokenFeeRatio(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.GT(math.LegacyNewDec(1)) || v.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("fee ratio for minting tokens [%s] should be between [0, 1]", v.String())
	}
	return nil
}

func validateIssueTokenBaseFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("base fee for issuing token should not be negative")
	}
	return nil
}
