package types

import (
	math "math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (pool FarmPool) ExpiredHeight() uint64 {
	var expiredHeight = uint64(math.MaxUint64)
	for _, r := range pool.Rules {
		inteval := r.TotalReward.Quo(r.RewardPerBlock).Uint64()
		if inteval+pool.BeginHeight < expiredHeight {
			expiredHeight = inteval + pool.BeginHeight
		}
	}
	return expiredHeight + 1
}

func (pool FarmPool) IsExpired(ctx sdk.Context) bool {
	return pool.EndHeight <= uint64(ctx.BlockHeight())
}

func (pool FarmPool) CaclRewards(farmInfo FarmInfo, deltaAmt sdk.Int) (rewards, rewardDebt sdk.Coins) {
	for _, r := range pool.Rules {
		if farmInfo.Locked.GT(sdk.ZeroInt()) {
			pendingRewardTotal := r.RewardPerShare.MulInt(farmInfo.Locked).TruncateInt()
			pendingReward := pendingRewardTotal.Sub(farmInfo.RewardDebt.AmountOf(r.Reward))
			rewards = rewards.Add(sdk.NewCoin(r.Reward, pendingReward))
		}

		locked := farmInfo.Locked.Add(deltaAmt)
		debt := sdk.NewCoin(r.Reward, r.RewardPerShare.MulInt(locked).TruncateInt())
		rewardDebt = rewardDebt.Add(debt)
	}
	return rewards, rewardDebt
}

type RewardRules []RewardRule

func (rs RewardRules) Contains(reward sdk.Coins) bool {
	var allRewards sdk.Coins
	for _, r := range rs {
		allRewards = allRewards.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
	}
	return reward.DenomsSubsetOf(allRewards)
}
