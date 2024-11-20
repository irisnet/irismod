package htlc

import (
	"context"
	"fmt"

	tmbytes "github.com/cometbft/cometbft/libs/bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"mods.irisnet.org/modules/htlc/keeper"
	"mods.irisnet.org/modules/htlc/types"
)

// BeginBlocker handles block beginning logic for HTLC
func BeginBlocker(c context.Context, k keeper.Keeper) error {
	ctx := sdk.UnwrapSDKContext(c)
	ctx = ctx.WithLogger(ctx.Logger().With("handler", "beginBlock").With("module", "irismod/htlc"))

	currentBlockHeight := uint64(ctx.BlockHeight())
	k.IterateHTLCExpiredQueueByHeight(
		ctx, currentBlockHeight,
		func(id tmbytes.HexBytes, h types.HTLC) (stop bool) {
			// refund HTLC
			if err := k.RefundHTLC(ctx, h, id); err != nil {
				return err
			}
			// delete from the expiration queue
			k.DeleteHTLCFromExpiredQueue(ctx, currentBlockHeight, id)

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeRefundHTLC,
					sdk.NewAttribute(types.AttributeKeyID, id.String()),
				),
			})

			ctx.Logger().Info(fmt.Sprintf("HTLC [%s] is refunded", id.String()))

			return false
		},
	)

	k.UpdateTimeBasedSupplyLimits(ctx)

	return nil
}
