package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/farm/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Pools(goctx context.Context,
	request *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	var pools []*types.FarmPoolEntry
	k.IteratorAllPools(ctx, func(pool *types.FarmPool) {
		var totalReward sdk.Coins
		var remainingReward sdk.Coins
		var rewardPerBlock sdk.Coins
		k.IteratorRewardRules(ctx, pool.Name, func(r types.RewardRule) {
			totalReward = totalReward.Add(sdk.NewCoin(r.Reward, r.TotalReward))
			remainingReward = remainingReward.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
			rewardPerBlock = rewardPerBlock.Add(sdk.NewCoin(r.Reward, r.RewardPerBlock))
		})

		pools = append(pools, &types.FarmPoolEntry{
			Name:               pool.Name,
			Creator:            pool.Creator,
			BeginHeight:        pool.BeginHeight,
			EndHeight:          pool.EndHeight,
			Destructible:       pool.Destructible,
			Expired:            uint64(ctx.BlockHeight()) >= pool.EndHeight,
			TotalLpTokenLocked: pool.TotalLpTokenLocked,
			TotalReward:        totalReward,
			RemainingReward:    remainingReward,
			RewardPerBlock:     rewardPerBlock,
		})
	})

	return &types.QueryPoolsResponse{List: pools}, nil
}

func (k Keeper) Farmer(goctx context.Context,
	request *types.QueryFarmerRequest) (*types.QueryFarmerResponse, error) {
	var list []*types.LockedInfo
	var err error
	var farmInfos []types.FarmInfo

	ctx := sdk.UnwrapSDKContext(goctx)
	cacheCtx, _ := ctx.CacheContext()
	if len(request.PoolName) == 0 {
		k.IteratorFarmInfo(cacheCtx, request.Farmer, func(farmInfo types.FarmInfo) {
			farmInfos = append(farmInfos, farmInfo)
		})
	} else {
		farmInfo, existed := k.GetFarmInfo(cacheCtx, request.PoolName, request.Farmer)
		if existed {
			farmInfos = append(farmInfos, farmInfo)
		}
	}
	if len(farmInfos) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrNotExistFarmer, "not found farmer: %s", request.Farmer)
	}

	for _, farmer := range farmInfos {
		pool, _ := k.GetPool(cacheCtx, farmer.PoolName)
		pool, err = k.UpdatePool(cacheCtx, pool, sdk.ZeroInt(), false)
		if err != nil {
			return nil, err
		}
		rewards, _ := pool.CaclRewards(farmer, sdk.ZeroInt())
		list = append(list, &types.LockedInfo{
			PoolName:      farmer.PoolName,
			Locked:        sdk.NewCoin(pool.TotalLpTokenLocked.Denom, farmer.Locked),
			PendingReward: rewards,
		})
	}

	return &types.QueryFarmerResponse{
		List:   list,
		Height: ctx.BlockHeight(),
	}, nil
}
