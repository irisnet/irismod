package types

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	proto "github.com/gogo/protobuf/proto"
)

const (
	Namespace          = "irismod:"
	KeyMediaFieldValue = "value"
)

var (
	ClassKeyName             = fmt.Sprintf("%s%s", Namespace, "name")
	ClassKeySymbol           = fmt.Sprintf("%s%s", Namespace, "symbol")
	ClassKeyDescription      = fmt.Sprintf("%s%s", Namespace, "description")
	ClassKeyURIhash          = fmt.Sprintf("%s%s", Namespace, "uri_hash")
	ClassKeyMintRestricted   = fmt.Sprintf("%s%s", Namespace, "mint_restricted")
	ClassKeyUpdateRestricted = fmt.Sprintf("%s%s", Namespace, "update_restricted")
	ClassKeyCreator          = fmt.Sprintf("%s%s", Namespace, "creator")
	ClassKeySchema           = fmt.Sprintf("%s%s", Namespace, "schema")
	TokenKeyName             = fmt.Sprintf("%s%s", Namespace, "name")
)

type (
	ClassResolver struct {
		cdc              codec.Codec
		getModuleAddress func(string) sdk.AccAddress
	}
	TokenResolver struct{ cdc codec.Codec }
	MediaField    struct {
		Value interface{} `json:"value"`
		Mime  string      `json:"mime,omitempty"`
	}
)

func NewClassResolver(cdc codec.Codec,
	getModuleAddress func(string) sdk.AccAddress,
) ClassResolver {
	return ClassResolver{
		cdc:              cdc,
		getModuleAddress: getModuleAddress,
	}
}

// Encode encode any into the metadata data format defined by ics721
func (cmr ClassResolver) Encode(class nft.Class) (string, error) {
	var message proto.Message
	if err := cmr.cdc.UnpackAny(class.Data, &message); err != nil {
		return "", err
	}

	metadata, ok := message.(*DenomMetadata)
	if !ok {
		return "", errors.New("unsupport classMetadata")
	}

	kvals := make(map[string]interface{})
	if len(metadata.Data) > 0 {
		if err := json.Unmarshal([]byte(metadata.Data), &kvals); err != nil {
			//when classData is not a legal json, there is no need to parse the data
			return base64.RawStdEncoding.EncodeToString([]byte(metadata.Data)), nil
		}
	}
	kvals[ClassKeyName] = MediaField{Value: class.Name}
	kvals[ClassKeySymbol] = MediaField{Value: class.Symbol}
	kvals[ClassKeyDescription] = MediaField{Value: class.Description}
	kvals[ClassKeyURIhash] = MediaField{Value: class.UriHash}
	kvals[ClassKeyMintRestricted] = MediaField{Value: metadata.MintRestricted}
	kvals[ClassKeyUpdateRestricted] = MediaField{Value: metadata.UpdateRestricted}
	kvals[ClassKeyCreator] = MediaField{Value: metadata.Creator}
	kvals[ClassKeySchema] = MediaField{Value: metadata.Schema}
	data, err := json.Marshal(kvals)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil
}

func (cmr ClassResolver) Decode(classID, classURI, classInfo string) (nft.Class, error) {
	classInfoBz, err := base64.RawStdEncoding.DecodeString(classInfo)
	if err != nil {
		return nft.Class{}, err
	}

	var (
		name             = ""
		symbol           = ""
		description      = ""
		uriHash          = ""
		mintRestricted   = true
		updateRestricted = true
		schema           = ""
		creator          = cmr.getModuleAddress(ModuleName).String()
	)

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(classInfoBz, &dataMap); err != nil {
		any, err := codectypes.NewAnyWithValue(&DenomMetadata{
			Creator:          creator,
			Schema:           schema,
			MintRestricted:   mintRestricted,
			UpdateRestricted: updateRestricted,
			Data:             string(classInfoBz),
		})
		if err != nil {
			return nft.Class{}, err
		}
		return nft.Class{
			Id:          classID,
			Uri:         classURI,
			Name:        name,
			Symbol:      symbol,
			Description: description,
			UriHash:     uriHash,
			Data:        any,
		}, nil
	}
	if v, ok := dataMap[ClassKeyName]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				name = vStr
				delete(dataMap, ClassKeyName)
			}
		}
	}

	if v, ok := dataMap[ClassKeySymbol]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				symbol = vStr
				delete(dataMap, ClassKeySymbol)
			}
		}
	}

	if v, ok := dataMap[ClassKeyDescription]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				description = vStr
				delete(dataMap, ClassKeyDescription)
			}
		}
	}

	if v, ok := dataMap[ClassKeyURIhash]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				uriHash = vStr
				delete(dataMap, ClassKeyURIhash)
			}
		}
	}

	if v, ok := dataMap[ClassKeyMintRestricted]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vBool, ok := vMap[KeyMediaFieldValue].(bool); ok {
				mintRestricted = vBool
				delete(dataMap, ClassKeyMintRestricted)
			}
		}
	}

	if v, ok := dataMap[ClassKeyUpdateRestricted]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vBool, ok := vMap[KeyMediaFieldValue].(bool); ok {
				updateRestricted = vBool
				delete(dataMap, ClassKeyUpdateRestricted)
			}
		}
	}

	if v, ok := dataMap[ClassKeyCreator]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				creator = vStr
				delete(dataMap, ClassKeyCreator)
			}
		}
	}

	if v, ok := dataMap[ClassKeySchema]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				schema = vStr
				delete(dataMap, ClassKeySchema)
			}
		}
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return nft.Class{}, err
	}

	any, err := codectypes.NewAnyWithValue(&DenomMetadata{
		Creator:          creator,
		Schema:           schema,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		Data:             string(data),
	})
	if err != nil {
		return nft.Class{}, err
	}

	return nft.Class{
		Id:          classID,
		Uri:         classURI,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		UriHash:     uriHash,
		Data:        any,
	}, nil
}

func NewTokenResolver(cdc codec.Codec) TokenResolver {
	return TokenResolver{
		cdc: cdc,
	}
}

func (cmr TokenResolver) Encode(any *codectypes.Any) (string, error) {
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

	kvals[TokenKeyName] = MediaField{Value: nftMetadata.Name}
	data, err := json.Marshal(kvals)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(data), nil
}

func (cmr TokenResolver) Decode(tokenInfo string) (*codectypes.Any, error) {
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
	if v, ok := dataMap[TokenKeyName]; ok {
		if vMap, ok := v.(map[string]interface{}); ok {
			if vStr, ok := vMap[KeyMediaFieldValue].(string); ok {
				name = vStr
				delete(dataMap, TokenKeyName)
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
