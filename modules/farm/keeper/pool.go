package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

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
	pool := types.FarmPool{
		Name:                   name,
		Creator:                creator.String(),
		BeginHeight:            beginHeight,
		EndHeight:              endHeight,
		LastHeightDistrRewards: 0,
		Destructible:           destructible,
	}
	//save farm pool
	k.SetPool(ctx, pool)
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

// Destroy creates an new farm pool
func (k Keeper) DestroyPool(ctx sdk.Context, name string,
	creator sdk.AccAddress) error {
	//TODO
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
