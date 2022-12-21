package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/irisnet/irismod/modules/nft/types"
)

var _ types.QueryServer = Keeper{}

// Supply queries the total supply of a given denom or owner
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
		supply = k.GetBalance(ctx, request.DenomId, owner)
	}
	return &types.QuerySupplyResponse{Amount: supply}, nil
}

// NFTsOfOwner queries the NFTs of the specified owner
func (k Keeper) NFTsOfOwner(c context.Context, request *types.QueryNFTsOfOwnerRequest) (*types.QueryNFTsOfOwnerResponse, error) {
	r := &nft.QueryNFTsRequest{
		ClassId:    request.DenomId,
		Owner:      request.Owner,
		Pagination: shapePageRequest(request.Pagination),
	}

	result, err := k.nk.NFTs(c, r)
	if err != nil {
		return nil, err
	}

	var denomMap = make(map[string][]string)
	var denoms []string
	for _, token := range result.Nfts {
		if denomMap[token.ClassId] == nil {
			denomMap[token.ClassId] = []string{}
			denoms = append(denoms, token.ClassId)
		}
		denomMap[token.ClassId] = append(denomMap[token.ClassId], token.Id)
	}

	var idc []types.IDCollection
	for _, denomID := range denoms {
		idc = append(idc, types.IDCollection{
			DenomId:  denomID,
			TokenIds: denomMap[denomID],
		})
	}

	response := &types.QueryNFTsOfOwnerResponse{
		Owner: &types.Owner{
			Address:       request.Owner,
			IDCollections: idc,
		},
		Pagination: result.Pagination,
	}

	return response, nil
}

// Collection queries the NFTs of the specified denom
func (k Keeper) Collection(c context.Context, request *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	denom, err := k.GetDenomInfo(ctx, request.DenomId)
	if err != nil {
		return nil, err
	}

	r := &nft.QueryNFTsRequest{
		ClassId:    request.DenomId,
		Pagination: shapePageRequest(request.Pagination),
	}

	result, err := k.nk.NFTs(c, r)
	if err != nil {
		return nil, err
	}

	var nfts []types.BaseNFT
	for _, token := range result.Nfts {
		owner := k.nk.GetOwner(ctx, request.DenomId, token.Id)

		var nftMetadata types.NFTMetadata
		if err := k.cdc.Unmarshal(token.Data.GetValue(), &nftMetadata); err != nil {
			return nil, err
		}

		nfts = append(nfts, types.BaseNFT{
			Id:      token.Id,
			URI:     token.Uri,
			UriHash: token.UriHash,
			Name:    nftMetadata.Name,
			Owner:   owner.String(),
			Data:    nftMetadata.Data,
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

// Denom queries the definition of a given denom
func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	denom, err := k.GetDenomInfo(ctx, request.DenomId)
	if err != nil {
		return nil, err
	}
	return &types.QueryDenomResponse{Denom: denom}, nil
}

// Denoms queries all the denoms
func (k Keeper) Denoms(c context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	result, err := k.nk.Classes(c, &nft.QueryClassesRequest{
		Pagination: shapePageRequest(req.Pagination),
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

// NFT queries the NFT for the given denom and token ID
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

// Rental Plugin Query Service

// User queries the user of an nft
func (k Keeper) UserOf(goCtx context.Context, msg *types.QueryUserOfRequest) (*types.QueryUserOfResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.DenomId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.DenomId, msg.NftId)
	}

	rental, err := k.GetRentalInfo(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, err
	}

	return &types.QueryUserOfResponse{User: rental.User}, nil
}

// Expires queries the expires of an nft
func (k Keeper) UserExpires(goCtx context.Context, msg *types.QueryUserExpiresRequest) (*types.QueryUserExpiresResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.DenomId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.DenomId, msg.NftId)
	}

	rental, err := k.GetRentalInfo(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, err
	}

	return &types.QueryUserExpiresResponse{Expires: rental.Expires}, nil
}

// HasUser queries if an nft has the user
func (k Keeper) HasUser(goCtx context.Context, msg *types.QueryHasUserRequest) (*types.QueryHasUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if exist := k.nk.HasNFT(ctx, msg.DenomId, msg.NftId); !exist {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "%s-%s is not existent", msg.DenomId, msg.NftId)
	}

	rental, err := k.GetRentalInfo(ctx, msg.DenomId, msg.NftId)
	if err != nil {
		return nil, err
	}

	// if expires or user not existent, return false
	var resp types.QueryHasUserResponse
	if ctx.BlockTime().Unix() > rental.Expires || rental.User == "" {
		resp.HasUser = false
	}

	return &resp, nil
}
