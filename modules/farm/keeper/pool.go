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

	if pool.EndHeight < uint64(ctx.BlockHeight()) {
		return sdkerrors.Wrapf(types.ErrExpiredPool, "not exist pool [%s]", poolName) //TODO
	}

	if err := k.Refund(ctx, poolName, creator); err != nil {
		return sdkerrors.Wrapf(types.ErrNotExistPool, "not exist pool [%s]", poolName)
	}

	pool.EndHeight = uint64(ctx.BlockHeight())
	k.SetPool(ctx, *pool)
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

// AppendReward creates an new farm pool
func (k Keeper) Refund(ctx sdk.Context, poolName string, creator sdk.AccAddress) error {
	rules := k.GetPoolRules(ctx, poolName)
	var remainingReward sdk.Coins

	for _, r := range rules {
		remainingReward = remainingReward.Add(sdk.NewCoin(r.Reward, r.RemainingReward))
		r.RemainingReward = sdk.ZeroInt()
		//update remaining reward
		//TODO update RewardPerShare
		k.SetPoolRule(ctx, poolName, r)
	}

	//refund the total remaining reward to owner
	if err := k.bk.SendCoinsFromModuleToAccount(ctx,
		types.ModuleName, creator, remainingReward); err != nil {
		return err
	}
	return nil
}
