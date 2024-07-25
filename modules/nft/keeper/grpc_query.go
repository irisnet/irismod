package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mods.irisnet.org/modules/nft/types"
)

var _ types.QueryServer = Keeper{}

// Supply queries the total supply of a given denom or owner.
//
// The function takes a context and a QuerySupplyRequest as parameters.
// It returns a QuerySupplyResponse and an error.
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

// Owner queries the NFTs of the specified owner.
//
// The function takes a context and a QueryOwnerRequest as parameters.
// It returns a QueryOwnerResponse and an error.
//
// The function first unwraps the context to get the SDK context.
// Then it converts the owner address from a Bech32 string to an SDK AccAddress.
// If the conversion fails, it returns an error.
//
// The function initializes an empty Owner struct with the owner address.
// It also initializes an empty map to store the token IDs.
//
// The function retrieves the store key and creates a new prefix store.
// It then paginates the NFT store using the provided pagination request.
// For each key-value pair in the store, it splits the key into the denom ID and token ID.
// If the denom ID is not provided in the request, it is extracted from the key.
// The function checks if the token ID is already present in the map.
// If it is, it appends the token ID to the existing list.
// Otherwise, it creates a new entry in the map and adds the token ID.
//
// After pagination, the function iterates over the IDCollections and assigns the corresponding token IDs.
//
// Finally, it returns a QueryOwnerResponse containing the Owner and the pagination result.
func (k Keeper) Owner(c context.Context, request *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	ownerAddress, err := sdk.AccAddressFromBech32(request.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid owner address %s", request.Owner)
	}

	owner := types.Owner{
		Address:       ownerAddress.String(),
		IDCollections: types.IDCollections{},
	}
	idsMap := make(map[string][]string)
	store := ctx.KVStore(k.storeKey)
	nftStore := prefix.NewStore(store, types.KeyOwner(ownerAddress, request.DenomId, ""))
	pageRes, err := query.Paginate(nftStore, shapePageRequest(request.Pagination), func(key []byte, value []byte) error {
		denomID := request.DenomId
		tokenID := string(key)
		if len(request.DenomId) == 0 {
			denomID, tokenID, _ = types.SplitKeyDenom(key)
		}
		if ids, ok := idsMap[denomID]; ok {
			idsMap[denomID] = append(ids, tokenID)
		} else {
			idsMap[denomID] = []string{tokenID}
			owner.IDCollections = append(
				owner.IDCollections,
				types.IDCollection{DenomId: denomID},
			)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(owner.IDCollections); i++ {
		owner.IDCollections[i].TokenIds = idsMap[owner.IDCollections[i].DenomId]
	}
	return &types.QueryOwnerResponse{Owner: &owner, Pagination: pageRes}, nil
}

// Collection retrieves a collection based on the given QueryCollectionRequest.
//
// Parameters:
// - c: The context.Context object.
// - request: The QueryCollectionRequest object.
//
// Returns:
// - *types.QueryCollectionResponse: The QueryCollectionResponse object.
// - error: An error if any occurred.
func (k Keeper) Collection(c context.Context, request *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	collection, pageRes, err := k.GetPaginateCollection(ctx, request, request.DenomId)
	if err != nil {
		return nil, err
	}
	return &types.QueryCollectionResponse{Collection: &collection, Pagination: pageRes}, nil
}

// Denom retrieves a denom based on the given QueryDenomRequest.
//
// Parameters:
// - c: The context.Context object.
// - request: The QueryDenomRequest object.
//
// Returns:
// - *types.QueryDenomResponse: The QueryDenomResponse object.
// - error: An error if the denom ID does not exist.
func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	denomObject, found := k.GetDenom(ctx, request.DenomId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", request.DenomId)
	}

	return &types.QueryDenomResponse{Denom: &denomObject}, nil
}

// Denoms retrieves a list of denoms based on the given QueryDenomsRequest.
//
// Parameters:
// - c: The context.Context object.
// - req: The QueryDenomsRequest object.
//
// Returns:
// - *types.QueryDenomsResponse: The QueryDenomsResponse object.
// - error: An error if the pagination is invalid.
func (k Keeper) Denoms(c context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var denoms []types.Denom
	store := ctx.KVStore(k.storeKey)
	denomStore := prefix.NewStore(store, types.KeyDenomID(""))
	pageRes, err := query.Paginate(denomStore, shapePageRequest(req.Pagination), func(key []byte, value []byte) error {
		var denom types.Denom
		k.cdc.MustUnmarshal(value, &denom)
		denoms = append(denoms, denom)
		return nil
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryDenomsResponse{
		Denoms:     denoms,
		Pagination: pageRes,
	}, nil
}

// NFT retrieves an NFT based on the given QueryNFTRequest.
//
// Parameters:
// - c: The context.Context object.
// - request: The QueryNFTRequest object.
//
// Returns:
// - *types.QueryNFTResponse: The QueryNFTResponse object.
// - error: An error if the NFT is not found or has an invalid type.
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
