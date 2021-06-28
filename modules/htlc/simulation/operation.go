package simulation

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/irisnet/irismod/modules/htlc/keeper"
	"github.com/irisnet/irismod/modules/htlc/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"math/rand"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateHTLC = "op_weight_msg_create_htlc"
	OpWeightMsgClaimHTLC  = "op_weight_msg_claim_htlc"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simulation.WeightedOperations {

	var weightCreateHtlc, weightClaimHtlc int

	appParams.GetOrGenerate(
		cdc, OpWeightMsgCreateHTLC, &weightCreateHtlc, nil,
		func(_ *rand.Rand) {
			weightCreateHtlc = 100
		},
	)
	appParams.GetOrGenerate(
		cdc, OpWeightMsgClaimHTLC, &weightClaimHtlc, nil,
		func(_ *rand.Rand) {
			weightClaimHtlc = 50
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightCreateHtlc,
			SimulateMsgCreateHtlc(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightClaimHtlc,
			SimulateMsgClaimHtlc(k, ak, bk),
		),
	}
}

func SimulateMsgCreateHtlc(ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		sender, _ := simtypes.RandomAcc(r, accs)
		to, _ := simtypes.RandomAcc(r, accs)
		recvOnOtherChain, _ := simtypes.RandomAcc(r, accs)
		senderOnOtherChain, _ := simtypes.RandomAcc(r, accs)

		account := ak.GetAccount(ctx, sender.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		if spendable.IsZero() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, "Insufficient funds"), nil, nil
		}
		amount := simtypes.RandSubsetCoins(r, spendable)

		balance, hasNeg := spendable.SafeSub(amount)
		if hasNeg {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, "Insufficient funds"), nil, nil
		}
		hashLock := "0000000000000000000000000000000000000000000000000000000000000000"
		//hashLock := GenHashLock(r)
		//tt := hex.EncodeToString(hashLock)
		timeLock := uint64(220)
		tranfer := false
		msg := &types.MsgCreateHTLC{
			Sender:               sender.Address.String(),
			To:                   to.Address.String(),
			ReceiverOnOtherChain: recvOnOtherChain.Address.String(),
			SenderOnOtherChain:   senderOnOtherChain.Address.String(),
			Amount:               amount,
			HashLock:             hashLock,
			Timestamp:            uint64(0),
			TimeLock:             timeLock,
			Transfer:             tranfer,
		}

		fees, err := simtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, err.Error()), nil, err
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
			sender.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}
func SimulateMsgClaimHtlc(k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		htlc := genRandomHtlc(ctx, k, r)
		if htlc.Size() == 0{
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, "not exist htlc"), nil, nil
		}
		sender, err := sdk.AccAddressFromBech32(htlc.Sender)
		if err != nil{
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, "invalid address"), nil, nil
		}
		account := ak.GetAccount(ctx, sender)

		simAccount, found := simtypes.FindAccount(accs, account.GetAddress())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, "account not found"), nil, nil
		}

		secret, err := GenRandomSecret()
		if err != nil{
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, err.Error()), nil, err
		}
		msg := &types.MsgClaimHTLC{Sender: htlc.Sender, Id: htlc.Id, Secret: string(secret)}

		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateHTLC, err.Error()), nil, err
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

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

//func  GenHashLock(r *rand.Rand) []byte{
//	hashlock ,_ := hex.DecodeString(simtypes.RandStringOfLength(r, 64))
//	return hashlock
//}


func genRandomHtlc(ctx sdk.Context, k keeper.Keeper, r *rand.Rand)types.HTLC{
	htlcs := []types.HTLC{}
	k.IterateHTLCs(
		ctx,
		func(_ tmbytes.HexBytes, h types.HTLC) (stop bool) {
			htlcs = append(htlcs, h)
			return false
		},
	)
	if len(htlcs) == 0{
		return types.HTLC{}
	}
	return htlcs[r.Intn(len(htlcs))]
}


func GenRandomSecret() ([]byte, error) {
	bytes := make([]byte, 64)
	if _, err := rand.Read(bytes); err != nil {
		return []byte{}, err
	}
	return bytes, nil
}


