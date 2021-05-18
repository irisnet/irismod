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
	k.IteratorActivePool(ctx, func(pool *types.FarmPool) {
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
			Owner:              pool.Creator,
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
	var farmerInfos []types.Farmer

	ctx := sdk.UnwrapSDKContext(goctx)
	cacheCtx, _ := ctx.CacheContext()
	if len(request.PoolName) == 0 {
		k.IteratorFarmer(cacheCtx, request.Farmer, func(farmer types.Farmer) error {
			farmerInfos = append(farmerInfos, farmer)
			pool, _ := k.GetPool(cacheCtx, farmer.PoolName)
			pool, err = k.UpdatePool(cacheCtx, pool, sdk.ZeroInt(), false)
			if err != nil {
				return err
			}
			rewards, _ := pool.CaclRewards(farmer, sdk.ZeroInt())
			list = append(list, &types.LockedInfo{
				PoolName:      farmer.PoolName,
				Locked:        sdk.NewCoin(pool.TotalLpTokenLocked.Denom, farmer.Locked),
				PendingReward: rewards,
			})
			return nil
		})
	} else {
		farmer, existed := k.GetFarmer(cacheCtx, request.PoolName, request.Farmer)
		if existed {
			farmerInfos = append(farmerInfos, *farmer)
		}
	}
	if len(farmerInfos) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrNotExistFarmer, "not found farmer: %s", request.Farmer)
	}

	for _, farmer := range farmerInfos {
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
