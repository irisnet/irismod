package farm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/farm/keeper"
)

// BeginBlocker handles block beginning logic for farm
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	ctx = ctx.WithLogger(ctx.Logger().
		With("handler", "beginBlocker").
		With("module", "irismod/farm"),
	)

	for _, pool := range k.IteratorExpiredPool(ctx) {
		if err := k.Refund(ctx, pool); err != nil {
			panic(err)
		}
	}
}
