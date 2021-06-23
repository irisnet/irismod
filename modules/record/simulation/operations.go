package simulation

import (
	"fmt"
	"math/rand"

	"github.com/irisnet/irismod/modules/record/keeper"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irisnet/irismod/modules/record/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateRecord = "op_weight_msg_create_record"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper) simulation.WeightedOperations {
	var weightCreate int
	appParams.GetOrGenerate(
		cdc, OpWeightMsgCreateRecord, &weightCreate, nil,
		func(_ *rand.Rand) {
			weightCreate = 50
		},
	)
	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightCreate,
			SimulateCreateRecord(ak, bk, k),
		),
	}
}

// SimulateCreateRecord tests and runs a single msg create a new record
func SimulateCreateRecord(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		record := genRandomRecord(r, k, ctx)
		if record.Size() == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.EventTypeCreateRecord, "record not exist"), nil, nil
		}

		creator, _ := sdk.AccAddressFromBech32(record.Creator)
		msg := types.NewMsgCreateRecord(record.Contents, creator.String())

		simAccount, found := simtypes.FindAccount(accs, creator)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.EventTypeCreateRecord, "creator not found"), nil, fmt.Errorf("account %s not found", record.Creator)
		}

		account := ak.GetAccount(ctx, creator)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.EventTypeCreateRecord, err.Error()), nil, err
		}
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, _ := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			simAccount.PrivKey,
		)

		if _, _, err = app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.EventTypeCreateRecord, err.Error()), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "simulate issue token"), nil, nil
	}
}

func genRecord(r *rand.Rand, accs []simtypes.Account) (types.Record, error) {
	var record types.Record
	txHash := make([]byte, 32)
	if _, err := r.Read(txHash); err != nil {
		return record, err
	}

	record.TxHash = tmbytes.HexBytes(txHash).String()

	for i := 0; i <= r.Intn(10); i++ {
		record.Contents = append(record.Contents, types.Content{
			Digest:     simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 1, 50)),
			DigestAlgo: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 1, 50)),
			URI:        simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 0, 50)),
			Meta:       simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 0, 50)),
		})
	}

	acc, _ := simtypes.RandomAcc(r, accs)
	record.Creator = acc.Address.String()

	return record, nil
}

func genRandomRecord(r *rand.Rand, k keeper.Keeper, ctx sdk.Context) types.Record {
	recordsIterator := k.RecordsIterator(ctx)
	defer recordsIterator.Close()

	var records []types.Record
	for ; recordsIterator.Valid(); recordsIterator.Next() {
		var record types.Record
		types.ModuleCdc.MustUnmarshalBinaryBare(recordsIterator.Value(), &record)
		records = append(records, record)
	}
	if len(records) > 0 {
		return records[r.Intn(len(records))]
	}
	return types.Record{}
}
