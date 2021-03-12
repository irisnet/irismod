package htlc

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/irisnet/irismod/modules/htlc/keeper"
	"github.com/irisnet/irismod/modules/htlc/types"
)

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	k.SetPreviousBlockTime(ctx, data.PreviousBlockTime)
	k.SetParams(ctx, data.Params)

	// TODO
	for _, htlc := range data.PendingHtlcs {
		id, err := hex.DecodeString(htlc.Id)
		if err != nil {
			panic(err.Error())
		}

		k.SetHTLC(ctx, htlc, id)
		k.AddHTLCToExpiredQueue(ctx, htlc.ExpirationHeight, id)
	}

	// TODO
	for _, supply := range data.Supplies {
		k.SetAssetSupply(ctx, supply, supply.CurrentSupply.Denom)
	}
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	pendingHTLCs := []types.HTLC{}
	k.IterateHTLCs(
		ctx,
		func(_ tmbytes.HexBytes, h types.HTLC) (stop bool) {
			if h.State == types.Open {
				pendingHTLCs = append(pendingHTLCs, h)
			}
			return false
		},
	)

	supplies := k.GetAllAssetSupplies(ctx)
	previousBlockTime, found := k.GetPreviousBlockTime(ctx)
	if !found {
		previousBlockTime = types.DefaultPreviousBlockTime
	}

	return types.NewGenesisState(
		k.GetParams(ctx),
		pendingHTLCs,
		supplies,
		previousBlockTime,
	)
}

func PrepForZeroHeightGenesis(ctx sdk.Context, k keeper.Keeper) {
	k.IterateHTLCs(
		ctx,
		func(id tmbytes.HexBytes, h types.HTLC) (stop bool) {
			if h.State == types.Open {
				h.ExpirationHeight = h.ExpirationHeight - uint64(ctx.BlockHeight()) + 1
				k.SetHTLC(ctx, h, id)
			}
			return false
		},
	)
	// TODO: update asset supplies and previous block time
}
