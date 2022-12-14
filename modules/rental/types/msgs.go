package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgSetUser = "set_user"
)

var (
	_ sdk.Msg = &MsgSetUser{}
)

func NewMsgSetUser(classId, nftId, renter string,
	expire uint64, sender string) *MsgSetUser {
	return &MsgSetUser{
		ClassId: classId,
		NftId:   nftId,
		Renter:  renter,
		Expire:  expire,
		Sender:  sender,
	}
}

func (MsgSetUser) Route() string { return RouterKey }

func (MsgSetUser) Type() string { return TypeMsgSetUser }

func (MsgSetUser) ValidateBasic() error {
	panic("Fixme")
}

func (MsgSetUser) GetSignBytes() []byte {
	panic("Fixme")
}

func (MsgSetUser) GetSigners() []sdk.AccAddress {
	panic("Fixme")
}
