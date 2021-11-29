package v160

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/nft/types"
)

func MigrateStore(ctx sdk.Context,
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	issueDenom IssueDenomFn,
	mintNFT MintNFTFn,
) error {
	store := ctx.KVStore(storeKey)
	denoms, err := migrateDenoms(ctx, store, cdc, issueDenom)
	if err != nil {
		return err
	}

	if err = migrateNFTs(ctx, store, cdc, denoms, mintNFT); err != nil {
		return err
	}
	return nil
}

func migrateDenoms(ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	issueDenom IssueDenomFn,
) (denoms []types.Denom, err error) {
	iterator := sdk.KVStorePrefixIterator(store, KeyDenomID(""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var denom types.Denom
		cdc.MustUnmarshal(iterator.Value(), &denom)
		denoms = append(denoms, denom)

		// delete denom from store
		store.Delete(iterator.Key())
		store.Delete(KeyDenomName(denom.Name))
		store.Delete(KeyCollection(denom.Id))

		creator, err := sdk.AccAddressFromBech32(denom.Creator)
		if err != nil {
			return nil, err
		}
		err = issueDenom(
			ctx,
			denom.Id,
			denom.Name,
			denom.Schema,
			denom.Symbol,
			creator,
			denom.MintRestricted,
			denom.UpdateRestricted,
		)
		if err != nil {
			return nil, err
		}
	}
	return denoms, nil
}

func migrateNFTs(ctx sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
	denoms []types.Denom,
	mintNFT MintNFTFn,
) error {
	for _, denom := range denoms {
		iterator := sdk.KVStorePrefixIterator(store, KeyNFT(denom.Id, ""))
		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			var baseNFT types.BaseNFT
			cdc.MustUnmarshal(iterator.Value(), &baseNFT)

			// delete nft from store
			store.Delete(iterator.Key())

			creator, err := sdk.AccAddressFromBech32(denom.Creator)
			if err != nil {
				return err
			}

			receiver, err := sdk.AccAddressFromBech32(baseNFT.Owner)
			if err != nil {
				return err
			}

			// delete owner from store
			store.Delete(KeyOwner(receiver, denom.Id, baseNFT.Id))

			if err = mintNFT(ctx,
				denom.Id,
				baseNFT.Id,
				baseNFT.Name,
				baseNFT.URI,
				baseNFT.Data,
				creator,
				receiver,
			); err != nil {
				return err
			}
		}
	}
	return nil
}
