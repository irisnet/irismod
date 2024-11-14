package v1

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

// NewParams constructs a new Params instance
func NewParams(tokenTaxRate math.LegacyDec, issueTokenBaseFee sdk.Coin,
	mintTokenFeeRatio math.LegacyDec, enableErc20 bool, beacon string,
) Params {
	return Params{
		TokenTaxRate:      tokenTaxRate,
		IssueTokenBaseFee: issueTokenBaseFee,
		MintTokenFeeRatio: mintTokenFeeRatio,
		EnableErc20:       enableErc20,
		Beacon:            beacon,
	}
}

// DefaultParams return the default params
func DefaultParams() Params {
	defaultToken := GetNativeToken()
	return Params{
		TokenTaxRate:      math.LegacyNewDecWithPrec(4, 1), // 0.4 (40%)
		IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, math.NewInt(60000)),
		MintTokenFeeRatio: math.LegacyNewDecWithPrec(1, 1), // 0.1 (10%)
		EnableErc20:       true,
	}
}

// Validate validates the given params
func (p Params) Validate() error {
	if err := validateTaxRate(p.TokenTaxRate); err != nil {
		return err
	}
	if err := validateMintTokenFeeRatio(p.MintTokenFeeRatio); err != nil {
		return err
	}
	if err := validateIssueTokenBaseFee(p.IssueTokenBaseFee); err != nil {
		return err
	}
	return validateBeacon(p.Beacon)
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

func validateBeacon(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(v) == 0 {
		return nil
	}
	if !common.IsHexAddress(v) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "beacon expecting a hex address, got %s", v)
	}
	return nil
}
