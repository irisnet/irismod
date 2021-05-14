package types

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func ExpiredHeight(begin uint64,
	rewardPerBlock sdk.Coins,
	totalReward sdk.Coins) (uint64, error) {
	var expiredHeight = uint64(math.MaxUint64)
	if len(rewardPerBlock) != len(totalReward) {
		return expiredHeight, sdkerrors.Wrapf(ErrNotMatch, "The length of `rewardPerBlock` and `totalReward` must be the same")
	}

	for i, coin := range totalReward {
		perBlock := rewardPerBlock[i]
		if perBlock.Denom != coin.Denom {
			return expiredHeight, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "not sort")
		}
		inteval := coin.Amount.Quo(perBlock.Amount).Uint64()
		if inteval+begin < expiredHeight {
			expiredHeight = inteval + begin
		}
	}
	return expiredHeight, nil
}
