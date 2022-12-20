package types

import (
	"encoding/json"
	"testing"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestClassMetadataResolverMarshal(t *testing.T) {
	creator, err := sdk.AccAddressFromHexUnsafe(crypto.AddressHash([]byte("test_consumer")).String())
	require.NoError(t, err, "AccAddressFromHexUnsafe failed")

	customData := map[string]MediaField{
		"key1": {Value: "value1"},
		"key2": {Value: "value2"},
	}
	customDataBz, _ := json.Marshal(customData)
	denomMetadata := DenomMetadata{
		Creator:          creator.String(),
		Schema:           "{}",
		MintRestricted:   true,
		UpdateRestricted: true,
		Data:             string(customDataBz),
	}

	bz, err := json.Marshal(denomMetadata)
	require.NoError(t, err, " denomMetadata json.Marshal failed")
	t.Logf("%s", bz)

	any, err := codectypes.NewAnyWithValue(&denomMetadata)
	require.NoError(t, err, " denomMetadata codectypes.NewAnyWithValue failed")

	getModuleAddress := func(_ string) sdk.AccAddress {
		return creator
	}

	resolver := NewClassMetadataResolver(GetEncoding(), getModuleAddress)
	result, err := resolver.Encode(any)
	require.NoError(t, err, " denomMetadata resolver.Marshal failed")
	t.Log(result)

	expClass, err := resolver.Decode(result)
	require.NoError(t, err, " denomMetadata resolver.Decode failed")
	require.True(t, expClass.Equal(any), "not equal")
}

func GetEncoding() codec.Codec {
	interfaceRegistry := types.NewInterfaceRegistry()
	interfaceRegistry.RegisterImplementations(
		(*proto.Message)(nil),
		&DenomMetadata{},
		&NFTMetadata{},
	)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return marshaler
}
