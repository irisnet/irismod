package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/nft/types"
)

func Migrate(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.Codec, k NFTKeeper) error {
	denoms, err := migrateDenoms(ctx, storeKey, cdc, k)
	if err != nil {
		return err
	}
	return migrateTokens(ctx, storeKey, cdc, denoms, k)
}

func migrateDenoms(ctx sdk.Context,
	storeKey storetypes.StoreKey,
	cdc codec.Codec,
	k NFTKeeper,
) (denoms []string, err error) {
	store := ctx.KVStore(storeKey)
	iterator := sdk.KVStorePrefixIterator(store, KeyDenom(""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var denom types.Denom
		cdc.MustUnmarshal(iterator.Value(), &denom)

		//delete unused key
		store.Delete(KeyDenom(denom.Id))
		store.Delete(KeyDenomName(denom.Name))
		store.Delete(KeyCollection(denom.Id))

		creator, err := sdk.AccAddressFromBech32(denom.Creator)
		if err != nil {
			return denoms, err
		}

		if err := k.SaveDenom(ctx, denom.Id,
			denom.Name,
			denom.Schema,
			denom.Symbol,
			creator,
			denom.MintRestricted,
			denom.UpdateRestricted,
			denom.Description,
			denom.Uri,
			denom.UriHash,
			denom.Data,
		); err != nil {
			return denoms, err
		}
		denoms = append(denoms, denom.Id)

	}
	return denoms, nil
}

func migrateTokens(ctx sdk.Context,
	storeKey storetypes.StoreKey,
	cdc codec.Codec,
	denoms []string,
	k NFTKeeper,
) error {
	store := ctx.KVStore(storeKey)

	var iterator sdk.Iterator
	defer func() {
		if iterator != nil {
			_ = iterator.Close()
		}
	}()

	for _, denomID := range denoms {
		iterator = sdk.KVStorePrefixIterator(store, KeyNFT(denomID, ""))
		for ; iterator.Valid(); iterator.Next() {
			var baseNFT types.BaseNFT
			cdc.MustUnmarshal(iterator.Value(), &baseNFT)

			owner, err := sdk.AccAddressFromBech32(baseNFT.Owner)
			if err != nil {
				return err
			}

			//delete unused key
			store.Delete(KeyNFT(denomID, baseNFT.Id))
			store.Delete(KeyOwner(owner, denomID, baseNFT.Id))

			if err := k.SaveNFT(ctx, denomID,
				baseNFT.Id,
				baseNFT.Name,
				baseNFT.URI,
				baseNFT.UriHash,
				baseNFT.Data,
				owner,
			); err != nil {
				return err
			}
		}
	}
	return nil
}
