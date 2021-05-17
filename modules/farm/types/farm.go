package types

import (
	math "math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (fp FarmPool) ExpiredHeight() uint64 {
	var expiredHeight = uint64(math.MaxUint64)
	for _, r := range fp.Rules {
		inteval := r.TotalReward.Quo(r.RewardPerBlock).Uint64()
		if inteval+fp.BeginHeight < expiredHeight {
			expiredHeight = inteval + fp.BeginHeight
		}
	}
	return expiredHeight + 1
}

func (fp FarmPool) CaclRewards(farmer Farmer, deltaAmt sdk.Int) (rewards, dewardDebt sdk.Coins) {
	dewardDebt = farmer.RewardDebt
	for _, r := range fp.Rules {
		if farmer.Locked.GT(sdk.ZeroInt()) {
			pendingRewardTotal := r.RewardPerShare.MulInt(farmer.Locked).TruncateInt()
			pendingReward := pendingRewardTotal.Sub(farmer.RewardDebt.AmountOf(r.Reward))
			rewards = rewards.Add(sdk.NewCoin(r.Reward, pendingReward))
		}

		locked := farmer.Locked.Add(deltaAmt)
		rewardDebt := sdk.NewCoin(r.Reward, r.RewardPerShare.MulInt(locked).TruncateInt())
		dewardDebt = dewardDebt.Add(rewardDebt)
	}
	return rewards, dewardDebt
}
