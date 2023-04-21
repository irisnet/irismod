package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = &MsgConvertNFT{}
	_ sdk.Msg = &MsgConvertERC721{}
	_ sdk.Msg = &MsgRegisterDenom{}
	_ sdk.Msg = &MsgRegisterERC721{}
)

const (
	TypeMsgRegisterDenom  = "register_denom"
	TypeMsgRegisterERC721 = "register_ERC721"
	TypeMsgConvertNFT     = "convert_NFT"
	TypeMsgConvertERC721  = "convert_ERC721"
)

// NewRegisterDenomMsg creates a new instance of MsgRegisterDenom
func NewRegisterDenomMsg(denomId string, sender sdk.AccAddress) *MsgRegisterDenom {
	return &MsgRegisterDenom{
		DenomId: denomId,
		Sender:  sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgRegisterDenom) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterDenom) Type() string { return TypeMsgRegisterDenom }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterDenom) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// NewRegisterERC721Msg creates a new instance of MsgRegisterERC721
func NewRegisterERC721Msg(contractAddress string, sender sdk.AccAddress) *MsgRegisterERC721 {
	return &MsgRegisterERC721{
		ContractAddress: contractAddress,
		Sender:          sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgRegisterERC721) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterERC721) Type() string { return TypeMsgRegisterERC721 }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterERC721) ValidateBasic() error {
	if !common.IsHexAddress(msg.ContractAddress) {
		return errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid contract hex address '%s'", msg.ContractAddress)
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterERC721) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterERC721) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// NewConvertNFTMsg creates a new instance of MsgConvertNFT
func NewConvertNFTMsg(denomId string, tokenId string, receiver, sender sdk.AccAddress) *MsgConvertNFT {
	return &MsgConvertNFT{
		DenomId:  denomId,
		TokenId:  tokenId,
		Receiver: receiver.String(),
		Sender:   sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgConvertNFT) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertNFT) Type() string { return TypeMsgConvertNFT }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgConvertNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertNFT) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// NewConvertERC721Msg creates a new instance of MsgConvertERC721
func NewConvertERC721Msg(contractAddress string, tokenId string, receiver, sender sdk.AccAddress) *MsgConvertERC721 {
	return &MsgConvertERC721{
		ContractAddress: contractAddress,
		TokenId:         tokenId,
		Receiver:        receiver.String(),
		Sender:          sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgConvertERC721) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertERC721) Type() string { return TypeMsgConvertERC721 }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertERC721) ValidateBasic() error {
	if !common.IsHexAddress(msg.ContractAddress) {
		return errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid contract hex address '%s'", msg.ContractAddress)
	}
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgConvertERC721) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertERC721) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}
