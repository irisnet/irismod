package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/x/nft"
	gogotypes "github.com/gogo/protobuf/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/nft/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Supply(c context.Context, request *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var supply uint64
	switch {
	case len(request.Owner) == 0 && len(request.DenomId) > 0:
		supply = k.GetTotalSupply(ctx, request.DenomId)
	default:
		owner, err := sdk.AccAddressFromBech32(request.Owner)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid owner address %s", request.Owner)
		}
		supply = k.GetTotalSupplyOfOwner(ctx, request.DenomId, owner)
	}
	return &types.QuerySupplyResponse{Amount: supply}, nil
}

func (k Keeper) Owner(c context.Context, request *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error) {
	r := &nft.QueryNFTsOfClassRequest{
		ClassId:    request.DenomId,
		Owner:      request.Owner,
		Pagination: request.Pagination,
	}

	// TODO DenomId should be optional
	result, err := k.nk.NFTsOfClass(c, r)
	if err != nil {
		return nil, err
	}

	var denomMap = make(map[string][]string)
	for _, token := range result.Nfts {
		denomMap[token.ClassId] = append(denomMap[token.ClassId], token.Id)
	}

	var idc []types.IDCollection
	for denom, ids := range denomMap {
		idc = append(idc, types.IDCollection{DenomId: denom, TokenIds: ids})
	}

	response := &types.QueryOwnerResponse{
		Owner: &types.Owner{
			Address:       request.Owner,
			IDCollections: idc,
		},
		Pagination: result.Pagination,
	}

	return response, nil
}

func (k Keeper) Collection(c context.Context, request *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	denom, err := k.GetDenomInfo(ctx, request.DenomId)
	if err != nil {
		return nil, err
	}

	r := &nft.QueryNFTsOfClassRequest{
		ClassId:    request.DenomId,
		Pagination: request.Pagination,
	}

	result, err := k.nk.NFTsOfClass(c, r)
	if err != nil {
		return nil, err
	}

	var nfts []types.BaseNFT
	for _, token := range result.Nfts {
		owner := k.nk.GetOwner(ctx, request.DenomId, token.Id)

		var data = &gogotypes.StringValue{}
		if err := k.cdc.Unmarshal(token.GetData().Value, data); err != nil {
			return nil, err
		}
		nfts = append(nfts, types.BaseNFT{
			Id:    token.Id,
			URI:   token.Uri,
			Owner: owner.String(),
			Data:  data.GetValue(),
		})
	}

	collection := &types.Collection{
		Denom: *denom,
		NFTs:  nfts,
	}

	response := &types.QueryCollectionResponse{
		Collection: collection,
		Pagination: result.Pagination,
	}

	return response, nil
}

func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	denom, err := k.GetDenomInfo(ctx, request.DenomId)
	if err != nil {
		return nil, err
	}
	return &types.QueryDenomResponse{Denom: denom}, nil
}

func (k Keeper) Denoms(c context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	result, err := k.nk.Classes(c, &nft.QueryClassesRequest{
		Pagination: req.Pagination,
	})
	if err != nil {
		return nil, err
	}

	var denoms []types.Denom
	for _, denom := range result.Classes {
		denom, err := k.GetDenomInfo(ctx, denom.Id)
		if err != nil {
			return nil, err
		}
		denoms = append(denoms, *denom)
	}

	return &types.QueryDenomsResponse{
		Denoms:     denoms,
		Pagination: result.Pagination,
	}, nil
}

func (k Keeper) NFT(c context.Context, request *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	nft, err := k.GetNFT(ctx, request.DenomId, request.TokenId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", request.TokenId, request.DenomId)
	}

	baseNFT, ok := nft.(types.BaseNFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type NFT %s from collection %s", request.TokenId, request.DenomId)
	}

	return &types.QueryNFTResponse{NFT: &baseNFT}, nil
}
