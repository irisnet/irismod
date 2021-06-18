package simulation

import (
	"fmt"
	"math/rand"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	cosmossimappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irisnet/irismod/modules/service/keeper"
	"github.com/irisnet/irismod/modules/service/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgDefineService         = "op_weight_msg_define_service"
	OpWeightMsgBindService           = "op_weight_msg_bind_service"
	OpWeightMsgUpdateServiceBinding  = "op_weight_msg_update_service_binding"
	OpWeightMsgSetWithdrawAddress    = "op_weight_msg_set_withdraw_address"
	OpWeightMsgDisableServiceBinding = "op_weight_msg_disable_service_binding"
	OpWeightMsgEnableServiceBinding  = "op_weight_msg_enable_service_binding"
	OpWeightMsgRefundServiceDeposit  = "op_weight_msg_refund_service_deposit"
	OpWeightMsgCallService           = "op_weight_msg_call_service"
	OpWeightMsgRespondService        = "op_weight_msg_respond_service"
	OpWeightMsgStartRequestContext   = "op_weight_msg_start_request_context"
	OpWeightMsgPauseRequestContext   = "op_weight_msg_pause_request_context"
	OpWeightMsgKillRequestContext    = "op_weight_msg_kill_request_context"
	OpWeightMsgUpdateRequestContext  = "op_weight_msg_update_request_context"
	OpWeightMsgWithdrawEarnedFees    = "op_weight_msg_withdraw_earned_fees"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgDefineService         int
		weightMsgBindService           int
		weightMsgUpdateServiceBinding  int
		weightMsgSetWithdrawAddress    int
		weightMsgDisableServiceBinding int
		weightMsgEnableServiceBinding  int
		weightMsgRefundServiceDeposit  int
		weightMsgCallService           int
		weightMsgRespondService        int
		weightMsgStartRequestContext   int
		weightMsgPauseRequestContext   int
		weightMsgKillRequestContext    int
		weightMsgUpdateRequestContext  int
		weightMsgWithdrawEarnedFees    int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDefineService, &weightMsgDefineService, nil,
		func(_ *rand.Rand) {
			weightMsgDefineService = DefaultWeightMsgDefineService
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBindService, &weightMsgBindService, nil,
		func(_ *rand.Rand) {
			weightMsgBindService = DefaultWeightMsgBindService
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateServiceBinding, &weightMsgUpdateServiceBinding, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateServiceBinding = DefaultWeightMsgUpdateServiceBinding
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSetWithdrawAddress, &weightMsgSetWithdrawAddress, nil,
		func(_ *rand.Rand) {
			weightMsgSetWithdrawAddress = DefaultWeightMsgSetWithdrawAddress
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDisableServiceBinding, &weightMsgDisableServiceBinding, nil,
		func(_ *rand.Rand) {
			weightMsgDisableServiceBinding = DefaultWeightMsgDisableServiceBinding
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgEnableServiceBinding, &weightMsgEnableServiceBinding, nil,
		func(_ *rand.Rand) {
			weightMsgEnableServiceBinding = DefaultWeightMsgEnableServiceBinding
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRefundServiceDeposit, &weightMsgRefundServiceDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgRefundServiceDeposit
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRefundServiceDeposit, &weightMsgRefundServiceDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgRefundServiceDeposit
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCallService, &weightMsgCallService, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgCallService
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRespondService, &weightMsgRespondService, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgRespondService
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgStartRequestContext, &weightMsgStartRequestContext, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgStartRequestContext
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgPauseRequestContext, &weightMsgPauseRequestContext, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgPauseRequestContext
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgKillRequestContext, &weightMsgKillRequestContext, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgKillRequestContext
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateRequestContext, &weightMsgUpdateRequestContext, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = DefaultWeightMsgUpdateRequestContext
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgWithdrawEarnedFees, &weightMsgWithdrawEarnedFees, nil,
		func(_ *rand.Rand) {
			weightMsgRefundServiceDeposit = weightMsgWithdrawEarnedFees
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgDefineService,
			SimulateMsgDefineService(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgBindService,
			SimulateMsgBindService(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateServiceBinding,
			SimulateMsgUpdateServiceBinding(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgSetWithdrawAddress,
			SimulateMsgSetWithdrawAddress(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgDisableServiceBinding,
			SimulateMsgDisableServiceBinding(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgEnableServiceBinding,
			SimulateMsgEnableServiceBinding(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRefundServiceDeposit,
			SimulateMsgRefundServiceDeposit(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgCallService,
			SimulateMsgCallService(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRespondService,
			SimulateMsgRespondService(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgStartRequestContext,
			SimulateMsgStartRequestContext(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgPauseRequestContext,
			SimulateMsgPauseRequestContext(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgKillRequestContext,
			SimulateMsgKillRequestContext(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateRequestContext,
			SimulateMsgKillRequestContext(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateRequestContext,
			SimulateMsgUpdateRequestContext(ak, bk, k),
		),
		//simulation.NewWeightedOperation(
		//	weightMsgWithdrawEarnedFees,
		//	SimulateMsgWithdrawEarnedFees(ak, bk, k),
		//),
	}
}

// SimulateMsgDefineService generates a MsgDefineService with random values.
func SimulateMsgDefineService(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)

		serviceName := simtypes.RandStringOfLength(r, 70)
		serviceDescription := simtypes.RandStringOfLength(r, 280)
		authorDescription := simtypes.RandStringOfLength(r, 280)
		tags := []string{simtypes.RandStringOfLength(r, 20), simtypes.RandStringOfLength(r, 20)}
		schemas := `{"input":{"type":"object"},"output":{"type":"object"}}`

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := types.NewMsgDefineService(serviceName, serviceDescription, tags, simAccount.Address.String(), authorDescription, schemas)

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgBindService generates a MsgBindService with random values.
func SimulateMsgBindService(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)

		def := GenServiceDefinition(r, k, ctx)

		serviceName := def.Name
		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		deposit := GenDeposit(r, spendable)
		pricing := fmt.Sprintf(`{"price":"%d%s"}`, simtypes.RandIntBetween(r, 100, 1000), sdk.DefaultBondDenom)
		qos := uint64(simtypes.RandIntBetween(r, 10, 100))
		options := "{}"

		msg := types.NewMsgBindService(serviceName, simAccount.Address.String(), deposit, pricing, qos, options, def.Author)

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgUpdateServiceBinding generates a MsgUpdateServiceBinding with random values.
func SimulateMsgUpdateServiceBinding(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)

		serviceName := simtypes.RandStringOfLength(r, 20)
		deposit := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 100, 1000)))))
		pricing := fmt.Sprintf(`{"price":"%d%s"}`, simtypes.RandIntBetween(r, 100, 1000), sdk.DefaultBondDenom)
		qos := uint64(simtypes.RandIntBetween(r, 10, 100))
		// XXX
		options := "{}"

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := types.NewMsgUpdateServiceBinding(serviceName, simAccount.Address.String(), deposit, pricing, qos, options, simAccount.Address.String())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgSetWithdrawAddress generates a MsgSetWithdrawAddress with random values.
func SimulateMsgSetWithdrawAddress(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)
		withdrawalAccount, _ := simtypes.RandomAcc(r, accs)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := types.NewMsgSetWithdrawAddress(simAccount.Address.String(), withdrawalAccount.Address.String())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgDisableServiceBinding generates a MsgDisableServiceBinding with random values.
func SimulateMsgDisableServiceBinding(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)
		serviceName := simtypes.RandStringOfLength(r, 20)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := types.NewMsgDisableServiceBinding(serviceName, simAccount.Address.String(), simAccount.Address.String())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgEnableServiceBinding generates a MsgEnableServiceBinding with random values.
func SimulateMsgEnableServiceBinding(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)

		serviceName := simtypes.RandStringOfLength(r, 20)
		deposit := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 100, 1000)))))

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := types.NewMsgEnableServiceBinding(serviceName, simAccount.Address.String(), deposit, simAccount.Address.String())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgRefundServiceDeposit generates a MsgRefundServiceDeposit with random values.
func SimulateMsgRefundServiceDeposit(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		simAccount, _ := simtypes.RandomAcc(r, accs)
		serviceName := simtypes.RandStringOfLength(r, 20)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg := types.NewMsgRefundServiceDeposit(serviceName, simAccount.Address.String(), simAccount.Address.String())

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgCallService generates a MsgCallService with random values.
func SimulateMsgCallService(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAccount.Address)

		definition := GenServiceDefinition(r, k, ctx)
		if definition.Size() == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCallService, "definition not exist"), nil, nil
		}

		providers := GetProviders(definition, k, ctx)
		if len(providers) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCallService, "providers not exist"), nil, nil
		}

		serviceName := definition.Name
		consumer := simAccount.Address.String()
		input := `{"header":{},"body":{}}`
		serviceFeeCap := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 2, 10))))}
		timeout := r.Int63n(k.MaxRequestTimeout(ctx))

		// Temporarily disabled in irishub-v1.0.0
		repeated := false
		repeatedFrequency := uint64(0)
		repeatedTotal := int64(0)

		msg := types.NewMsgCallService(serviceName, providers, consumer, input,
			serviceFeeCap, timeout, repeated, repeatedFrequency, repeatedTotal)

		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		spendable, hasNeg := spendable.SafeSub(serviceFeeCap)
		if hasNeg {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCallService, "Insufficient funds"), nil, nil
		}

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgRespondService generates a MsgRespondService with random values.
func SimulateMsgRespondService(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		requestId := GenRequestContextId(r, k, ctx)
		if len(requestId) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRespondService, "requestId not exist"), nil, nil
		}

		request, found := k.GetRequest(ctx, requestId)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRespondService, "request not found"), nil, nil
		}

		provider, err := sdk.AccAddressFromBech32(request.Provider)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRespondService, "invalid address"), nil, nil
		}

		result := `{"code":200,"message":""}`
		output := `{"header":{},"body":{}}`

		msg := types.NewMsgRespondService(requestId.String(), request.Provider, result, output)

		acc, found := simtypes.FindAccount(accs, provider)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRespondService, "account not found"), nil, nil
		}

		account := ak.GetAccount(ctx, acc.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil

	}
}

// SimulateMsgPauseRequestContext generates a MsgSPauseRequestContext with random values.
func SimulateMsgPauseRequestContext(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// request must be running
		requestId := GenRunningContextId(r, k, ctx)
		if len(requestId) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPauseRequestContext, "requestId not exist"), nil, nil
		}

		request, found := k.GetRequest(ctx, requestId)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPauseRequestContext, "request not found"), nil, nil
		}

		consumer, err := sdk.AccAddressFromBech32(request.Consumer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPauseRequestContext, "invalid address"), nil, nil
		}

		msg := types.NewMsgPauseRequestContext(requestId.String(), consumer.String())

		acc, found := simtypes.FindAccount(accs, consumer)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgPauseRequestContext, "account not found"), nil, nil
		}

		account := ak.GetAccount(ctx, acc.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgStartRequestContext generates a MsgStartRequestContext with random values.
func SimulateMsgStartRequestContext(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// request must be paused
		requestId := GenPausedRequestContextId(r, k, ctx)
		if len(requestId) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequestContext, "requestId not exist"), nil, nil
		}

		request, found := k.GetRequest(ctx, requestId)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequestContext, "request not found"), nil, nil
		}

		consumer, err := sdk.AccAddressFromBech32(request.Consumer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequestContext, "invalid address"), nil, nil
		}

		msg := types.NewMsgPauseRequestContext(requestId.String(), consumer.String())

		acc, found := simtypes.FindAccount(accs, consumer)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequestContext, "account not found"), nil, nil
		}

		account := ak.GetAccount(ctx, acc.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgKillRequestContext generates a MsgKillRequestContext with random values.
func SimulateMsgKillRequestContext(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		requestId := GenRequestContextId(r, k, ctx)
		if len(requestId) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgKillRequestContext, "requestId not exist"), nil, nil
		}

		request, found := k.GetRequest(ctx, requestId)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgKillRequestContext, "request not found"), nil, nil
		}

		consumer, err := sdk.AccAddressFromBech32(request.Consumer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgKillRequestContext, "invalid address"), nil, nil
		}

		msg := types.NewMsgPauseRequestContext(requestId.String(), consumer.String())

		acc, found := simtypes.FindAccount(accs, consumer)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgKillRequestContext, "account not found"), nil, nil
		}

		account := ak.GetAccount(ctx, acc.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgUpdateRequestContext generates a MsgUpdateRequestContext with random values.
func SimulateMsgUpdateRequestContext(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		requestId := GenRequestContextId(r, k, ctx)
		if len(requestId) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequestContext, "requestId not exist"), nil, nil
		}

		request, found := k.GetRequest(ctx, requestId)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequestContext, "request not found"), nil, nil
		}

		consumer, err := sdk.AccAddressFromBech32(request.Consumer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequestContext, "invalid address"), nil, nil
		}

		definition, found := k.GetServiceDefinition(ctx, request.ServiceName)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequestContext, "definition not found"), nil, nil
		}
		providers := GetProviders(definition, k, ctx)

		serviceFeeCap := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 2, 10))))}
		timeout := r.Int63n(k.MaxRequestTimeout(ctx))
		repeatedFrequency := uint64(0)
		repeatedTotal := int64(0)

		msg := types.NewMsgUpdateRequestContext(requestId.String(), providers, serviceFeeCap, timeout, repeatedFrequency, repeatedTotal, consumer.String())

		acc, found := simtypes.FindAccount(accs, consumer)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgKillRequestContext, "account not found"), nil, nil
		}

		account := ak.GetAccount(ctx, acc.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		spendable, hasNeg := spendable.SafeSub(serviceFeeCap)

		if hasNeg {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgCallService, "Insufficient funds"), nil, nil
		}

		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
		}

		txGen := cosmossimappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate mock tx"), nil, err
		}

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgWithdrawEarnedFees generates a MsgWithdrawEarnedFees with random values.
//func SimulateMsgWithdrawEarnedFees(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
//	return func(
//		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
//	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
//
//
//
//
//
//
//	}
//}

// GenServiceDefinition randomized serviceDefinition
func GenServiceDefinition(r *rand.Rand, k keeper.Keeper, ctx sdk.Context) types.ServiceDefinition {
	var definitions []types.ServiceDefinition
	k.IterateServiceDefinitions(
		ctx,
		func(definition types.ServiceDefinition) bool {
			definitions = append(definitions, definition)
			return false
		},
	)
	if len(definitions) > 0 {
		return definitions[r.Intn(len(definitions))]
	}
	return types.ServiceDefinition{}
}

// GenServiceBinding randomized serviceBinding
func GenServiceBinding(r *rand.Rand, k keeper.Keeper, ctx sdk.Context) types.ServiceBinding {
	var bindings []types.ServiceBinding
	k.IterateServiceBindings(
		ctx,
		func(binding types.ServiceBinding) bool {
			bindings = append(bindings, binding)
			return false
		},
	)
	if len(bindings) > 0 {
		return bindings[r.Intn(len(bindings))]
	}
	return types.ServiceBinding{}
}

// GenRequestContextId randomized requestContext
func GenRequestContextId(r *rand.Rand, k keeper.Keeper, ctx sdk.Context) tmbytes.HexBytes {
	var requestIds []tmbytes.HexBytes
	k.IterateRequestContexts(
		ctx,
		func(requestContextID tmbytes.HexBytes, requestContext types.RequestContext) bool {
			requestIds = append(requestIds, requestContextID)
			return false
		},
	)
	if len(requestIds) > 0 {
		return requestIds[r.Intn(len(requestIds))]
	}
	return tmbytes.HexBytes{}
}

// GenRunningContextId randomized runningContextId
func GenRunningContextId(r *rand.Rand, k keeper.Keeper, ctx sdk.Context) tmbytes.HexBytes {
	var requestIds []tmbytes.HexBytes
	k.IterateRequestContexts(
		ctx,
		func(requestContextID tmbytes.HexBytes, requestContext types.RequestContext) bool {
			if requestContext.State == types.RUNNING {
				requestIds = append(requestIds, requestContextID)
				return false
			}
			return false
		},
	)
	if len(requestIds) > 0 {
		return requestIds[r.Intn(len(requestIds))]
	}
	return tmbytes.HexBytes{}
}

// GenPausedRequestContextId randomized pausedRequestContextId
func GenPausedRequestContextId(r *rand.Rand, k keeper.Keeper, ctx sdk.Context) tmbytes.HexBytes {
	var requestIds []tmbytes.HexBytes
	k.IterateRequestContexts(
		ctx,
		func(requestContextID tmbytes.HexBytes, requestContext types.RequestContext) bool {
			if requestContext.State == types.PAUSED {
				requestIds = append(requestIds, requestContextID)
				return false
			}
			return false
		},
	)
	if len(requestIds) > 0 {
		return requestIds[r.Intn(len(requestIds))]
	}
	return tmbytes.HexBytes{}
}

func GetProviders(definition types.ServiceDefinition, k keeper.Keeper, ctx sdk.Context) (providers []string) {
	if definition.Size() == 0 {
		return
	}

	owner, err := sdk.AccAddressFromBech32(definition.Author)
	if err != nil {
		return
	}

	bindings := k.GetOwnerServiceBindings(ctx, owner, definition.Name)
	if len(bindings) > 0 {
		for _, binding := range bindings {
			providers = append(providers, binding.Provider)
		}
	}
	return
}

//GenDeposit randomized deposit
func GenDeposit(r *rand.Rand, spendable sdk.Coins) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, simtypes.RandomAmount(r, spendable.AmountOf(sdk.DefaultBondDenom))))
}

//GenBoolValue randomized bool value
func GenBoolValue(r *rand.Rand) bool {
	return r.Int()%2 == 0
}
