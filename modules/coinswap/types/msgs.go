package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSwapOrder{}
	_ sdk.Msg = &MsgAddLiquidity{}
	_ sdk.Msg = &MsgRemoveLiquidity{}
)

const (
	// FormatUniABSPrefix define the prefix of liquidity token
	FormatUniABSPrefix = "swap/"
	// FormatUniDenom define the name of liquidity token
	FormatUniDenom = "swap/%s"

	// TypeMsgAddLiquidity define the event type of the MsgAddLiquidity
	TypeMsgAddLiquidity = "add_liquidity"
	// TypeMsgRemoveLiquidity define the event type of the MsgRemoveLiquidity
	TypeMsgRemoveLiquidity = "remove_liquidity"
	// TypeMsgSwapOrder define the event type of the MsgSwapOrder
	TypeMsgSwapOrder = "swap_order"
)

/* --------------------------------------------------------------------------- */
// MsgSwapOrder
/* --------------------------------------------------------------------------- */

// MsgSwapOrder - struct for swapping a coin
// Input and Output can either be exact or calculated.
// An exact coin has the senders desired buy or sell amount.
// A calculated coin has the desired denomination and bounded amount
// the sender is willing to buy or sell in this order.

// NewMsgSwapOrder creates a new MsgSwapOrder object.
func NewMsgSwapOrder(
	input Input,
	output Output,
	deadline int64,
	isBuyOrder bool,
) *MsgSwapOrder {
	return &MsgSwapOrder{
		Input:      input,
		Output:     output,
		Deadline:   deadline,
		IsBuyOrder: isBuyOrder,
	}
}

// Route implements Msg.
func (msg MsgSwapOrder) Route() string { return RouterKey }

// Type implements Msg.
func (msg MsgSwapOrder) Type() string { return TypeMsgSwapOrder }

// ValidateBasic implements Msg.
func (msg MsgSwapOrder) ValidateBasic() error {
	if err := msg.Input.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid input (%s)", err.Error())
	}

	if strings.HasPrefix(msg.Input.Coin.Denom, FormatUniABSPrefix) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("unsupported input coin type: %s", msg.Input.Coin.String()))
	}

	if err := msg.Output.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid output (%s)", err.Error())
	}

	if strings.HasPrefix(msg.Output.Coin.Denom, FormatUniABSPrefix) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("unsupported output coin type: %s", msg.Output.Coin.String()))
	}

	if msg.Input.Coin.Denom == msg.Output.Coin.Denom {
		return sdkerrors.Wrap(ErrEqualDenom, "invalid swap")
	}

	if msg.Deadline <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deadline %d must be greater than 0", msg.Deadline))
	}

	if _, err := sdk.AccAddressFromBech32(msg.Input.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid input address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Output.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid output address (%s)", err)
	}
	return nil
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgSwapOrder) Normalize() *MsgSwapOrder {
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgSwapOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgSwapOrder) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Input.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

/* --------------------------------------------------------------------------- */
// MsgAddLiquidity
/* --------------------------------------------------------------------------- */

// NewMsgAddLiquidity creates a new MsgAddLiquidity object.
func NewMsgAddLiquidity(
	maxToken sdk.Coin,
	exactStandardAmt sdk.Int,
	minLiquidity sdk.Int,
	deadline int64,
	sender string,
) *MsgAddLiquidity {
	return &MsgAddLiquidity{
		MaxToken:         maxToken,
		ExactStandardAmt: exactStandardAmt,
		MinLiquidity:     minLiquidity,
		Deadline:         deadline,
		Sender:           sender,
	}
}

// Route implements Msg.
func (msg MsgAddLiquidity) Route() string { return RouterKey }

// Type implements Msg.
func (msg MsgAddLiquidity) Type() string { return TypeMsgAddLiquidity }

// ValidateBasic implements Msg.
func (msg MsgAddLiquidity) ValidateBasic() error {
	if err := msg.MaxToken.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid maxToken (%s)", err.Error())
	}

	if msg.MaxToken.Denom == StandardDenom {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("max token must not be standard token: %s", StandardDenom))
	}

	if strings.HasPrefix(msg.MaxToken.Denom, FormatUniABSPrefix) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "max token must be non-liquidity token")
	}

	if !msg.ExactStandardAmt.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "standard token amount must be positive")
	}

	if msg.MinLiquidity.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "minimum liquidity can not be negative")
	}

	if msg.Deadline <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deadline %d must be greater than 0", msg.Deadline))
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgAddLiquidity) Normalize() *MsgAddLiquidity {
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgAddLiquidity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgAddLiquidity) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

/* --------------------------------------------------------------------------- */
// MsgRemoveLiquidity
/* --------------------------------------------------------------------------- */

// NewMsgRemoveLiquidity creates a new MsgRemoveLiquidity object
func NewMsgRemoveLiquidity(
	minToken sdk.Int,
	withdrawLiquidity sdk.Coin,
	minStandardAmt sdk.Int,
	deadline int64,
	sender string,
) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		MinToken:          minToken,
		WithdrawLiquidity: withdrawLiquidity,
		MinStandardAmt:    minStandardAmt,
		Deadline:          deadline,
		Sender:            sender,
	}
}

// Route implements Msg.
func (msg MsgRemoveLiquidity) Route() string { return RouterKey }

// Type implements Msg.
func (msg MsgRemoveLiquidity) Type() string { return TypeMsgRemoveLiquidity }

// ValidateBasic implements Msg.
func (msg MsgRemoveLiquidity) ValidateBasic() error {
	if msg.MinToken.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "minimum token amount can not be negative")
	}

	if err := msg.WithdrawLiquidity.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid withdrawLiquidity (%s)", err.Error())
	}

	if err := CheckUniDenom(msg.WithdrawLiquidity.Denom); err != nil {
		return err
	}

	if msg.MinStandardAmt.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("minimum standard token amount %s can not be negative", msg.MinStandardAmt.String()))
	}

	if msg.Deadline <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deadline %d must be greater than 0", msg.Deadline))
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// Normalize return a string with spaces removed and lowercase
func (msg *MsgRemoveLiquidity) Normalize() *MsgRemoveLiquidity {
	return msg
}

// GetSignBytes implements Msg.
func (msg MsgRemoveLiquidity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements Msg.
func (msg MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
