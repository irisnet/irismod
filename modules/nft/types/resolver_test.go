package types

import (
	"encoding/json"
	"testing"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestClassMetadataResolverEncodeAndDecode(t *testing.T) {
	creator, err := sdk.AccAddressFromHexUnsafe(crypto.AddressHash([]byte("test_consumer")).String())
	require.NoError(t, err, "AccAddressFromHexUnsafe failed")

	// dataMap := map[string]MediaField{
	// 	"irismod:key1": {Value: "value1"},
	// 	"irismod:key2": {Value: "value2"},
	// }
	// dataBytes, err := json.Marshal(dataMap)
	// require.NoError(t, err, " dataMap json.Marshal failed")
	// t.Logf("%s", dataBytes)

	denomMetadata := DenomMetadata{
		Creator:          creator.String(),
		Schema:           "{}",
		MintRestricted:   true,
		UpdateRestricted: true,
		Data:             "{\"key1\":\"value1\",\"key2\":\"value2\"}",
	}

	bz, err := json.Marshal(denomMetadata)
	require.NoError(t, err, " denomMetadata json.Marshal failed")
	t.Logf("%s", bz)

	any, err := codectypes.NewAnyWithValue(&denomMetadata)
	require.NoError(t, err, " denomMetadata codectypes.NewAnyWithValue failed")

	getModuleAddress := func(_ string) sdk.AccAddress {
		return creator
	}
	class := nft.Class{
		Id:          "cat",
		Name:        "kitty",
		Symbol:      "symbol",
		Description: "digital cat",
		Uri:         "uri",
		UriHash:     "uri_hash",
		Data:        any,
	}

	cdc := GetEncoding()
	resolver := NewClassResolver(cdc, getModuleAddress)
	result, err := resolver.Encode(class)
	require.NoError(t, err, " class resolver.Encode failed")
	t.Log(result)

	expClass, err := resolver.Decode(class.Id, class.Uri, result)
	require.NoError(t, err, " class resolver.Decode failed")

	exp, err := cdc.MarshalInterfaceJSON(&class)
	require.NoError(t, err, " class resolver.Decode failed")
	t.Logf("%s", exp)

	act, err := cdc.MarshalInterfaceJSON(&expClass)
	require.NoError(t, err, " class resolver.Decode failed")
	t.Logf("%s", act)

	require.Equal(t, act, exp, "not equal")
}

func GetEncoding() codec.Codec {
	interfaceRegistry := types.NewInterfaceRegistry()
	interfaceRegistry.RegisterImplementations(
		(*proto.Message)(nil),
		&nft.Class{},
		&nft.NFT{},
		&DenomMetadata{},
		&NFTMetadata{},
	)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return marshaler
}
