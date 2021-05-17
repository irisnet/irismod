package farm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/farm/keeper"
	"github.com/irisnet/irismod/modules/farm/types"
)

// BeginBlocker handles block beginning logic for farm
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	ctx = ctx.WithLogger(ctx.Logger().
		With("handler", "beginBlocker").
		With("module", "irismod/farm"),
	)

	k.IteratorExpiredPool(ctx, func(pool *types.FarmPool) {
		ctx.Logger().Info(
			"The farm pool has expired, refund to creator",
			"poolName", pool.Name,
			"endHeight", pool.EndHeight,
			"lastHeightDistrRewards", pool.LastHeightDistrRewards,
			"totalLpTokenLocked", pool.TotalLpTokenLocked,
		)
		if err := k.Refund(ctx, pool); err != nil {
			panic(err)
		}
	})
}
