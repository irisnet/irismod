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

func NewKeeper(cdc codec.Marshaler,
	storeKey sdk.StoreKey,
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
		paramSpace:       paramSpace,
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
func (k Keeper) GetPool(ctx sdk.Context, poolName string) (*types.FarmPool, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetFarmPoolKey(poolName))
	if len(bz) == 0 {
		return nil, false
	}

	var pool types.FarmPool
	k.cdc.MustUnmarshalBinaryBare(bz, &pool)
	return &pool, true
}

func (k Keeper) SetRewardRule(ctx sdk.Context, poolName string, rule types.RewardRule) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&rule)
	store.Set(types.GetFarmPoolRuleKey(poolName, rule.Reward), bz)
}

func (k Keeper) GetRewardRules(ctx sdk.Context, poolName string) (rules types.RewardRules) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetFarmPoolRulePrefix(poolName))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.RewardRule
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &r)
		rules = append(rules, &r)
	}
	return
}

func (k Keeper) IteratorRewardRules(ctx sdk.Context, poolName string, fun func(r types.RewardRule)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetFarmPoolRulePrefix(poolName))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.RewardRule
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &r)
		fun(r)
	}
}
