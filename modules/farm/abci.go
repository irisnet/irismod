package farm

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"mods.irisnet.org/modules/farm/keeper"
	"mods.irisnet.org/modules/farm/types"
)

// EndBlocker handles block beginning logic for farm
func EndBlocker(c context.Context, k keeper.Keeper) {
	ctx := sdk.UnwrapSDKContext(c)
	logger := k.Logger(ctx).With("handler", "endBlocker")
	k.IteratorExpiredPool(ctx, ctx.BlockHeight(), func(pool types.FarmPool) {
		logger.Info(
			"The farm pool has expired, refund to creator",
			"poolId", pool.Id,
			"endHeight", pool.EndHeight,
			"lastHeightDistrRewards", pool.LastHeightDistrRewards,
			"totalLptLocked", pool.TotalLptLocked,
			"creator", pool.Creator,
		)
		if _, err := k.Refund(ctx, pool); err != nil {
			logger.Error("The farm pool refund failed",
				"poolId", pool.Id,
				"creator", pool.Creator,
				"errMsg", err.Error(),
			)
		}
	})
}
