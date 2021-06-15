package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (pool FarmPool) ExpiredHeight() uint64 {
	if len(pool.Rules) == 0 {
		return pool.StartHeight
	}

	targetInteval := pool.Rules[0].TotalReward.Quo(pool.Rules[0].RewardPerBlock).Uint64()
	for _, r := range pool.Rules[1:] {
		inteval := r.TotalReward.Quo(r.RewardPerBlock).Uint64()
		if targetInteval > inteval {
			targetInteval = inteval
		}
	}
	return pool.StartHeight + targetInteval
}

func (pool FarmPool) RemainingHeight() uint64 {
	if len(pool.Rules) == 0 {
		return 0
	}

	targetInteval := pool.Rules[0].RemainingReward.Quo(pool.Rules[0].RewardPerBlock).Uint64()
	for _, r := range pool.Rules[1:] {
		inteval := r.RemainingReward.Quo(r.RewardPerBlock).Uint64()
		if targetInteval > inteval {
			targetInteval = inteval
		}
	}
	return targetInteval
}

func (pool FarmPool) IsExpired(height int64) bool {
	return pool.EndHeight < uint64(height)
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

func (rs RewardRules) UpdateWith(rewarPerBlockd sdk.Coins) RewardRules {
	for _, r := range rs {
		rewardAmt := rewarPerBlockd.AmountOf(r.Reward)
		if rewardAmt.IsPositive() {
			r.RewardPerBlock = rewardAmt
		}
	}
	return rs
}
