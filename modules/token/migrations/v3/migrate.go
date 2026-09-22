package v3

import (
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/ethereum/go-ethereum/common"

	"mods.irisnet.org/modules/token/types"
	v1 "mods.irisnet.org/modules/token/types/v1"
)

// Migrate removes the legacy token registration whose symbol overlaps a base denomination.
// Bank balances and supply are preserved.
func Migrate(ctx sdk.Context, key storetypes.StoreKey, cdc codec.BinaryCodec, bank types.BankKeeper) error {
	const symbol, minUnit = "uiris", "xuiris"
	configuredNative := v1.GetNativeToken()
	if configuredNative.Symbol != "iris" || configuredNative.MinUnit != symbol {
		return nil
	}

	store := ctx.KVStore(key)
	nativeRecord := store.Get(types.KeySymbol("iris"))
	nativeIndex := store.Get(types.KeyMinUint(symbol))
	if nativeRecord == nil || nativeIndex == nil {
		return fmt.Errorf("incomplete iris token namespace")
	}
	var native v1.Token
	if err := cdc.Unmarshal(nativeRecord, &native); err != nil {
		return fmt.Errorf("decode iris token: %w", err)
	}
	if native.Symbol != "iris" || native.MinUnit != symbol {
		return fmt.Errorf("unexpected iris token identity: %s/%s", native.Symbol, native.MinUnit)
	}
	var nativeSymbol gogotypes.StringValue
	if err := cdc.Unmarshal(nativeIndex, &nativeSymbol); err != nil {
		return fmt.Errorf("decode iris token min-unit index: %w", err)
	}
	if nativeSymbol.Value != native.Symbol {
		return fmt.Errorf("unexpected iris token min-unit index: %s", nativeSymbol.Value)
	}

	bz := store.Get(types.KeySymbol(symbol))
	if bz == nil {
		return nil
	}
	var token v1.Token
	if err := cdc.Unmarshal(bz, &token); err != nil {
		return fmt.Errorf("decode legacy token: %w", err)
	}
	if token.Symbol != symbol || token.MinUnit != minUnit {
		return fmt.Errorf("unexpected legacy token identity: %s/%s", token.Symbol, token.MinUnit)
	}

	indexKeys := [][]byte{types.KeyMinUint(minUnit)}
	if token.Owner != "" {
		owner, err := sdk.AccAddressFromBech32(token.Owner)
		if err != nil {
			return fmt.Errorf("invalid legacy token owner: %w", err)
		}
		indexKeys = append(indexKeys, types.KeyTokens(owner, symbol))
	}
	if token.Contract != "" {
		if !common.IsHexAddress(token.Contract) {
			return fmt.Errorf("invalid legacy token contract: %s", token.Contract)
		}
		indexKeys = append(indexKeys, types.KeyContract(token.Contract))
	}
	for i, indexKey := range indexKeys {
		bz := store.Get(indexKey)
		if bz == nil && i > 0 {
			continue
		}
		var index gogotypes.StringValue
		if bz == nil {
			return fmt.Errorf("missing legacy token min-unit index")
		}
		if err := cdc.Unmarshal(bz, &index); err != nil {
			return fmt.Errorf("decode legacy token index %x: %w", indexKey, err)
		}
		if index.Value != symbol {
			return fmt.Errorf("unexpected legacy token index %x: %s", indexKey, index.Value)
		}
	}

	// The v0.50 bank keeper exposes metadata through its collections map.
	var bk bankkeeper.BaseKeeper
	switch b := bank.(type) {
	case bankkeeper.BaseKeeper:
		bk = b
	case *bankkeeper.BaseKeeper:
		if b == nil {
			return fmt.Errorf("nil bank keeper for token migration")
		}
		bk = *b
	default:
		return fmt.Errorf("unsupported bank keeper for token migration: %T", bank)
	}
	metadata, err := bk.BaseViewKeeper.DenomMetadata.Get(ctx, minUnit)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return fmt.Errorf("read legacy token metadata: %w", err)
	}
	if err == nil && metadata.Base != minUnit {
		return fmt.Errorf("unexpected legacy token metadata base: %s", metadata.Base)
	}

	cacheCtx, write := ctx.CacheContext()
	if err := bk.BaseViewKeeper.DenomMetadata.Remove(cacheCtx, minUnit); err != nil {
		return fmt.Errorf("remove legacy token metadata: %w", err)
	}
	cacheStore := cacheCtx.KVStore(key)
	cacheStore.Delete(types.KeySymbol(symbol))
	for _, indexKey := range indexKeys {
		cacheStore.Delete(indexKey)
	}
	cacheStore.Delete(types.KeyBurnTokenAmt(minUnit))
	write()
	return nil
}
