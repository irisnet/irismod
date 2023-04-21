package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global erc20 module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to modules/erc20 and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

	// AminoCdc is a amino codec created to support amino JSON compatible msgs.
	AminoCdc = codec.NewAminoCodec(amino)
)

const (
	// Amino names
	registerERC721Name = "irismod/erc721-converter/MsgRegisterERC721"
	registerDenomName  = "irismod/erc721-converter/MsgRegisterDenom"
	convertERC721Name  = "irismod/erc721-converter/MsgConvertERC721"
	convertNativeName  = "irismod/erc721-converter/MsgConvertNFT"
)

// NOTE: This is required for the GetSignBytes function
func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterERC721{},
		&MsgRegisterDenom{},
		&MsgConvertERC721{},
		&MsgConvertNFT{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

}

// RegisterLegacyAminoCodec registers the necessary x/erc20 interfaces and
// concrete types on the provided LegacyAmino codec. These types are used for
// Amino JSON serialization and EIP-712 compatibility.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterERC721{}, registerERC721Name, nil)
	cdc.RegisterConcrete(&MsgRegisterDenom{}, registerDenomName, nil)
	cdc.RegisterConcrete(&MsgConvertERC721{}, convertERC721Name, nil)
	cdc.RegisterConcrete(&MsgConvertNFT{}, convertNativeName, nil)
}
