package keeper

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irisnet/irismod/modules/farm/types"
)

// CreatePool creates an new farm pool
func (k Keeper) CreatePool(ctx sdk.Context, name string,
	description string,
	lpTokenDenom string,
	beginHeight uint64,
	rewardPerBlock sdk.Coins,
	totalReward sdk.Coins,
	destructible bool,
	creator sdk.AccAddress,
) error {
	//Escrow total reward
	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, types.ModuleName, totalReward); err != nil {
		return err
	}

	//send CreatePoolFee to feeCollectorName
	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, k.feeCollectorName, sdk.NewCoins(k.CreatePoolFee(ctx))); err != nil {
		return err
	}

	pool := types.FarmPool{
		Name:               name,
		Creator:            creator.String(),
		Description:        description,
		BeginHeight:        beginHeight,
		Destructible:       destructible,
		TotalLpTokenLocked: sdk.NewCoin(lpTokenDenom, sdk.ZeroInt()),
		Rules:              []*types.RewardRule{},
	}
	//save farm rule
	for i, total := range totalReward {
		rewardRule := types.RewardRule{
			Reward:          total.Denom,
			TotalReward:     total.Amount,
			RemainingReward: total.Amount,
			RewardPerBlock:  rewardPerBlock[i].Amount,
			RewardPerShare:  sdk.ZeroDec(),
		}
		k.SetRewardRule(ctx, name, rewardRule)
		pool.Rules = append(pool.Rules, &rewardRule)
	}
	pool.EndHeight = pool.ExpiredHeight()
	//save farm pool
	k.SetPool(ctx, pool)
	// put to expired farm pool queue
	k.EnqueueActivePool(ctx, name, pool.EndHeight)
	return nil
}

// Destroy destroy an exist farm pool
func (k Keeper) DestroyPool(ctx sdk.Context, poolName string,
	creator sdk.AccAddress) error {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}

	if creator.String() != pool.Creator {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "creator [%s] is not the creator of the pool", creator.String())
	}

	if !pool.Destructible {
		return sdkerrors.Wrapf(
			types.ErrInvalidOperate, "pool [%s] is not destructible", poolName)
	}

	if pool.EndHeight <= uint64(ctx.BlockHeight()) {
		return sdkerrors.Wrapf(types.ErrExpiredPool,
			"pool [%s] has expired at height[%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}

	if err := k.Refund(ctx, pool); err != nil {
		return sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}
	return nil
}

// AppendReward creates an new farm pool
func (k Keeper) AppendReward(ctx sdk.Context, poolName string,
	reward sdk.Coins,
	creator sdk.AccAddress,
) (remaining sdk.Coins, err error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return remaining, sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}

	if creator.String() != pool.Creator {
		return remaining, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "creator [%s] is not the creator of the pool", creator.String())
	}

	if pool.EndHeight <= uint64(ctx.BlockHeight()) {
		return remaining, sdkerrors.Wrapf(types.ErrExpiredPool,
			"pool [%s] has expired at height[%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}

	rules := k.GetRewardRules(ctx, poolName)
	if !rules.Contains(reward) {
		return remaining, sdkerrors.Wrapf(types.ErrInvalidAppend, reward.String())
	}

	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, types.ModuleName, reward); err != nil {
		return remaining, err
	}

	var heightIncr = uint64(math.MaxUint64)
	for _, r := range rules {
		r.TotalReward = r.TotalReward.Add(reward.AmountOf(r.Reward))
		r.RemainingReward = r.RemainingReward.Add(reward.AmountOf(r.Reward))
		k.SetRewardRule(ctx, poolName, *r)

		delta := reward.AmountOf(r.Reward).Quo(r.RewardPerBlock).Uint64()
		if delta < heightIncr {
			heightIncr = delta
		}

		remaining = remaining.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
	}

	//if the expiration height does not change, there is no need to update the pool and the expired queue
	if heightIncr == 0 {
		return remaining, nil
	}

	// remove from Expired Pool at old height
	k.DequeueActivePool(ctx, poolName, pool.EndHeight)

	pool.EndHeight = pool.EndHeight + uint64(heightIncr)
	k.SetPool(ctx, *pool)
	// put to expired farm pool queue at new height
	k.EnqueueActivePool(ctx, poolName, pool.EndHeight)
	return remaining, nil
}

// Stake is responsible for the user to mortgage the lp token to the system and get back the reward accumulated before then
func (k Keeper) Stake(ctx sdk.Context, poolName string,
	lpToken sdk.Coin,
	sender sdk.AccAddress,
) (reward sdk.Coins, err error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return reward, sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}

	if pool.EndHeight <= uint64(ctx.BlockHeight()) {
		return reward, sdkerrors.Wrapf(types.ErrExpiredPool,
			"pool [%s] has expired at height[%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}

	if lpToken.Denom != pool.TotalLpTokenLocked.Denom {
		return reward, sdkerrors.Wrapf(types.ErrNotMatch,
			"pool [%s] only accept [%s] token, but got [%s]",
			poolName, pool.TotalLpTokenLocked.Denom, lpToken.Denom)
	}

	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		sender, types.ModuleName, sdk.NewCoins(lpToken)); err != nil {
		return reward, err
	}

	farmInfo, exist := k.GetFarmInfo(ctx, poolName, sender.String())
	if !exist {
		farmInfo = types.FarmInfo{
			PoolName:   poolName,
			Address:    sender.String(),
			Locked:     sdk.ZeroInt(),
			RewardDebt: sdk.NewCoins(),
		}
	}

	//update pool reward shards
	pool, err = k.UpdatePool(ctx, pool, lpToken.Amount, false)
	if err != nil {
		return nil, err
	}
	rewards, rewardDebt := pool.CaclRewards(farmInfo, lpToken.Amount)

	//reward users
	if reward.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, rewards); err != nil {
			return reward, err
		}
	}

	farmInfo.RewardDebt = rewardDebt
	farmInfo.Locked = farmInfo.Locked.Add(lpToken.Amount)
	k.SetFarmInfo(ctx, farmInfo)
	return rewards, nil
}

// AppendReward creates an new farm pool
func (k Keeper) Unstake(ctx sdk.Context, poolName string,
	lpToken sdk.Coin,
	sender sdk.AccAddress) (sdk.Coins, error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}

	if pool.EndHeight <= uint64(ctx.BlockHeight()) {
		return nil, sdkerrors.Wrapf(types.ErrExpiredPool,
			"pool [%s] has expired at height[%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}

	if lpToken.Denom != pool.TotalLpTokenLocked.Denom {
		return nil, sdkerrors.Wrapf(types.ErrNotMatch,
			"pool [%s] only accept [%s] token, but got [%s]",
			poolName, pool.TotalLpTokenLocked.Denom, lpToken.Denom)
	}

	farmInfo, exist := k.GetFarmInfo(ctx, poolName, sender.String())
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistFarmer,
			"farmer [%s] not found in pool[%s]",
			sender.String(),
			poolName,
		)
	}

	if farmInfo.Locked.LT(lpToken.Amount) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"farmer locked lp token %s, but unstake %s",
			farmInfo.Locked.String(),
			lpToken.Amount.String(),
		)
	}

	amtAdded := lpToken.Amount.Neg()
	//update pool reward shards
	pool, err := k.UpdatePool(ctx, pool, amtAdded, false)
	if err != nil {
		return nil, err
	}

	rewards, rewardDebt := pool.CaclRewards(farmInfo, amtAdded)
	//reward users
	if rewards.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, rewards); err != nil {
			return nil, err
		}
	}

	farmInfo.RewardDebt = rewardDebt
	farmInfo.Locked = farmInfo.Locked.Add(amtAdded)

	if farmInfo.Locked.IsZero() {
		k.DeleteFarmInfo(ctx, poolName, sender.String())
		return rewards, nil
	}

	k.SetFarmInfo(ctx, farmInfo)
	return rewards, nil
}

// Harvest creates an new farm pool
func (k Keeper) Harvest(ctx sdk.Context, poolName string,
	sender sdk.AccAddress) (sdk.Coins, error) {
	pool, exist := k.GetPool(ctx, poolName)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}

	if pool.EndHeight <= uint64(ctx.BlockHeight()) {
		return nil, sdkerrors.Wrapf(types.ErrExpiredPool,
			"pool [%s] has expired at height[%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}

	farmInfo, exist := k.GetFarmInfo(ctx, poolName, sender.String())
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrNotExistFarmer,
			"farmer [%s] not found in pool[%s]",
			sender.String(),
			poolName,
		)
	}

	amtAdded := sdk.ZeroInt()
	//update pool reward shards
	pool, err := k.UpdatePool(ctx, pool, amtAdded, false)
	if err != nil {
		return nil, err
	}

	rewards, rewardDebt := pool.CaclRewards(farmInfo, amtAdded)
	//reward users
	if rewards.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, rewards); err != nil {
			return nil, err
		}
	}

	farmInfo.RewardDebt = rewardDebt
	k.SetFarmInfo(ctx, farmInfo)
	return rewards, nil
}

// Refund refund the remaining reward to pool creator
func (k Keeper) Refund(ctx sdk.Context, pool *types.FarmPool) (err error) {
	pool, err = k.UpdatePool(ctx, pool, sdk.ZeroInt(), true)
	if err != nil {
		return err
	}

	creator, err := sdk.AccAddressFromBech32(pool.Creator)
	if err != nil {
		return err
	}

	var remainingTotal sdk.Coins
	for _, r := range pool.Rules {
		remainingTotal = remainingTotal.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
	}

	//refund the total remaining reward to creator
	if err := k.bk.SendCoinsFromModuleToAccount(ctx,
		types.ModuleName, creator, remainingTotal); err != nil {
		return err
	}

	//remove record
	k.DequeueActivePool(ctx, pool.Name, pool.EndHeight)
	return nil
}

func (k Keeper) UpdatePool(ctx sdk.Context,
	pool *types.FarmPool,
	amount sdk.Int,
	isDestroy bool,
) (*types.FarmPool, error) {
	height := uint64(ctx.BlockHeight())
	if height < pool.LastHeightDistrRewards {
		return nil, sdkerrors.Wrapf(types.ErrExpiredHeight, "invalid height: %d, current: %d", height, pool.LastHeightDistrRewards)
	}

	rules := k.GetRewardRules(ctx, pool.Name)
	if len(rules) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrNotExistPool, "the rule of the farm pool[%s] not exist", pool.Name)
	}

	//when there are multiple farm operations in the same block, the value needs to be updated once
	if height > pool.LastHeightDistrRewards &&
		pool.TotalLpTokenLocked.Amount.GT(sdk.ZeroInt()) {
		blockInterval := height - pool.LastHeightDistrRewards
		for _, r := range rules {
			rewardAdded := r.RewardPerBlock.MulRaw(int64(blockInterval))
			if r.RemainingReward.LT(rewardAdded) {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
					"the remaining reward of the pool[%s] is %s, but got %s",
					pool.Name,
					sdk.NewCoin(r.Reward, r.RemainingReward).String(),
					sdk.NewCoin(r.Reward, rewardAdded).String(),
				)
			}
			newRewardPerShare := sdk.NewDecFromInt(rewardAdded).QuoInt(pool.TotalLpTokenLocked.Amount)
			r.RewardPerShare = r.RewardPerShare.Add(newRewardPerShare)
			r.RemainingReward = r.RemainingReward.Sub(rewardAdded)
			k.SetRewardRule(ctx, pool.Name, *r)
		}
	}

	pool.TotalLpTokenLocked = sdk.NewCoin(
		pool.TotalLpTokenLocked.Denom,
		pool.TotalLpTokenLocked.Amount.Add(amount),
	)
	pool.LastHeightDistrRewards = uint64(ctx.BlockHeight())
	if isDestroy {
		pool.EndHeight = uint64(ctx.BlockHeight())
	}
	k.SetPool(ctx, *pool)

	pool.Rules = rules
	return pool, nil
}
