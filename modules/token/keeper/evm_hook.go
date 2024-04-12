package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/irisnet/irismod/contracts"
)

// PostTxProcessing processes the transaction receipt for ERC20 token swap to native.
//
// Parameters:
//   ctx: the context in which the function is executed.
//   msg: the core message associated with the transaction.
//   receipt: the Ethereum receipt containing the transaction logs.
// Return type: error
func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	erc20 := contracts.ERC20TokenContract.ABI
	for _, log := range receipt.Logs {
		// Note: the `SwapToNative` event contains 1 topics
		if len(log.Topics) != 1 {
			continue
		}

		// Check if event is included in ERC20
		eventID := log.Topics[0]
		event, err := erc20.EventByID(eventID)
		if err != nil {
			continue
		}

		// Check if event is a `SwapToNative` event.
		if event.Name != contracts.EventSwapToNative {
			k.Logger(ctx).Info("emitted event", "name", event.Name, "signature", event.Sig)
			continue
		}

		eventArgs, err := erc20.Unpack(event.Name, log.Data)
		if err != nil {
			k.Logger(ctx).Error("failed to unpack SwapToNative event", "error", err.Error())
			continue
		}

		if len(eventArgs) != 3 {
			k.Logger(ctx).Error("invalid SwapToNative event args")
			continue
		}

		sender := sdk.AccAddress(msg.From().Bytes())
		to, ok := eventArgs[0].(string)
		if !ok || len(to) == 0 {
			k.Logger(ctx).Error("invalid SwapToNative event args `to`")
			continue
		}
		toAddr, err := sdk.AccAddressFromBech32(to)
		if err != nil {
			k.Logger(ctx).Error("invalid SwapToNative event args `to`", "error", err.Error())
			continue
		}

		amount, ok := eventArgs[1].(*big.Int)
		if !ok || amount.Cmp(big.NewInt(0)) == 0 {
			k.Logger(ctx).Error("invalid SwapToNative event args `amount`")
			continue
		}

		isERIS, ok := eventArgs[2].(bool)
		if !ok {
			k.Logger(ctx).Error("invalid SwapToNative event args `isERIS`")
			continue
		}

		paid := sdk.NewCoin(k.evmKeeper.FeeDenom(), sdk.NewIntFromBigInt(amount))
		if isERIS {
			_, _, err := k.SwapFeeToken(ctx, paid, sender, toAddr)
			return err
		}
	}
	return nil
}
