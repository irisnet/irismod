package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/mt/types"
)

// GetOwner gets all the ID collections owned by an address and denom ID
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress, denomId string) types.Owner {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, denomId, ""))
	defer iterator.Close()

	owner := types.Owner{
		Address:       address.String(),
		IdCollections: types.IDCollections{},
	}
	idsMap := make(map[string][]types.Balance)

	for ; iterator.Valid(); iterator.Next() {
		_, denomID, mtID, _ := types.SplitKeyOwner(iterator.Key())

		balance := k.getBalance(ctx, address, denomID, mtID)
		if ids, ok := idsMap[denomID]; ok {
			idsMap[denomID] = append(ids, balance)
		} else {
			idsMap[denomID] = append([]types.Balance{}, balance)
			owner.IdCollections = append(
				owner.IdCollections,
				types.IDCollection{DenomId: denomID},
			)
		}
	}

	for i := 0; i < len(owner.IdCollections); i++ {
		owner.IdCollections[i].Balances = idsMap[owner.IdCollections[i].DenomId]
	}

	return owner
}

// GetOwners gets all the ID collections
func (k Keeper) GetOwners(ctx sdk.Context) (owners types.Owners) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.KeyOwner(nil, "", ""))
	defer iterator.Close()

	idcsMap := make(map[string]types.IDCollections)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		address, denomId, mtId, _ := types.SplitKeyOwner(key)
		if _, ok := idcsMap[address.String()]; !ok {
			idcsMap[address.String()] = types.IDCollections{}
			owners = append(
				owners,
				types.Owner{Address: address.String()},
			)
		}
		idcs := idcsMap[address.String()]
		balance := k.getBalance(ctx, address, denomId, mtId)
		idcs = idcs.Add(denomId, balance)
		idcsMap[address.String()] = idcs
	}
	for i, owner := range owners {
		owners[i].IdCollections = idcsMap[owner.Address]
	}

	return owners
}

func (k Keeper) deleteOwner(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, denomID, tokenID))
}

func (k Keeper) setOwner(ctx sdk.Context, denomID, tokenID string, amount uint64, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	SrcBalanceAmount := amount - k.getBalance(ctx, owner, denomID, tokenID).Amount

	amountByte := make([]byte, 8)
	binary.BigEndian.PutUint64(amountByte, SrcBalanceAmount)

	store.Set(types.KeyOwner(owner, denomID, tokenID), amountByte)
}

func (k Keeper) setBalance(ctx sdk.Context,
	denomID, tokenID string,
	amount uint64,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	SrcBalanceAmount := amount - k.getBalance(ctx, owner, denomID, tokenID).Amount

	amountByte := make([]byte, 8)
	binary.BigEndian.PutUint64(amountByte, SrcBalanceAmount)

	store.Set(types.KeyOwner(owner, denomID, tokenID), amountByte)
}

func (k Keeper) getBalance(ctx sdk.Context,
	owner sdk.AccAddress,
	denomID, tokenID string,
) (balance types.Balance) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyOwner(owner, denomID, tokenID))
	if len(bz) == 0 {
		return balance
	}

	k.cdc.MustUnmarshal(bz, &balance)
	return balance
}

func (k Keeper) swapOwner(ctx sdk.Context, denomID, tokenID string, amount uint64, srcOwner, dstOwner sdk.AccAddress) {

	// delete old owner key
	k.setOwner(ctx, denomID, tokenID, amount, srcOwner)

	// set new owner key
	k.setBalance(ctx, denomID, tokenID, amount, dstOwner)
}
