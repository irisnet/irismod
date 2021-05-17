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
) (err error) {
	//Escrow total reward
	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, types.ModuleName, totalReward); err != nil {
		return err
	}

	//send CreatePoolFee to feeCollectorName
	fee := k.CreatePoolFee(ctx)
	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, k.feeCollectorName, sdk.NewCoins(fee)); err != nil {
		return err
	}

	fp := types.FarmPool{
		Name:         name,
		Creator:      creator.String(),
		BeginHeight:  beginHeight,
		Destructible: destructible,
		Rule:         []*types.FarmRule{},
	}
	//save farm rule
	for i, total := range totalReward {
		farmRule := types.FarmRule{
			Reward:          total.Denom,
			TotalReward:     total.Amount,
			RemainingReward: total.Amount,
			RewardPerBlock:  rewardPerBlock[i].Amount,
			RewardPerShare:  sdk.ZeroDec(),
		}
		k.SetPoolRule(ctx, name, farmRule)
		fp.Rule = append(fp.Rule, &farmRule)
	}
	fp.EndHeight = fp.ExpiredHeight()
	//save farm pool
	k.SetPool(ctx, fp)
	// put to expired farm pool queue
	k.EnqueueExpiredPool(ctx, name, fp.EndHeight)
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

	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		creator, types.ModuleName, reward); err != nil {
		return remaining, err
	}

	var heightIncr = uint64(math.MaxUint64)
	for _, r := range k.GetPoolRules(ctx, poolName) {
		r.TotalReward = r.TotalReward.Add(reward.AmountOf(r.Reward))
		r.RemainingReward = r.RemainingReward.Add(reward.AmountOf(r.Reward))
		k.SetPoolRule(ctx, poolName, r)

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
	k.DequeueExpiredPool(ctx, poolName, pool.EndHeight)

	pool.EndHeight = pool.EndHeight + uint64(heightIncr)
	k.SetPool(ctx, *pool)
	// put to expired farm pool queue at new height
	k.EnqueueExpiredPool(ctx, poolName, pool.EndHeight)
	return remaining, nil
}

// Stake is responsible for the user to mortgage the lp token to the system and get back the reward accumulated before then
func (k Keeper) Stake(ctx sdk.Context, poolName string,
	amount sdk.Coin,
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

	if amount.Denom != pool.TotalLpTokenLocked.Denom {
		return reward, sdkerrors.Wrapf(types.ErrNotMatch,
			"pool [%s] only accept [%s] token, but got [%s]",
			poolName, pool.TotalLpTokenLocked.Denom, amount.Denom)
	}

	if err := k.bk.SendCoinsFromAccountToModule(ctx,
		sender, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return reward, err
	}

	//update pool reward shards
	rulesUpdated, err := k.UpdatePool(ctx, pool, amount.Amount)
	if err != nil {
		return nil, err
	}

	farmer, exist := k.GetFarmer(ctx, poolName, sender.String())
	if !exist {
		farmer = &types.Farmer{
			Address:    sender.String(),
			Locked:     sdk.ZeroInt(),
			RewardDebt: sdk.NewCoins(),
		}
	}

	locked := farmer.Locked.Add(amount.Amount)
	for _, r := range rulesUpdated {
		if exist {
			pendingRewardTotal := r.RewardPerShare.MulInt(farmer.Locked).TruncateInt()
			pendingReward := pendingRewardTotal.Sub(farmer.RewardDebt.AmountOf(r.Reward))
			reward = reward.Add(sdk.NewCoin(r.Reward, pendingReward))
		}
		rewardDebt := sdk.NewCoin(r.Reward, r.RewardPerShare.MulInt(locked).TruncateInt())
		farmer.RewardDebt = farmer.RewardDebt.Add(rewardDebt)
	}

	//reward users
	if reward.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, reward); err != nil {
			return reward, err
		}
	}

	farmer.Locked = locked
	k.SetFarmer(ctx, poolName, *farmer)
	return reward, nil
}

// AppendReward creates an new farm pool
func (k Keeper) Unstake(ctx sdk.Context, poolName string,
	amount sdk.Coin,
	sender sdk.AccAddress) (reward sdk.Coins, err error) {
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

	if amount.Denom != pool.TotalLpTokenLocked.Denom {
		return reward, sdkerrors.Wrapf(types.ErrNotMatch,
			"pool [%s] only accept [%s] token, but got [%s]",
			poolName, pool.TotalLpTokenLocked.Denom, amount.Denom)
	}

	farmer, exist := k.GetFarmer(ctx, poolName, sender.String())
	if !exist {
		return reward, sdkerrors.Wrapf(types.ErrNotExistFarmer,
			"farmer [%s] not found in pool[%s]",
			sender.String(),
			poolName,
		)
	}

	if farmer.Locked.LT(amount.Amount) {
		return reward, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"farmer locked lp token %s, but unstake %s",
			farmer.Locked.String(),
			amount.Amount.String(),
		)
	}

	locked := farmer.Locked.Sub(amount.Amount)

	//update pool reward shards
	rulesUpdated, err := k.UpdatePool(ctx, pool, amount.Amount.Neg())
	if err != nil {
		return nil, err
	}

	for _, r := range rulesUpdated {
		pendingRewardTotal := r.RewardPerShare.MulInt(farmer.Locked).TruncateInt()
		pendingReward := pendingRewardTotal.Sub(farmer.RewardDebt.AmountOf(r.Reward))
		reward = reward.Add(sdk.NewCoin(r.Reward, pendingReward))

		rewardDebt := sdk.NewCoin(r.Reward, r.RewardPerShare.MulInt(locked).TruncateInt())
		farmer.RewardDebt = farmer.RewardDebt.Add(rewardDebt)
	}

	//reward users
	if reward.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, reward); err != nil {
			return reward, err
		}
	}

	if locked.IsZero() {
		k.DeleteFarmer(ctx, poolName, sender.String())
		return reward, nil
	}

	farmer.Locked = locked
	k.SetFarmer(ctx, poolName, *farmer)
	return reward, nil
}

// Harvest creates an new farm pool
func (k Keeper) Harvest(ctx sdk.Context, poolName string,
	sender sdk.AccAddress) (reward sdk.Coins, err error) {
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

	farmer, exist := k.GetFarmer(ctx, poolName, sender.String())
	if !exist {
		return reward, sdkerrors.Wrapf(types.ErrNotExistFarmer,
			"farmer [%s] not found in pool[%s]",
			sender.String(),
			poolName,
		)
	}

	//update pool reward shards
	rulesUpdated, err := k.UpdatePool(ctx, pool, sdk.ZeroInt())
	if err != nil {
		return nil, err
	}

	for _, r := range rulesUpdated {
		pendingRewardTotal := r.RewardPerShare.MulInt(farmer.Locked).TruncateInt()
		pendingReward := pendingRewardTotal.Sub(farmer.RewardDebt.AmountOf(r.Reward))
		reward = reward.Add(sdk.NewCoin(r.Reward, pendingReward))

		rewardDebt := sdk.NewCoin(r.Reward, r.RewardPerShare.MulInt(farmer.Locked).TruncateInt())
		farmer.RewardDebt = farmer.RewardDebt.Add(rewardDebt)
	}

	//reward users
	if reward.IsAllPositive() {
		if err = k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, reward); err != nil {
			return reward, err
		}
	}
	k.SetFarmer(ctx, poolName, *farmer)
	return reward, nil
}

// Refund refund the remaining reward to pool creator
func (k Keeper) Refund(ctx sdk.Context, pool *types.FarmPool) error {
	rules := k.GetPoolRules(ctx, pool.Name)
	var remainingTotal sdk.Coins

	blockInterval := uint64(ctx.BlockHeight()) - pool.LastHeightDistrRewards
	for _, r := range rules {
		rewardAdded := r.RewardPerBlock.ModRaw(int64(blockInterval))
		newRewardPerShare := sdk.NewDecFromInt(rewardAdded).QuoInt(pool.TotalLpTokenLocked.Amount)
		r.RewardPerShare = r.RewardPerShare.Add(newRewardPerShare)
		r.RemainingReward = r.RemainingReward.Sub(rewardAdded)
		remainingTotal = remainingTotal.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
		k.SetPoolRule(ctx, pool.Name, r)
	}

	creator, err := sdk.AccAddressFromBech32(pool.Creator)
	if err != nil {
		return err
	}

	//refund the total remaining reward to creator
	if err := k.bk.SendCoinsFromModuleToAccount(ctx,
		types.ModuleName, creator, remainingTotal); err != nil {
		return err
	}

	//remove record
	k.DequeueExpiredPool(ctx, pool.Name, pool.EndHeight)

	//update LastHeightDistrRewards
	pool.LastHeightDistrRewards = uint64(ctx.BlockHeight())
	//update EndHeight
	pool.EndHeight = uint64(ctx.BlockHeight())
	k.SetPool(ctx, *pool)
	return nil
}

func (k Keeper) UpdatePool(ctx sdk.Context, pool *types.FarmPool, amount sdk.Int) ([]types.FarmRule, error) {
	height := uint64(ctx.BlockHeight())
	if height < pool.LastHeightDistrRewards {
		return nil, sdkerrors.Wrapf(types.ErrExpiredHeight, "invalid height: %d, current: %d", height, pool.LastHeightDistrRewards)
	}

	rules := k.GetPoolRules(ctx, pool.Name)
	if len(rules) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrNotExistPool, "the rule of the farm pool[%s] not exist", pool.Name)
	}

	if height == pool.LastHeightDistrRewards {
		return rules, nil
	}

	if pool.TotalLpTokenLocked.IsZero() {
		return rules, nil
	}

	blockInterval := height - pool.LastHeightDistrRewards
	for _, r := range rules {
		rewardAdded := r.RewardPerBlock.ModRaw(int64(blockInterval))
		newRewardPerShare := sdk.NewDecFromInt(rewardAdded).QuoInt(pool.TotalLpTokenLocked.Amount)
		r.RewardPerShare = r.RewardPerShare.Add(newRewardPerShare)
		r.RemainingReward = r.RemainingReward.Sub(rewardAdded)
		k.SetPoolRule(ctx, pool.Name, r)
	}

	if amount.IsZero() {
		return rules, nil
	}

	totalLocked := pool.TotalLpTokenLocked.Amount.Add(amount)
	pool.TotalLpTokenLocked = sdk.NewCoin(pool.TotalLpTokenLocked.Denom, totalLocked)
	pool.LastHeightDistrRewards = uint64(ctx.BlockHeight())
	k.SetPool(ctx, *pool)
	return rules, nil
}
