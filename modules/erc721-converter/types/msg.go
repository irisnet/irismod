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
)

const (
	TypeMsgConvertNFT    = "convert_NFT"
	TypeMsgConvertERC721 = "convert_ERC721"
)

// NewConvertNFTMsg creates a new instance of MsgConvertNFT
func NewConvertNFTMsg(classId string, tokenId string, receiver, sender sdk.AccAddress) *MsgConvertNFT {
	return &MsgConvertNFT{
		ClassId:  classId,
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
func NewConvertERC721Msg(contractAddress string, tokenId sdk.Int, receiver, sender sdk.AccAddress) *MsgConvertERC721 {
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
