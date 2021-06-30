package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irisnet/irismod/modules/random/keeper"
	"github.com/irisnet/irismod/modules/random/types"
	"github.com/irisnet/irismod/simapp/helpers"
)

// WeightedOperations generates a MsgRequestRandom with random values.
func WeightedOperations(k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simulation.WeightedOperations {
	var weightMsgRequestRandom = 100
	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgRequestRandom,
			func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
				accs []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

				simAccount, _ := simtypes.RandomAcc(r, accs)
				blockInterval := simtypes.RandIntBetween(r, 10, 100)

				account := ak.GetAccount(ctx, simAccount.Address)

				spendable := bk.SpendableCoins(ctx, account.GetAddress())

				msg := types.NewMsgRequestRandom(simAccount.Address.String(), uint64(blockInterval), false, nil)

				fees, err := simtypes.RandomFees(r, ctx, spendable)
				if err != nil {
					return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
				}

				txGen := simappparams.MakeTestEncodingConfig().TxConfig
				tx, err := helpers.GenTx(
					r,
					txGen,
					[]sdk.Msg{msg},
					fees,
					helpers.DefaultGenTxGas,
					chainID,
					[]uint64{account.GetAccountNumber()},
					[]uint64{account.GetSequence()},
					simAccount.PrivKey,
				)
				if err != nil {
					return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
				}

				if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
					return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
				}

				return simtypes.NewOperationMsg(msg, true, ""), nil, nil
			},
		),
	}
}
