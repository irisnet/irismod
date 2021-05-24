package simulation

import (
	"math/rand"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/irisnet/irismod/modules/coinswap/keeper"
	"github.com/irisnet/irismod/modules/coinswap/types"
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

		//simulation.NewWeightedOperation(
		//	weightRemove,
		//	SimulateMsgRemoveLiquidity(k, ak, bk),
		//),
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

		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		//if outputCoin.Size() == 0{
		//	return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, "outputCoin should not be empty"), nil, err
		//}

		isBuyOrder = randBoolean(r)

		//
		coins := bk.GetSupply(ctx).GetTotal()
		randomInputCoin := randomCoinInpool(r, ctx, spendable, k)
		randomOutputCoin := randomCoinInpool(r, ctx, coins, k)
		isDoubleSwap := (randomOutputCoin.Denom != standardDenom) && (randomInputCoin.Denom != standardDenom)
		if isBuyOrder && isDoubleSwap {
			inputCoin, outputCoin, isBuyOrder = doubleSwapBill(randomInputCoin, randomOutputCoin, ctx, k)
		} else if isBuyOrder && !isDoubleSwap {
			inputCoin, outputCoin, isBuyOrder = singleSwapBill(randomInputCoin, randomOutputCoin, ctx, k)
		} else if !isBuyOrder && isDoubleSwap {
			inputCoin, outputCoin, isBuyOrder = doubleSwapSellOrder(randomInputCoin, randomOutputCoin, ctx, k)
		} else if !isBuyOrder && !isDoubleSwap {
			inputCoin, outputCoin, isBuyOrder = singleSwapSellOrder(randomInputCoin, randomOutputCoin, ctx, k)
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

		var (
			maxToken     sdk.Coin
			minLiquidity sdk.Int
		)

		standardDenom := k.GetStandardDenom(ctx)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		exactStandardAmt := simtypes.RandomAmount(r, spendable.AmountOf(standardDenom))

		if !exactStandardAmt.IsPositive() {
			return
		}

		maxToken = RandomSpendableToken(r, spendable)
		if maxToken.Denom == standardDenom {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, "tokenDenom should not be standardDenom"), nil, err
		}

		if strings.HasPrefix(maxToken.Denom, types.FormatUniABSPrefix) {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, "tokenDenom should not be liquidity token"), nil, err
		}

		if !maxToken.Amount.IsPositive() {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, "maxToken must is positive"), nil, err
		}

		uniDenom := types.GetUniDenomFromDenom(maxToken.Denom)

		reservePool, err := k.GetReservePool(ctx, uniDenom)

		if err != nil {
			minLiquidity = exactStandardAmt
		} else {
			standardReserveAmt := reservePool.AmountOf(standardDenom)
			liquidity := reservePool.AmountOf(uniDenom)
			minLiquidity = liquidity.Mul(exactStandardAmt).Quo(standardReserveAmt)

			if !maxToken.Amount.Sub(reservePool.AmountOf(maxToken.GetDenom()).Mul(exactStandardAmt).Quo(standardReserveAmt)).IsPositive() {
				return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, "insufficient funds"), nil, err
			}
		}

		deadline := randDeadline(r)
		msg := types.NewMsgAddLiquidity(
			maxToken,
			exactStandardAmt,
			minLiquidity,
			deadline,
			account.GetAddress().String(),
		)

		var fees sdk.Coins
		coinsTemp, hasNeg := spendable.SafeSub(sdk.NewCoins(sdk.NewCoin(standardDenom, exactStandardAmt), maxToken))
		if !hasNeg {
			fees, err = simtypes.RandomFees(r, ctx, coinsTemp)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to generate fees"), nil, nil
			}
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

		var (
			minToken          sdk.Int
			minStandardAmt    sdk.Int
			withdrawLiquidity sdk.Coin
		)

		spendable := bk.SpendableCoins(ctx, account.GetAddress())
		//token := RandomSpendableToken(r, spendable)
		token := RandomSpendableToken(r, spendable)

		////should not be standardDenom
		if token.Denom == standardDenom {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, "tokenDenom  should not be standardDenom"), nil, err
		}

		//token.Denom must be "swapDenom"
		tokenDenom, err := types.GetCoinDenomFromUniDenom(token.Denom)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, err.Error()), nil, err
		}

		reservePool, err := k.GetReservePool(ctx, token.Denom)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveLiquidity, err.Error()), nil, err
		} else {
			standardReserveAmt := reservePool.AmountOf(standardDenom)
			tokenReserveAmt := reservePool.AmountOf(tokenDenom)

			withdrawLiquidity = sdk.NewCoin(token.GetDenom(), simtypes.RandomAmount(r, sdk.MinInt(token.Amount, tokenReserveAmt)))
			liquidityReserve := bk.GetSupply(ctx).GetTotal().AmountOf(token.Denom)
			minToken = withdrawLiquidity.Amount.Mul(tokenReserveAmt).Quo(liquidityReserve)
			minStandardAmt = withdrawLiquidity.Amount.Mul(standardReserveAmt).Quo(liquidityReserve)
		}

		deadline := randDeadline(r)
		msg := types.NewMsgRemoveLiquidity(
			minToken,
			withdrawLiquidity,
			minStandardAmt,
			deadline,
			account.GetAddress().String(),
		)
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
	return sdk.NewCoin(denom, amount)
	//return sdk.NewInt64Coin(denom, r.Int63n(amount.Int64()))
}

func RandomSpendableToken(r *rand.Rand, spendableCoin sdk.Coins) sdk.Coin {
	token := spendableCoin[r.Intn(len(spendableCoin))]
	return sdk.NewCoin(token.Denom, simtypes.RandomAmount(r, token.Amount))
}

func RandomSpendableToken2(r *rand.Rand, spendableCoin sdk.Coins, ctx sdk.Context, k keeper.Keeper) sdk.Coin {
	//standardDenom := k.GetStandardDenom(ctx)
	for {
		token := spendableCoin[r.Intn(len(spendableCoin))]
		_, err := types.GetCoinDenomFromUniDenom(token.Denom)
		if err != nil {
			continue
		}
		return sdk.NewCoin(token.Denom, simtypes.RandomAmount(r, token.Amount))
		//if token.Denom != standardDenom {
		//	return sdk.NewCoin(token.Denom, simtypes.RandomAmount(r, token.Amount))
		//}
	}
}

func RandomSpendableTokenInPool(r *rand.Rand, spendableCoin sdk.Coins, ctx sdk.Context, k keeper.Keeper) sdk.Coin {
	standardDenom := k.GetStandardDenom(ctx)
	token := spendableCoin[r.Intn(len(spendableCoin))]
LOOP:
	if token.Denom != standardDenom {
		// token must be in pool
		if _, err := k.GetReservePool(ctx, types.GetUniDenomFromDenom(token.Denom)); err != nil {
			goto LOOP
		} else {
			return RandomCoin(r, token.Denom, simtypes.RandomAmount(r, token.Amount))
		}
	}
	goto LOOP
}

// randomBoughtCoinInpool
func randomCoinInpool(r *rand.Rand, ctx sdk.Context, coins sdk.Coins, k keeper.Keeper) sdk.Coin {
	for _, coin := range coins {
		uniDenom := types.GetUniDenomFromDenom(coin.GetDenom())
		if _, err := k.GetReservePool(ctx, uniDenom); err != nil {
			continue
		}
		return sdk.NewCoin(coin.Denom, simtypes.RandomAmount(r, coin.Amount))
	}
	return sdk.Coin{}
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
	var delta = time.Duration(simtypes.RandIntBetween(r, 10, 100)) * time.Second
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

func randBoolean(r *rand.Rand) bool {
	return r.Int()%2 == 0
}