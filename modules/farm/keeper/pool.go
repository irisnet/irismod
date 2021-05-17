package keeper

import (
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
	endHeight, err := types.ExpiredHeight(beginHeight, rewardPerBlock, totalReward)
	if err != nil {
		return err
	}

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

	//save farm pool
	k.SetPool(ctx, types.FarmPool{
		Name:         name,
		Creator:      creator.String(),
		BeginHeight:  beginHeight,
		EndHeight:    endHeight,
		Destructible: destructible,
	})
	//save farm rule
	for i, total := range totalReward {
		k.SetPoolRule(ctx, name, types.FarmRule{
			Reward:          total.Denom,
			TotalReward:     total.Amount,
			RemainingReward: total.Amount,
			RewardPerBlock:  rewardPerBlock[i].Amount,
			RewardPerShare:  sdk.ZeroDec(),
		})
	}
	// put to expired farm pool queue
	k.Enqueue(ctx, name, endHeight)
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

	if pool.EndHeight < uint64(ctx.BlockHeight()) {
		return sdkerrors.Wrapf(types.ErrExpiredPool,
			"pool [%s] has expired at height[%d], current [%d]",
			poolName,
			pool.EndHeight,
			ctx.BlockHeight(),
		)
	}

	if err := k.Refund(ctx, pool, creator); err != nil {
		return sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}
	return nil
}

// AppendReward creates an new farm pool
func (k Keeper) AppendReward(ctx sdk.Context, name string,
	reward sdk.Coins,
	creator sdk.AccAddress,
) (remaining sdk.Coins, err error) {
	//TODO
	return nil, nil
}

// AppendReward creates an new farm pool
func (k Keeper) Stake(ctx sdk.Context, name string,
	amount sdk.Coin,
	sender sdk.AccAddress,
) (remaining sdk.Coins, err error) {
	//TODO
	return nil, nil
}

// AppendReward creates an new farm pool
func (k Keeper) Unstake(ctx sdk.Context, name string,
	amount sdk.Coin,
	sender sdk.AccAddress) (remaining sdk.Coins, err error) {
	//TODO
	return nil, nil
}

// AppendReward creates an new farm pool
func (k Keeper) Harvest(ctx sdk.Context, name string,
	creator sdk.AccAddress) (remaining sdk.Coins, err error) {
	//TODO
	return nil, nil
}

// Refund Refund the remaining reward to creator
func (k Keeper) Refund(ctx sdk.Context, pool *types.FarmPool, creator sdk.AccAddress) error {
	rules := k.GetPoolRules(ctx, pool.Name)
	var remainingReward sdk.Coins
	for _, r := range rules {
		rewardAdded := r.RewardPerBlock.ModRaw(ctx.BlockHeight() - int64(pool.LastHeightDistrRewards))
		r.RewardPerShare = sdk.NewDecFromInt(rewardAdded).QuoInt(pool.TotalLpTokenLocked.Amount)
		r.RemainingReward = r.RemainingReward.Sub(rewardAdded)
		remainingReward = remainingReward.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
		k.SetPoolRule(ctx, pool.Name, r)
	}

	//refund the total remaining reward to owner
	if err := k.bk.SendCoinsFromModuleToAccount(ctx,
		types.ModuleName, creator, remainingReward); err != nil {
		return err
	}

	//update LastHeightDistrRewards
	pool.LastHeightDistrRewards = uint64(ctx.BlockHeight())
	//update EndHeight
	pool.EndHeight = uint64(ctx.BlockHeight())
	k.SetPool(ctx, *pool)
	return nil
}
