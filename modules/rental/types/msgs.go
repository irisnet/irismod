package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

const (
	TypeMsgSetUser = "set_user"
)

var (
	_ sdk.Msg = &MsgSetUser{}
)

func NewMsgSetUser(classId, nftId, user string,
	expires uint64, sender string) *MsgSetUser {
	return &MsgSetUser{
		ClassId: classId,
		NftId:   nftId,
		User:    user,
		Expires: expires,
		Sender:  sender,
	}
}

// Route Implements LegacyMsg
func (m MsgSetUser) Route() string { return RouterKey }

// Type Implements LegacyMsg
func (m MsgSetUser) Type() string { return TypeMsgSetUser }

// GetSignBytes Implements LegacyMsg
func (m MsgSetUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the Msg.ValidateBasic method.
func (m MsgSetUser) ValidateBasic() error {
	if err := nfttypes.ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.Wrapf(ErrInvalidClassID, "Invalid class id (%s)", m.ClassId)
	}

	if err := nfttypes.ValidateNFTID(m.NftId); err != nil {
		return sdkerrors.Wrapf(ErrInvalidNftID, "Invalid nft id (%s)", m.NftId)
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	_, err = sdk.AccAddressFromBech32(m.User)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid user address (%s)", m.User)

	}

	if m.Expires == 0 {
		return sdkerrors.Wrapf(ErrInvalidExpires, "Invalid expires (%d)", m.Expires)
	}

	return nil
}

// GetSigners implements the Msg.GetSigners method.
func (m MsgSetUser) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}
