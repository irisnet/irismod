package types

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

const (
	Namespace          = "irismod:"
	KeyMediaFieldValue = "value"
)

var (
	MintRestrictedFieldKey   = fmt.Sprintf("%s%s", Namespace, "mint_restricted")
	UpdateRestrictedFieldKey = fmt.Sprintf("%s%s", Namespace, "update_restricted")
	CreatorFieldKey          = fmt.Sprintf("%s%s", Namespace, "creator")
	SchemaFieldKey           = fmt.Sprintf("%s%s", Namespace, "schema")
	NameFieldKey             = fmt.Sprintf("%s%s", Namespace, "name")
)

type (
	ClassMetadataResolver struct {
		cdc              codec.Codec
		getModuleAddress func(string) sdk.AccAddress
	}
	TokenMetadataResolver struct{ cdc codec.Codec }
	Type                  interface{ bool | string }
	MediaField            struct {
		Value interface{} `json:"value"`
		Mime  string      `json:"mime,omitempty"`
	}
)

func NewClassMetadataResolver(cdc codec.Codec, getModuleAddress func(string) sdk.AccAddress) ClassMetadataResolver {
	return ClassMetadataResolver{
		cdc:              cdc,
		getModuleAddress: getModuleAddress,
	}
}

// Encode encode any into the metadata data format defined by ics721
func (cmr ClassMetadataResolver) Encode(any *codectypes.Any) (string, error) {
	var message proto.Message
	if err := cmr.cdc.UnpackAny(any, &message); err != nil {
		return "", err
	}

	denomMetadata, ok := message.(*DenomMetadata)
	if !ok {
		return "", errors.New("unsupport classMetadata")
	}

	kvals := make(map[string]interface{})
	if err := json.Unmarshal([]byte(denomMetadata.Data), &kvals); err != nil {
		//when classData is not a legal json, there is no need to parse the data
		return base64.RawStdEncoding.EncodeToString([]byte(denomMetadata.Data)), nil
	}

	kvals[MintRestrictedFieldKey] = MediaField{Value: denomMetadata.MintRestricted}
	kvals[UpdateRestrictedFieldKey] = MediaField{Value: denomMetadata.UpdateRestricted}
	kvals[CreatorFieldKey] = MediaField{Value: denomMetadata.Creator}
	kvals[SchemaFieldKey] = MediaField{Value: denomMetadata.Schema}
	data, err := json.Marshal(kvals)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil
}

func (cmr ClassMetadataResolver) Decode(classInfo string) (*codectypes.Any, error) {
	classInfoBz, err := base64.RawStdEncoding.DecodeString(classInfo)
	if err != nil {
		return nil, err
	}

	var (
		mintRestricted   = true
		updateRestricted = true
		schema           = ""
		creator          = cmr.getModuleAddress(ModuleName).String()
	)

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(classInfoBz, &dataMap); err != nil {
		return codectypes.NewAnyWithValue(&DenomMetadata{
			Creator:          creator,
			MintRestricted:   mintRestricted,
			UpdateRestricted: updateRestricted,
			Data:             string(classInfoBz),
		})
	}

	if v, ok := dataMap[MintRestrictedFieldKey]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vBool, ok := vMap[KeyMediaFieldValue].(bool); ok {
				mintRestricted = vBool
				delete(dataMap, MintRestrictedFieldKey)
			}
		}
	}

	if v, ok := dataMap[UpdateRestrictedFieldKey]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vBool, ok := vMap[KeyMediaFieldValue].(bool); ok {
				updateRestricted = vBool
				delete(dataMap, UpdateRestrictedFieldKey)
			}
		}
	}

	if v, ok := dataMap[CreatorFieldKey]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				creator = vStr
				delete(dataMap, CreatorFieldKey)
			}
		}
	}

	if v, ok := dataMap[SchemaFieldKey]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				schema = vStr
				delete(dataMap, SchemaFieldKey)
			}
		}
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return nil, err
	}

	return codectypes.NewAnyWithValue(&DenomMetadata{
		Creator:          creator,
		Schema:           schema,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		Data:             string(data),
	})
}

func NewTokenMetadataResolver(cdc codec.Codec) TokenMetadataResolver {
	return TokenMetadataResolver{
		cdc: cdc,
	}
}

func (cmr TokenMetadataResolver) Encode(any *codectypes.Any) (string, error) {
	var message proto.Message
	if err := cmr.cdc.UnpackAny(any, &message); err != nil {
		return "", err
	}

	nftMetadata, ok := message.(*NFTMetadata)
	if !ok {
		return "", errors.New("unsupport classMetadata")
	}
	kvals := make(map[string]interface{})
	if err := json.Unmarshal([]byte(nftMetadata.Data), &kvals); err != nil {
		//when nftMetadata is not a legal json, there is no need to parse the data
		return base64.RawStdEncoding.EncodeToString([]byte(nftMetadata.Data)), nil
	}

	kvals[NameFieldKey] = MediaField{Value: nftMetadata.Name}
	data, err := json.Marshal(kvals)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil
}

func (cmr TokenMetadataResolver) Decode(tokenInfo string) (*codectypes.Any, error) {
	tokenInfoBz, err := base64.RawStdEncoding.DecodeString(tokenInfo)
	if err != nil {
		return nil, err
	}

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(tokenInfoBz, &dataMap); err != nil {
		return codectypes.NewAnyWithValue(&NFTMetadata{
			Data: string(tokenInfoBz),
		})
	}

	var name string
	if v, ok := dataMap[NameFieldKey]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				name = vStr
				delete(dataMap, MintRestrictedFieldKey)
			}
		}
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return nil, err
	}

	return codectypes.NewAnyWithValue(&NFTMetadata{
		Name: name,
		Data: string(data),
	})
}
