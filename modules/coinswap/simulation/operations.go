package simulation

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/irisnet/irismod/modules/coinswap/keeper"
	"github.com/irisnet/irismod/modules/coinswap/types"
	tokentypes "github.com/irisnet/irismod/modules/token/types"
	"math/rand"
	"time"
)

// Simulation operation weights constants
const (
	OpWeightMsgSwapOrder       = "op_weight_msg_swap_order"
	OpWeightMsgAddLiquidity    = "op_weight_msg_add_liquidity"
	OpWeightMsgRemoveLiquidity = "op_weight_msg_remove_liquidity"
)

func WeightedOperations(
	appParams simtypes.AppParams,
	cdc codec.JSONMarshaler,
	k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simulation.WeightedOperations {

	var (
		weightSwap   int
		weightAdd    int
		weightRemove int
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgSwapOrder, &weightSwap, nil,
		func(_ *rand.Rand) {
			weightSwap = 50
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgAddLiquidity, &weightAdd, nil,
		func(_ *rand.Rand) {
			weightAdd = 100
		},
	)

	appParams.GetOrGenerate(
		cdc, OpWeightMsgRemoveLiquidity, &weightRemove, nil,
		func(_ *rand.Rand) {
			weightRemove = 30
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightAdd,
			SimulateMsgAddLiquidity(k, ak, bk),
		),
		//simulation.NewWeightedOperation(
		//	weightSwap,
		//	SimulateMsgSwapOrder(k, ak, bk),
		//),

		simulation.NewWeightedOperation(
			weightRemove,
			SimulateMsgRemoveLiquidity(k, ak, bk),
		),
	}
}

//SimulateMsgSwapOrder  simulates  the swap of order
func SimulateMsgSwapOrder(k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (
		opMsg simtypes.OperationMsg, fOps []simtypes.FutureOperation, err error,
	) {
		var (
			inputCoin, outputCoin sdk.Coin
			isBuyOrder            bool
		)

		simAccount, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAccount.Address)

		standardDenom := k.GetStandardDenom(ctx)
		inputToken := RandomTokenInPool(r, ctx, accs, k)
		outputToken := RandomTokenInPool(r, ctx, accs, k)

		inputCoinTemp := RandomCoin(r, inputToken.MinUnit, sdk.NewIntFromUint64(inputToken.InitialSupply))
		outputCoinTemp := RandomCoin(r, outputToken.MinUnit, sdk.NewIntFromUint64(outputToken.InitialSupply))

		isDoubleSwap := (inputToken.MinUnit != standardDenom) && (outputToken.MinUnit != standardDenom)

		i := rand.Intn(4)
		switch i {
		case 0:
			if isDoubleSwap {
				inputCoin, outputCoin, isBuyOrder = doubleSwapBill(inputCoinTemp, outputCoinTemp, ctx, k)
			}
		case 1:
			if !isDoubleSwap {
				inputCoin, outputCoin, isBuyOrder = singleSwapBill(inputCoinTemp, outputCoinTemp, ctx, k)
			}
		case 2:
			if isDoubleSwap {
				inputCoin, outputCoin, isBuyOrder = doubleSwapSellOrder(inputCoinTemp, outputCoinTemp, ctx, k)
			}
		case 3:
			if !isDoubleSwap {
				inputCoin, outputCoin, isBuyOrder = singleSwapSellOrder(inputCoinTemp, outputCoinTemp, ctx, k)
			}
		}

		deadline := int64(time.Now().Add(time.Second * time.Duration(r.Intn(10))).Second())
		msg := types.NewMsgSwapOrder(
			types.Input{
				Address: simAccount.Address.String(),
				Coin:    inputCoin,
			},
			types.Output{
				Address: simAccount.Address.String(),
				Coin:    outputCoin,
			},
			deadline,
			isBuyOrder,
		)

		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

//SimulateMsgAddLiquidity  simulates  the addition of liquidity
func SimulateMsgAddLiquidity(k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (
		opMsg simtypes.OperationMsg, fOps []simtypes.FutureOperation, err error,
	) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAccount.Address)

		var maxToken sdk.Coin
		var minLiquidity sdk.Int

		irisAmt := simtypes.RandomAmount(r, sdk.NewIntFromUint64(tokentypes.GetNativeToken().InitialSupply))
		standardDenom := k.GetStandardDenom(ctx)
		token := RandomToken(r)
		fmt.Println("^^", token.Symbol)
		uniDenom := types.GetUniDenomFromDenom(token.GetMinUnit())
		liquidity := bk.GetSupply(ctx).GetTotal().AmountOf(uniDenom)
		reservePool, err := k.GetReservePool(ctx, uniDenom)

		if err != nil {
			maxToken = RandomCoin(r, token.MinUnit, sdk.NewIntFromUint64(token.InitialSupply))
			minLiquidity = simtypes.RandomAmount(r, irisAmt)
		} else {
			maxToken = RandomCoin(r, token.MinUnit, sdk.NewIntFromUint64(token.InitialSupply))
			standardReserveAmt := reservePool.AmountOf(standardDenom)
			minLiquidity = liquidity.Mul(irisAmt).Quo(standardReserveAmt)
		}

		deadline := randDeadline(r)
		msg := types.NewMsgAddLiquidity(
			maxToken,
			irisAmt,
			minLiquidity,
			deadline,
			account.GetAddress().String(),
		)

		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}

}

//SimulateMsgRemoveLiquidity  simulates  the removal of liquidity
func SimulateMsgRemoveLiquidity(k keeper.Keeper, ak types.AccountKeeper, bk types.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (
		opMsg simtypes.OperationMsg, fOps []simtypes.FutureOperation, err error,
	) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAccount.Address)
		standardDenom := k.GetStandardDenom(ctx)

		var minToken sdk.Int
		var minStandardAmt sdk.Int

		token := RandomTokenInPool(r, ctx, accs, k)

		uniDenom := types.GetUniDenomFromDenom(token.MinUnit)
		minTokenDenom, err := types.GetCoinDenomFromUniDenom(uniDenom)
		reservePool, err := k.GetReservePool(ctx, uniDenom)

		liquidityReserve := bk.GetSupply(ctx).GetTotal().AmountOf(uniDenom)
		standardReserveAmt := reservePool.AmountOf(standardDenom)
		tokenReserveAmt := reservePool.AmountOf(minTokenDenom)

		withdrawLiquidity := RandomCoin(r, token.MinUnit, sdk.NewIntFromUint64(token.InitialSupply))

		deadline := randDeadline(r)
		if err != nil {
			minToken = simtypes.RandomAmount(r, tokenReserveAmt)
			minStandardAmt = simtypes.RandomAmount(r, standardReserveAmt)
		} else {
			minToken = withdrawLiquidity.Amount.Mul(tokenReserveAmt).Quo(liquidityReserve)
			minStandardAmt = withdrawLiquidity.Amount.Mul(standardReserveAmt).Quo(liquidityReserve)
		}

		msg := types.NewMsgRemoveLiquidity(
			minToken,
			withdrawLiquidity,
			minStandardAmt,
			deadline,
			account.GetAddress().String(),
		)
		fmt.Println(msg.Deadline)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		fees, err := simtypes.RandomFees(r, ctx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, err
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

		if _, _, err := app.Deliver(txGen.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil

	}
}

func RandomCoin(r *rand.Rand, denom string, amount sdk.Int) sdk.Coin {
	return sdk.NewInt64Coin(denom, r.Int63n(amount.Int64()))
}

func RandomToken(r *rand.Rand) tokentypes.Token {

	index := r.Intn(len(tokens))

	fmt.Println("*******",index)
	token := tokens[index]
	//fmt.Println("&&", token.Symbol)
	//fmt.Println("&&len", len(tokens))
	for i, t := range tokens {
		fmt.Println("i@@@@@:",i)
		fmt.Println("t@@@@@:",t.Symbol)
	}
	fmt.Println("^^",token.Symbol)
	return token
}

// RandomTokenInPool generate token in pool
func RandomTokenInPool(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper) tokentypes.Token {
	for {
		token := randToken(r, accs)
		// token must be in reseve pool
		if reservePool, _ := k.GetReservePool(ctx, types.GetUniDenomFromDenom(token.GetMinUnit())); !reservePool.IsZero() {
			return token
		}
	}
}

// generate coin in reserve pool
func randomCoinInPool(r *rand.Rand, ctx sdk.Context, coins sdk.Coins, k keeper.Keeper) sdk.Coin {
	for {
		denom := coins[r.Intn(len(coins))].GetDenom()
		if reservePool, _ := k.GetReservePool(ctx, types.GetUniDenomFromDenom(denom)); !reservePool.IsZero() {
			return sdk.Coin{Denom: denom, Amount: reservePool.AmountOf(denom)}
		}
	}
}

func randDeadline(r *rand.Rand) int64 {
	var delta = time.Duration(simtypes.RandIntBetween(r,10,100)) * time.Second
	return time.Now().Add(delta).UnixNano()
}

//

//Double swap bill
func doubleSwapBill(inputCoin, outputCoin sdk.Coin, ctx sdk.Context, k keeper.Keeper) (sdk.Coin, sdk.Coin, bool) {
	standardDenom := k.GetStandardDenom(ctx)
	param := k.GetParams(ctx)

	//generate sold standard Coin
	uniDenom, _ := k.GetUniDenomFromDenoms(ctx, outputCoin.Denom, standardDenom)
	reservePool, _ := k.GetReservePool(ctx, uniDenom)
	outputReserve := reservePool.AmountOf(outputCoin.Denom)
	inputReserve := reservePool.AmountOf(standardDenom)
	soldStandardAmount := keeper.GetOutputPrice(outputCoin.Amount, inputReserve, outputReserve, param.Fee)
	soldStandardCoin := sdk.NewCoin(standardDenom, soldStandardAmount)

	//generate input coin
	uniDenom2, _ := k.GetUniDenomFromDenoms(ctx, soldStandardCoin.Denom, inputCoin.Denom)
	reservePool2, _ := k.GetReservePool(ctx, uniDenom2)
	outputReserve2 := reservePool2.AmountOf(soldStandardCoin.Denom)
	inputReserve2 := reservePool2.AmountOf(inputCoin.Denom)
	soldTokenAmt := keeper.GetOutputPrice(soldStandardCoin.Amount, inputReserve2, outputReserve2, param.Fee)
	inputCoin = sdk.NewCoin(inputCoin.Denom, soldTokenAmt)

	return inputCoin, outputCoin, true
}

//A single swap bill
func singleSwapBill(inputCoin, outputCoin sdk.Coin, ctx sdk.Context, k keeper.Keeper) (sdk.Coin, sdk.Coin, bool) {
	param := k.GetParams(ctx)

	uniDenom, _ := k.GetUniDenomFromDenoms(ctx, outputCoin.Denom, inputCoin.Denom)
	reservePool, _ := k.GetReservePool(ctx, uniDenom)
	outputReserve := reservePool.AmountOf(outputCoin.Denom)
	inputReserve := reservePool.AmountOf(inputCoin.Denom)
	soldTokenAmt := keeper.GetOutputPrice(outputCoin.Amount, inputReserve, outputReserve, param.Fee)
	inputCoin = sdk.NewCoin(inputCoin.Denom, soldTokenAmt)

	return inputCoin, outputCoin, true
}

//Double swap sell orders
func doubleSwapSellOrder(inputCoin, outputCoin sdk.Coin, ctx sdk.Context, k keeper.Keeper) (sdk.Coin, sdk.Coin, bool) {
	standardDenom := k.GetStandardDenom(ctx)

	param := k.GetParams(ctx)

	uniDenom, _ := k.GetUniDenomFromDenoms(ctx, inputCoin.Denom, standardDenom)
	reservePool, _ := k.GetReservePool(ctx, uniDenom)
	inputReserve := reservePool.AmountOf(inputCoin.Denom)
	outputReserve := reservePool.AmountOf(standardDenom)
	standardAmount := keeper.GetInputPrice(inputCoin.Amount, inputReserve, outputReserve, param.Fee)
	standardCoin := sdk.NewCoin(standardDenom, standardAmount)

	uniDenom2, _ := k.GetUniDenomFromDenoms(ctx, standardCoin.Denom, outputCoin.Denom)
	reservePool2, _ := k.GetReservePool(ctx, uniDenom2)
	inputReserve2 := reservePool2.AmountOf(standardCoin.Denom)
	outputReserve2 := reservePool2.AmountOf(outputCoin.Denom)
	boughtTokenAmt := keeper.GetInputPrice(standardCoin.Amount, inputReserve2, outputReserve2, param.Fee)
	outputCoin = sdk.NewCoin(outputCoin.Denom, boughtTokenAmt)

	return inputCoin, outputCoin, false
}

//A single swap sell order
func singleSwapSellOrder(inputCoin, outputCoin sdk.Coin, ctx sdk.Context, k keeper.Keeper) (sdk.Coin, sdk.Coin, bool) {
	param := k.GetParams(ctx)

	uniDenom, _ := k.GetUniDenomFromDenoms(ctx, inputCoin.Denom, outputCoin.Denom)
	reservePool, _ := k.GetReservePool(ctx, uniDenom)
	inputReserve := reservePool.AmountOf(inputCoin.Denom)
	outputReserve := reservePool.AmountOf(outputCoin.Denom)
	boughtTokenAmt := keeper.GetInputPrice(inputCoin.Amount, inputReserve, outputReserve, param.Fee)

	outputCoin = sdk.NewCoin(outputCoin.Denom, boughtTokenAmt)
	return inputCoin, outputCoin, false
}
