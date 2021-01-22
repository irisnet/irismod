package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateInput verifies whether the  parameters are legal
func ValidateInput(input Input) error {
	if err := input.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid input (%s)", err.Error())
	}

	if strings.HasPrefix(input.Coin.Denom, FormatUniABSPrefix) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid input denom,shoule not be begin with (%s)", FormatUniABSPrefix)
	}

	if _, err := sdk.AccAddressFromBech32(input.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid input address (%s)", err)
	}
	return nil
}

// ValidateOutput verifies whether the  parameters are legal
func ValidateOutput(output Output) error {
	if err := output.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid output (%s)", err.Error())
	}

	if strings.HasPrefix(output.Coin.Denom, FormatUniABSPrefix) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid output denom,shoule not be begin with (%s)", FormatUniABSPrefix)
	}

	if _, err := sdk.AccAddressFromBech32(output.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid output address (%s)", err)
	}
	return nil
}

// ValidateDeadline verifies whether the  parameters are legal
func ValidateDeadline(deadline int64) error {
	if deadline <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deadline %d must be greater than 0", deadline))
	}
	return nil
}

// ValidateMaxToken verifies whether the  parameters are legal
func ValidateMaxToken(maxToken sdk.Coin) error {
	if err := maxToken.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid maxToken (%s)", err.Error())
	}

	if maxToken.Denom == StandardDenom {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("max token must not be standard token: %s", StandardDenom))
	}

	if strings.HasPrefix(maxToken.Denom, FormatUniABSPrefix) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max token must be non-liquidity token")
	}
	return nil
}

// ValidateExactStandardAmt verifies whether the  parameters are legal
func ValidateExactStandardAmt(standardAmt sdk.Int) error {
	if !standardAmt.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "standard token amount must be positive")
	}
	return nil
}

// ValidateMinLiquidity verifies whether the  parameters are legal
func ValidateMinLiquidity(minLiquidity sdk.Int) error {
	if minLiquidity.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "minimum liquidity can not be negative")
	}
	return nil
}

// ValidateMinToken verifies whether the  parameters are legal
func ValidateMinToken(minToken sdk.Int) error {
	if minToken.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "minimum token amount can not be negative")
	}
	return nil
}

// ValidateWithdrawLiquidity verifies whether the  parameters are legal
func ValidateWithdrawLiquidity(liquidity sdk.Coin) error {
	if err := liquidity.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid withdrawLiquidity (%s)", err.Error())
	}

	if err := CheckUniDenom(liquidity.Denom); err != nil {
		return err
	}
	return nil
}

// ValidateMinStandardAmt verifies whether the  parameters are legal
func ValidateMinStandardAmt(minStandardAmt sdk.Int) error {
	if minStandardAmt.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("minimum standard token amount %s can not be negative", minStandardAmt.String()))
	}
	return nil
}
