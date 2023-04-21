package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// GetTokenPairs - get all registered token tokenPairs
func (k Keeper) GetTokenPairs(ctx sdk.Context) []types.TokenPair {
	tokenPairs := []types.TokenPair{}

	k.IterateTokenPairs(ctx, func(tokenPair types.TokenPair) (stop bool) {
		tokenPairs = append(tokenPairs, tokenPair)
		return false
	})

	return tokenPairs
}

// IterateTokenPairs iterates over all the stored token pairs
func (k Keeper) IterateTokenPairs(ctx sdk.Context, cb func(tokenPair types.TokenPair) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixTokenPair)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var tokenPair types.TokenPair
		k.cdc.MustUnmarshal(iterator.Value(), &tokenPair)

		if cb(tokenPair) {
			break
		}
	}
}

// GetTokenPairID returns the pair id from either of the registered tokens.
// Hex address or Denom can be used as token argument.
func (k Keeper) GetTokenPairID(ctx sdk.Context, token string) []byte {
	if common.IsHexAddress(token) {
		addr := common.HexToAddress(token)
		return k.GetERC721Map(ctx, addr)
	}
	return k.GetDenomMap(ctx, token)
}

// GetTokenPair gets a registered token pair from the identifier.
func (k Keeper) GetTokenPair(ctx sdk.Context, id []byte) (types.TokenPair, bool) {
	if id == nil {
		return types.TokenPair{}, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	var tokenPair types.TokenPair
	bz := store.Get(id)
	if len(bz) == 0 {
		return types.TokenPair{}, false
	}

	k.cdc.MustUnmarshal(bz, &tokenPair)
	return tokenPair, true
}

// SetTokenPair stores a token pair
func (k Keeper) SetTokenPair(ctx sdk.Context, tokenPair types.TokenPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	key := tokenPair.GetID()
	bz := k.cdc.MustMarshal(&tokenPair)
	store.Set(key, bz)
}

// DeleteTokenPair removes a token pair.
func (k Keeper) DeleteTokenPair(ctx sdk.Context, tokenPair types.TokenPair) {
	id := tokenPair.GetID()
	k.deleteTokenPair(ctx, id)
	k.deleteERC721Map(ctx, tokenPair.GetERC721Contract())
	k.deleteDenomMap(ctx, tokenPair.Denom)
}

// deleteTokenPair deletes the token pair for the given id
func (k Keeper) deleteTokenPair(ctx sdk.Context, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	store.Delete(id)
}

// GetERC721Map returns the token pair id for the given address
func (k Keeper) GetERC721Map(ctx sdk.Context, erc721 common.Address) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByERC721)
	return store.Get(erc721.Bytes())
}

// GetDenomMap returns the token pair id for the given denomination
func (k Keeper) GetDenomMap(ctx sdk.Context, denom string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByDenom)
	return store.Get([]byte(denom))
}

// SetERC721Map sets the token pair id for the given address
func (k Keeper) SetERC721Map(ctx sdk.Context, erc721 common.Address, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByERC721)
	store.Set(erc721.Bytes(), id)
}

// deleteERC721Map deletes the token pair id for the given address
func (k Keeper) deleteERC721Map(ctx sdk.Context, erc721 common.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByERC721)
	store.Delete(erc721.Bytes())
}

// SetDenomMap sets the token pair id for the denomination
func (k Keeper) SetDenomMap(ctx sdk.Context, denom string, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByDenom)
	store.Set([]byte(denom), id)
}

// deleteDenomMap deletes the token pair id for the given denom
func (k Keeper) deleteDenomMap(ctx sdk.Context, denom string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByDenom)
	store.Delete([]byte(denom))
}

// IsTokenPairRegistered - check if registered token tokenPair is registered
func (k Keeper) IsTokenPairRegistered(ctx sdk.Context, id []byte) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	return store.Has(id)
}

// IsERC721Registered check if registered ERC20 token is registered
func (k Keeper) IsERC721Registered(ctx sdk.Context, erc721 common.Address) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByERC721)
	return store.Has(erc721.Bytes())
}

// IsDenomRegistered check if registered coin denom is registered
func (k Keeper) IsDenomRegistered(ctx sdk.Context, denom string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByDenom)
	return store.Has([]byte(denom))
}
