package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/irisnet/irismod/modules/farm/types"
)

// Keeper of the farm store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler

	paramSpace paramstypes.Subspace
	// name of the fee collector
	feeCollectorName string
	bk               types.BankKeeper
	ck               types.CoinswapKeeper
}

func NewKeeper(storeKey sdk.StoreKey, cdc codec.Marshaler,
	bk types.BankKeeper,
	ck types.CoinswapKeeper,
	paramSpace paramstypes.Subspace,
	feeCollectorName string,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ParamKeyTable())
	}
	return Keeper{
		storeKey:         storeKey,
		cdc:              cdc,
		bk:               bk,
		ck:               ck,
		feeCollectorName: feeCollectorName,
	}
}

// CreatePool creates an new farm pool
func (k Keeper) SetPool(ctx sdk.Context, pool types.FarmPool) {
	pool.Rules = nil
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&pool)
	store.Set(types.GetFarmPoolKey(pool.Name), bz)
}

// GetPool return the specified farm pool
func (k Keeper) GetPool(ctx sdk.Context, poolName string) (fp *types.FarmPool, exist bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFarmPoolKey(poolName))
	if len(bz) == 0 {
		return nil, false
	}

	k.cdc.MustUnmarshalBinaryBare(bz, fp)
	return fp, true
}

func (k Keeper) SetPoolRule(ctx sdk.Context, poolName string, rule types.FarmRule) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&rule)
	store.Set(types.GetFarmPoolRuleKey(poolName, rule.Reward), bz)
}

func (k Keeper) GetPoolRules(ctx sdk.Context, poolName string) (rules []*types.FarmRule) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetFarmPoolRulePrefix(poolName))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.FarmRule
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &r)
		rules = append(rules, &r)
	}
	return
}
