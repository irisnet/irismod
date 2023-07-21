package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/irisnet/irismod/modules/farm/types"
)

// Keeper of the farm store
type Keeper struct {
	cdc                                 codec.Codec
	storeKey                            storetypes.StoreKey
	paramSpace                          paramstypes.Subspace
	validateLPToken                     types.ValidateLPToken
	bk                                  types.BankKeeper
	ak                                  types.AccountKeeper
	dk                                  types.DistrKeeper
	gk                                  types.GovKeeper
	feeCollectorName, communityPoolName string // name of the fee collector
}

func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	bk types.BankKeeper,
	ak types.AccountKeeper,
	dk types.DistrKeeper,
	gk types.GovKeeper,
	validateLPToken types.ValidateLPToken,
	paramSpace paramstypes.Subspace,
	feeCollectorName, communityPoolName string,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ParamKeyTable())
	}

	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// ensure farm module accounts are set
	if addr := ak.GetModuleAddress(types.RewardCollector); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.RewardCollector))
	}

	return Keeper{
		storeKey:          storeKey,
		cdc:               cdc,
		bk:                bk,
		ak:                ak,
		dk:                dk,
		gk:                gk,
		validateLPToken:   validateLPToken,
		paramSpace:        paramSpace,
		feeCollectorName:  feeCollectorName,
		communityPoolName: communityPoolName,
	}
}

// CreatePool creates an new farm pool
func (k Keeper) SetPool(ctx sdk.Context, pool types.FarmPool) {
	pool.Rules = nil
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pool)
	store.Set(types.KeyFarmPool(pool.Id), bz)
}

// GetPool return the specified farm pool
func (k Keeper) GetPool(ctx sdk.Context, poolId string) (types.FarmPool, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyFarmPool(poolId))
	if len(bz) == 0 {
		return types.FarmPool{}, false
	}

	var pool types.FarmPool
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}

func (k Keeper) SetRewardRules(ctx sdk.Context, poolId string, rules types.RewardRules) {
	for _, r := range rules {
		k.SetRewardRule(ctx, poolId, r)
	}
}

func (k Keeper) SetRewardRule(ctx sdk.Context, poolId string, rule types.RewardRule) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rule)
	store.Set(types.KeyRewardRule(poolId, rule.Reward), bz)
}

func (k Keeper) GetRewardRules(ctx sdk.Context, poolId string) (rules types.RewardRules) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixRewardRule(poolId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.RewardRule
		k.cdc.MustUnmarshal(iterator.Value(), &r)
		rules = append(rules, r)
	}
	return
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "irismod/farm")
}

func (k Keeper) IteratorRewardRules(ctx sdk.Context, poolId string, fun func(r types.RewardRule)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixRewardRule(poolId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var r types.RewardRule
		k.cdc.MustUnmarshal(iterator.Value(), &r)
		fun(r)
	}
}

func (k Keeper) GetSequence(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyFarmPoolSeq())
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetSequence(ctx sdk.Context, seq uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyFarmPoolSeq(), sdk.Uint64ToBigEndian(seq))
}
