package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const (
	TypeMsgSetDefaultRoyalty    = "set_default_royalty"
	TypeMsgSetTokenRoyalty      = "set_token_royalty"
	TypeMsgResetTokenRoyalty    = "reset_token_royalty"
	TypeMsgDeleteDefaultRoyalty = "delete_default_royalty"
)

var (
	_ sdk.Msg = &MsgSetDefaultRoyalty{}
	_ sdk.Msg = &MsgSetTokenRoyalty{}
	_ sdk.Msg = &MsgResetTokenRoyalty{}
	_ sdk.Msg = &MsgDeleteDefaultRoyalty{}
)

// NewMsgSetDefaultRoyalty is a constructor function for MsgSetName
func NewMsgSetDefaultRoyalty(denomID, receiver string, feeNumerator sdkmath.Uint, sender string) *MsgSetDefaultRoyalty {
	return &MsgSetDefaultRoyalty{
		DenomId:      denomID,
		Receiver:     receiver,
		FeeNumerator: feeNumerator,
		Sender:       sender,
	}
}

// Route Implements Msg
func (msg MsgSetDefaultRoyalty) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgSetDefaultRoyalty) Type() string { return TypeMsgSetDefaultRoyalty }

// ValidateBasic Implements Msg.
func (msg MsgSetDefaultRoyalty) ValidateBasic() error {
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateKeywords(msg.DenomId)
}

// GetSignBytes Implements Msg.
func (msg MsgSetDefaultRoyalty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSetDefaultRoyalty) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgSetTokenRoyalty is a constructor function for MsgSetName
func NewMsgSetTokenRoyalty(denomId, nftId, receiver string, feeNumerator sdkmath.Uint, sender string) *MsgSetTokenRoyalty {
	return &MsgSetTokenRoyalty{
		DenomId:      denomId,
		NftId:        nftId,
		Receiver:     receiver,
		FeeNumerator: feeNumerator,
		Sender:       sender,
	}
}

// Route Implements Msg
func (msg MsgSetTokenRoyalty) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgSetTokenRoyalty) Type() string { return TypeMsgSetTokenRoyalty }

// ValidateBasic Implements Msg.
func (msg MsgSetTokenRoyalty) ValidateBasic() error {
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateKeywords(msg.DenomId)
}

// GetSignBytes Implements Msg.
func (msg MsgSetTokenRoyalty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSetTokenRoyalty) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgResetTokenRoyalty is a constructor function for MsgSetName
func NewMsgResetTokenRoyalty(denomId, nftId, receiver string, feeNumerator sdkmath.Uint, sender string) *MsgResetTokenRoyalty {
	return &MsgResetTokenRoyalty{
		DenomId: denomId,
		NftId:   nftId,
		Sender:  sender,
	}
}

// Route Implements Msg
func (msg MsgResetTokenRoyalty) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgResetTokenRoyalty) Type() string { return TypeMsgResetTokenRoyalty }

// ValidateBasic Implements Msg.
func (msg MsgResetTokenRoyalty) ValidateBasic() error {
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateKeywords(msg.DenomId)
}

// GetSignBytes Implements Msg.
func (msg MsgResetTokenRoyalty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgResetTokenRoyalty) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgDeleteDefaultRoyalty is a constructor function for MsgSetName
func NewMsgDeleteDefaultRoyalty(denomId, nftId, receiver string, feeNumerator sdkmath.Uint, sender string) *MsgDeleteDefaultRoyalty {
	return &MsgDeleteDefaultRoyalty{
		DenomId: denomId,
		Sender:  sender,
	}
}

// Route Implements Msg
func (msg MsgDeleteDefaultRoyalty) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgDeleteDefaultRoyalty) Type() string { return TypeMsgDeleteDefaultRoyalty }

// ValidateBasic Implements Msg.
func (msg MsgDeleteDefaultRoyalty) ValidateBasic() error {
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return ValidateKeywords(msg.DenomId)
}

// GetSignBytes Implements Msg.
func (msg MsgDeleteDefaultRoyalty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgDeleteDefaultRoyalty) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
