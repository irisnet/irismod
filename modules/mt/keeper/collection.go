package keeper

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/irisnet/irismod/modules/mt/exported"
	"github.com/irisnet/irismod/modules/mt/types"
)

// SetCollection saves all MTs and returns an error if there already exists
func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) error {
	for _, mt := range collection.MTs {
		if err := k.MintMT(
			ctx,
			collection.Denom.Id,
			mt.GetID(),
			mt.GetName(),
			mt.GetURI(),
			mt.GetURIHash(),
			mt.GetData(),
			mt.GetOwner(),
		); err != nil {
			return err
		}
	}
	return nil
}

// GetCollection returns the collection by the specified denom ID
func (k Keeper) GetCollection(ctx sdk.Context, denomID string) (types.Collection, error) {
	denom, found := k.GetDenom(ctx, denomID)
	if !found {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not existed ", denomID)
	}

	mts := k.GetMTs(ctx, denomID)
	return types.NewCollection(denom, mts), nil
}

// GetPaginateCollection returns the collection by the specified denom ID
func (k Keeper) GetPaginateCollection(ctx sdk.Context, request *types.QueryCollectionRequest, denomID string) (types.Collection, *query.PageResponse, error) {
	denom, found := k.GetDenom(ctx, denomID)
	if !found {
		return types.Collection{}, nil, sdkerrors.Wrapf(types.ErrInvalidDenom, "denomID %s not existed ", denomID)
	}
	var mts []exported.MT
	store := ctx.KVStore(k.storeKey)
	mtStore := prefix.NewStore(store, types.KeyMT(denomID, ""))
	pageRes, err := query.Paginate(mtStore, request.Pagination, func(key []byte, value []byte) error {
		var baseMT types.BaseMT
		k.cdc.MustUnmarshal(value, &baseMT)
		mts = append(mts, baseMT)
		return nil
	})
	if err != nil {
		return types.Collection{}, nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}
	return types.NewCollection(denom, mts), pageRes, nil
}

// GetCollections returns all the collections
func (k Keeper) GetCollections(ctx sdk.Context) (cs []types.Collection) {
	for _, denom := range k.GetDenoms(ctx) {
		mts := k.GetMTs(ctx, denom.Id)
		cs = append(cs, types.NewCollection(denom, mts))
	}
	return cs
}

// GetTotalSupply returns the number of MTs by the specified denom ID
func (k Keeper) GetTotalSupply(ctx sdk.Context, denomID string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCollection(denomID))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(k.cdc, bz)
}

// GetTotalSupplyOfOwner returns the amount of MTs by the specified conditions
func (k Keeper) GetTotalSupplyOfOwner(ctx sdk.Context, id string, owner sdk.AccAddress) (supply uint64) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(owner, id, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		supply++
	}
	return supply
}

func (k Keeper) increaseSupply(ctx sdk.Context, denomID string) {
	supply := k.GetTotalSupply(ctx, denomID)
	supply++

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollection(denomID), bz)
}

func (k Keeper) decreaseSupply(ctx sdk.Context, denomID string) {
	supply := k.GetTotalSupply(ctx, denomID)
	supply--

	store := ctx.KVStore(k.storeKey)
	if supply == 0 {
		store.Delete(types.KeyCollection(denomID))
		return
	}

	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollection(denomID), bz)
}
