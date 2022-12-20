package types

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

const (
	Namespace = "irismod:"
)

var (
	MintRestrictedFieldKey   = fmt.Sprintf("%s%s", Namespace, "mint_restricted")
	UpdateRestrictedFieldKey = fmt.Sprintf("%s%s", Namespace, "update_restricted")
	CreatorFieldKey          = fmt.Sprintf("%s%s", Namespace, "creator")
	SchemaFieldKey           = fmt.Sprintf("%s%s", Namespace, "schema")
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

func (cmr ClassMetadataResolver) Encode(any *codectypes.Any) (string, error) {
	var message proto.Message
	if err := cmr.cdc.UnpackAny(any, &message); err != nil {
		return "", err
	}

	denomMetadata, ok := message.(*DenomMetadata)
	if !ok {
		return "", errors.New("unsupport classMetadata")
	}
	if !gjson.Valid(denomMetadata.Data) {
		//when classData is not a legal json, there is no need to parse the data
		return base64.RawStdEncoding.EncodeToString([]byte(denomMetadata.Data)), nil
	}

	kvals := make(map[string]interface{})
	if err := json.Unmarshal([]byte(denomMetadata.Data), &kvals); err != nil {
		return "", err
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

	if !gjson.ValidBytes(classInfoBz) {
		return codectypes.NewAnyWithValue(&DenomMetadata{
			Creator:          creator,
			MintRestricted:   mintRestricted,
			UpdateRestricted: updateRestricted,
			Data:             string(classInfoBz),
		})
	}

	jsonMap := make(map[string]MediaField)
	if err := json.Unmarshal(classInfoBz, &jsonMap); err != nil {
		return nil, err
	}

	if v, ok := jsonMap[MintRestrictedFieldKey]; ok {
		if vBool, ok := v.Value.(bool); ok {
			mintRestricted = vBool
			delete(jsonMap, MintRestrictedFieldKey)
		}
	}

	if v, ok := jsonMap[UpdateRestrictedFieldKey]; ok {
		if vBool, ok := v.Value.(bool); ok {
			updateRestricted = vBool
			delete(jsonMap, UpdateRestrictedFieldKey)
		}
	}

	if v, ok := jsonMap[CreatorFieldKey]; ok {
		if vStr, ok := v.Value.(string); ok {
			creator = vStr
			delete(jsonMap, CreatorFieldKey)
		}
	}

	if v, ok := jsonMap[SchemaFieldKey]; ok {
		if vStr, ok := v.Value.(string); ok {
			schema = vStr
			delete(jsonMap, SchemaFieldKey)
		}
	}

	data, err := json.Marshal(jsonMap)
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

	return "", nil
}

func (cmr TokenMetadataResolver) Decode(classInfo string) (*codectypes.Any, error) {
	return nil, nil
}
