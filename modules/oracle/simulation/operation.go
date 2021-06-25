package simulation

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/irisnet/irismod/modules/oracle/keeper"
	"github.com/irisnet/irismod/modules/oracle/types"
	"math/rand"
)


const (
	OpWeightMsgCreateFeed = "op_weight_msg_create_feed"
	OpWeightMsgPauseFeed  = "op_weight_msg_pause_feed"
	OpWeightMsgStartFeed  = "op_weight_msg_start_feed"
	OpWeightMsgEditFeed   = "op_weight_msg_edit_feed"
)


func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	k keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper) simulation.WeightedOperations {

	var weightCreate, weightPause, weightStart, WeightEdit int

	appParams.GetOrGenerate(
		cdc, OpWeightMsgCreateFeed, &weightCreate, nil,
		func(_ *rand.Rand) {
			weightCreate = 50
		},
	)
	appParams.GetOrGenerate(
		cdc, OpWeightMsgPauseFeed, &weightPause, nil,
		func(_ *rand.Rand) {
			weightPause = 50
		},
	)
	appParams.GetOrGenerate(
		cdc, OpWeightMsgStartFeed, &weightStart, nil,
		func(_ *rand.Rand) {
			weightStart = 50
		},
	)
	appParams.GetOrGenerate(
		cdc, OpWeightMsgEditFeed, &WeightEdit, nil,
		func(_ *rand.Rand) {
			WeightEdit = 50
		},
	)
	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightCreate,
			SimulateCreateFeed(k, ak, bk),
		),
	}
}

func SimulateCreateFeed(k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (
		opMsg simtypes.OperationMsg, fOps []simtypes.FutureOperation, err error,
	) {

		providers1 ,_ :=simtypes.RandomAcc(r, accs)
		providers2 ,_ :=simtypes.RandomAcc(r, accs)

		simAccount, _ 	  := simtypes.RandomAcc(r, accs)
		feedName 	      := simtypes.RandStringOfLength(r, 10)
		latestHistory     := uint64(simtypes.RandIntBetween(r, 1, 100))
		description       := simtypes.RandStringOfLength(r, 50)
		creator   	      := simAccount.Address.String()
		serviceName       := simtypes.RandStringOfLength(r, 10) //
		providers         := []string{providers1.Address.String(), providers2.Address.String()} //
		input 		  	  := `{"header":{},"body":{}}` //
		timeout 	  	  := int64(simtypes.RandIntBetween(r, 10, 100))
		srvFeeCap 	  	  := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 2, 10))))}
		repeatedFrequency := uint64(100)
		aggregateFunc     := GenAggregateFunc(r)
		valueJsonPath     := `{"input":{"type":"object"},"output":{"type":"object"}}`
		responseThreshold := uint32(simtypes.RandIntBetween(r, 1, len(providers)))

		msg := &types.MsgCreateFeed{
			feedName,
			latestHistory,
			description,
			creator,
			serviceName,
			providers,
			input,
			timeout,
			srvFeeCap,
			repeatedFrequency,
			aggregateFunc,
			valueJsonPath,
			responseThreshold,
		}
		account := ak.GetAccount(ctx, simAccount.Address)

		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.EventTypeCreateFeed, err.Error()), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
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

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.EventTypeCreateFeed, err.Error()), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func GenAggregateFunc(r *rand.Rand) string {
	slice := []string{"max", "min", "avg"}
	return slice[r.Intn(len(slice))]
}
