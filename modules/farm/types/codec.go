package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "irismod/farm/MsgCreatePool", nil)
	cdc.RegisterConcrete(&MsgDestroyPool{}, "irismod/farm/MsgDestroyPool", nil)
	cdc.RegisterConcrete(&MsgAppendReward{}, "irismod/farm/MsgAppendReward", nil)
	cdc.RegisterConcrete(&MsgStake{}, "irismod/farm/MsgStake", nil)
	cdc.RegisterConcrete(&MsgUnstake{}, "irismod/farm/MsgUnstake", nil)
	cdc.RegisterConcrete(&MsgHarvest{}, "irismod/farm/MsgHarvest", nil)
}

// RegisterInterfaces registers the interface
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
		&MsgDestroyPool{},
		&MsgAppendReward{},
		&MsgStake{},
		&MsgUnstake{},
		&MsgHarvest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
