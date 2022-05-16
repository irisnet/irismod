package keeper

import (
	"fmt"
	"math/big"
	"strconv"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/irisnet/irismod/modules/coinswap/types"
)

// Keeper of the coinswap store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         storetypes.StoreKey
	bk               types.BankKeeper
	ak               types.AccountKeeper
	paramSpace       paramstypes.Subspace
	feeCollectorName string
	blockedAddrs     map[string]bool
}

// NewKeeper returns a coinswap keeper. It handles:
// - creating new ModuleAccounts for each trading pair
// - burning and minting liquidity coins
// - sending to and from ModuleAccounts
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
	bk types.BankKeeper,
	ak types.AccountKeeper,
	blockedAddrs map[string]bool,
	feeCollectorName string,
) Keeper {
	// ensure coinswap module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:         key,
		bk:               bk,
		ak:               ak,
		cdc:              cdc,
		paramSpace:       paramSpace,
		blockedAddrs:     blockedAddrs,
		feeCollectorName: feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

// Swap execute swap order in specified pool
func (k Keeper) Swap(ctx sdk.Context, msg *types.MsgSwapOrder) error {
	var amount sdk.Int
	var err error

	standardDenom := k.GetStandardDenom(ctx)
	isDoubleSwap := (msg.Input.Coin.Denom != standardDenom) && (msg.Output.Coin.Denom != standardDenom)

	if msg.IsBuyOrder && isDoubleSwap {
		amount, err = k.doubleTradeInputForExactOutput(ctx, msg.Input, msg.Output)
	} else if msg.IsBuyOrder && !isDoubleSwap {
		amount, err = k.TradeInputForExactOutput(ctx, msg.Input, msg.Output)
	} else if !msg.IsBuyOrder && isDoubleSwap {
		amount, err = k.doubleTradeExactInputForOutput(ctx, msg.Input, msg.Output)
	} else if !msg.IsBuyOrder && !isDoubleSwap {
		amount, err = k.TradeExactInputForOutput(ctx, msg.Input, msg.Output)
	}
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSwap,
			sdk.NewAttribute(types.AttributeValueAmount, amount.String()),
			sdk.NewAttribute(types.AttributeValueSender, msg.Input.Address),
			sdk.NewAttribute(types.AttributeValueRecipient, msg.Output.Address),
			sdk.NewAttribute(types.AttributeValueIsBuyOrder, strconv.FormatBool(msg.IsBuyOrder)),
			sdk.NewAttribute(types.AttributeValueTokenPair, types.GetTokenPairByDenom(msg.Input.Coin.Denom, msg.Output.Coin.Denom)),
		),
	)

	return nil
}

// AddLiquidity adds liquidity to the specified pool
func (k Keeper) AddLiquidity(ctx sdk.Context, imsg interface{}) (sdk.Coin, error) {
	var msg *types.MsgAddLiquidity

	switch i := imsg.(type) {
	case *types.MsgAddUnilateralLiquidity:
		return k.addUnilateralLiquidity(ctx, i)
	case *types.MsgAddLiquidity:
		msg = i
	}

	standardDenom := k.GetStandardDenom(ctx)
	if standardDenom == msg.MaxToken.Denom {
		return sdk.Coin{}, sdkerrors.Wrapf(types.ErrInvalidDenom,
			"MaxToken: %s should not be StandardDenom", msg.MaxToken.String())
	}
	var mintLiquidityAmt sdk.Int
	var depositToken sdk.Coin
	var standardCoin = sdk.NewCoin(standardDenom, msg.ExactStandardAmt)

	poolId := types.GetPoolId(msg.MaxToken.Denom)
	pool, exists := k.GetPool(ctx, poolId)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdk.Coin{}, err
	}

	// calculate amount of UNI to be minted for sender
	// and coin amount to be deposited
	if !exists {
		// deduct the user's fee for creating a Liquidity pool
		if err := k.DeductPoolCreationFee(ctx, sender); err != nil {
			return sdk.Coin{}, err
		}

		mintLiquidityAmt = msg.ExactStandardAmt
		if mintLiquidityAmt.LT(msg.MinLiquidity) {
			return sdk.Coin{}, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("liquidity amount not met, user expected: no less than %s, actual: %s", msg.MinLiquidity.String(), mintLiquidityAmt.String()))
		}
		depositToken = sdk.NewCoin(msg.MaxToken.Denom, msg.MaxToken.Amount)
		pool = k.CreatePool(ctx, msg.MaxToken.Denom)
	} else {
		balances, err := k.GetPoolBalances(ctx, pool.EscrowAddress)
		if err != nil {
			return sdk.Coin{}, err
		}

		standardReserveAmt := balances.AmountOf(standardDenom)
		tokenReserveAmt := balances.AmountOf(msg.MaxToken.Denom)
		liquidity := k.bk.GetSupply(ctx, pool.LptDenom).Amount

		mintLiquidityAmt = (liquidity.Mul(msg.ExactStandardAmt)).Quo(standardReserveAmt)
		if mintLiquidityAmt.LT(msg.MinLiquidity) {
			return sdk.Coin{}, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("liquidity amount not met, user expected: no less than %s, actual: %s", msg.MinLiquidity.String(), mintLiquidityAmt.String()))
		}
		depositAmt := (tokenReserveAmt.Mul(msg.ExactStandardAmt)).Quo(standardReserveAmt).AddRaw(1)
		depositToken = sdk.NewCoin(msg.MaxToken.Denom, depositAmt)

		if depositAmt.GT(msg.MaxToken.Amount) {
			return sdk.Coin{}, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("token amount not met, user expected: no more than %s, actual: %s", msg.MaxToken.String(), depositToken.String()))
		}
	}

	reservePoolAddress, err := sdk.AccAddressFromBech32(pool.EscrowAddress)
	if err != nil {
		return sdk.Coin{}, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddLiquidity,
			sdk.NewAttribute(types.AttributeValueSender, msg.Sender),
			sdk.NewAttribute(types.AttributeValueTokenPair, types.GetTokenPairByDenom(msg.MaxToken.Denom, standardDenom)),
		),
	)
	return k.addLiquidity(ctx, sender, reservePoolAddress, standardCoin, depositToken, pool.LptDenom, mintLiquidityAmt)
}

func (k Keeper) addLiquidity(ctx sdk.Context,
	sender sdk.AccAddress,
	reservePoolAddress sdk.AccAddress,
	standardCoin, token sdk.Coin,
	lptDenom string,
	mintLiquidityAmt sdk.Int,
) (sdk.Coin, error) {
	depositedTokens := sdk.NewCoins(standardCoin, token)
	// transfer deposited token into coinswaps Account
	if err := k.bk.SendCoins(ctx, sender, reservePoolAddress, depositedTokens); err != nil {
		return sdk.Coin{}, err
	}

	mintToken := sdk.NewCoin(lptDenom, mintLiquidityAmt)
	mintTokens := sdk.NewCoins(mintToken)
	if err := k.bk.MintCoins(ctx, types.ModuleName, mintTokens); err != nil {
		return sdk.Coin{}, err
	}
	if err := k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, mintTokens); err != nil {
		return sdk.Coin{}, err
	}

	return mintToken, nil
}

func (k Keeper) addUnilateralLiquidity(ctx sdk.Context, msg *types.MsgAddUnilateralLiquidity) (sdk.Coin, error) {

	pool, exist := k.GetPoolBySequenceId(ctx, msg.PoolId)
	if !exist {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrReservePoolNotExists, fmt.Sprintf("liquidity pool id: %d ", msg.PoolId))
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdk.Coin{}, err
	}

	poolAddr, err := sdk.AccAddressFromBech32(pool.EscrowAddress)
	if err != nil {
		return sdk.Coin{}, err
	}

	balances, err := k.GetPoolBalances(ctx, pool.EscrowAddress)
	if err != nil {
		return sdk.Coin{}, err
	}

	// square = ( token_balance + ( 1- fee_unilateral ) * exact_token ) / token_balance * lpt_balance^2
	// 1 - fee_unilateral = numerator / denominator
	tokenBalanceAmt := balances.AmountOf(msg.ExactToken.Denom)
	lptBalanceAmt := k.bk.GetSupply(ctx, pool.LptDenom).Amount
	exactTokenAmt := msg.ExactToken.Amount

	deltaFeeUnilateral := sdk.OneDec().Sub(k.GetParams(ctx).UnilateralLiquidityFee)
	numerator := sdk.NewIntFromBigInt(deltaFeeUnilateral.BigInt())
	denominator := sdk.NewIntWithDecimal(1, sdk.Precision)

	square := denominator.Mul(tokenBalanceAmt).Add(numerator.Mul(exactTokenAmt)).Mul(lptBalanceAmt).Mul(lptBalanceAmt).Quo(denominator.Mul(tokenBalanceAmt))

	// lpt = square^0.5 - lpt_balance
	var squareBigInt = &big.Int{}
	squareBigInt.Sqrt(square.BigInt())
	mintLptAmt := sdk.NewIntFromBigInt(squareBigInt).Sub(lptBalanceAmt)

	if mintLptAmt.LT(msg.MinLiquidity) {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("liquidity amount not met, user expected: no less than %s, actual: %s", msg.MinLiquidity.String(), mintLptAmt.String()))
	}

	// event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddUnilateralLiquidity,
			sdk.NewAttribute(types.AttributeValueSender, msg.Sender),
			sdk.NewAttribute(types.AttributeValueTokenUnilateral, msg.ExactToken.Denom),
		),
	)

	// add liquidity
	exactCoins := sdk.NewCoins(msg.ExactToken)
	if err := k.bk.SendCoins(ctx, sender, poolAddr, exactCoins); err != nil {
		return sdk.Coin{}, err
	}

	// mint and send lpt
	mintLpt := sdk.NewCoin(pool.LptDenom, mintLptAmt)
	mintLpts := sdk.NewCoins(mintLpt)
	if err := k.bk.MintCoins(ctx, types.ModuleName, mintLpts); err != nil {
		return sdk.Coin{}, err
	}
	if err := k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, mintLpts); err != nil {
		return sdk.Coin{}, err
	}

	return mintLpt, nil
}

// RemoveLiquidity removes liquidity from the specified pool
func (k Keeper) RemoveLiquidity(ctx sdk.Context, imsg interface{}) (sdk.Coins, error) {
	var msg *types.MsgRemoveLiquidity

	switch i := imsg.(type) {
	case *types.MsgRemoveUnilateralLiquidity:
		return k.removeUnilateralLiquidity(ctx, i)
	case *types.MsgRemoveLiquidity:
		msg = i
	}

	standardDenom := k.GetStandardDenom(ctx)

	pool, exists := k.GetPoolByLptDenom(ctx, msg.WithdrawLiquidity.Denom)
	if !exists {
		return nil, sdkerrors.Wrapf(types.ErrReservePoolNotExists, "liquidity pool token: %s", msg.WithdrawLiquidity.Denom)
	}

	balances, err := k.GetPoolBalances(ctx, pool.EscrowAddress)
	if err != nil {
		return nil, err
	}

	lptDenom := msg.WithdrawLiquidity.Denom
	minTokenDenom := pool.CounterpartyDenom

	standardReserveAmt := balances.AmountOf(standardDenom)
	tokenReserveAmt := balances.AmountOf(minTokenDenom)
	liquidityReserve := k.bk.GetSupply(ctx, lptDenom).Amount
	if standardReserveAmt.LT(msg.MinStandardAmt) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientFunds, fmt.Sprintf("insufficient %s funds, user expected: %s, actual: %s", standardDenom, msg.MinStandardAmt.String(), standardReserveAmt.String()))
	}
	if tokenReserveAmt.LT(msg.MinToken) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientFunds, fmt.Sprintf("insufficient %s funds, user expected: %s, actual: %s", minTokenDenom, msg.MinToken.String(), tokenReserveAmt.String()))
	}
	if liquidityReserve.LT(msg.WithdrawLiquidity.Amount) {
		return nil, sdkerrors.Wrap(types.ErrInsufficientFunds, fmt.Sprintf("insufficient %s funds, user expected: %s, actual: %s", lptDenom, msg.WithdrawLiquidity.Amount.String(), liquidityReserve.String()))
	}

	// calculate amount of UNI to be burned for sender
	// and coin amount to be returned
	irisWithdrawnAmt := msg.WithdrawLiquidity.Amount.Mul(standardReserveAmt).Quo(liquidityReserve)
	tokenWithdrawnAmt := msg.WithdrawLiquidity.Amount.Mul(tokenReserveAmt).Quo(liquidityReserve)

	irisWithdrawCoin := sdk.NewCoin(standardDenom, irisWithdrawnAmt)
	tokenWithdrawCoin := sdk.NewCoin(minTokenDenom, tokenWithdrawnAmt)
	deductUniCoin := msg.WithdrawLiquidity

	if irisWithdrawCoin.Amount.LT(msg.MinStandardAmt) {
		return nil, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("iris amount not met, user expected: no less than %s, actual: %s", sdk.NewCoin(standardDenom, msg.MinStandardAmt).String(), irisWithdrawCoin.String()))
	}
	if tokenWithdrawCoin.Amount.LT(msg.MinToken) {
		return nil, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("token amount not met, user expected: no less than %s, actual: %s", sdk.NewCoin(minTokenDenom, msg.MinToken).String(), tokenWithdrawCoin.String()))
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveLiquidity,
			sdk.NewAttribute(types.AttributeValueSender, msg.Sender),
			sdk.NewAttribute(types.AttributeValueTokenPair, types.GetTokenPairByDenom(minTokenDenom, standardDenom)),
		),
	)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	poolAddr, err := sdk.AccAddressFromBech32(pool.EscrowAddress)
	if err != nil {
		return nil, err
	}

	return k.removeLiquidity(ctx, poolAddr, sender, deductUniCoin, irisWithdrawCoin, tokenWithdrawCoin)
}

func (k Keeper) removeLiquidity(ctx sdk.Context, poolAddr, sender sdk.AccAddress, deductUniCoin, irisWithdrawCoin, tokenWithdrawCoin sdk.Coin) (sdk.Coins, error) {
	deltaCoins := sdk.NewCoins(deductUniCoin)

	// send liquidity vouchers to be burned from sender account to module account
	if err := k.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, deltaCoins); err != nil {
		return nil, err
	}
	// burn liquidity vouchers of reserve pool from module account
	if err := k.bk.BurnCoins(ctx, types.ModuleName, deltaCoins); err != nil {
		return nil, err
	}

	// transfer withdrawn liquidity from coinswap reserve pool account to sender account
	coins := sdk.NewCoins(irisWithdrawCoin, tokenWithdrawCoin)

	return coins, k.bk.SendCoins(ctx, poolAddr, sender, coins)
}

func (k Keeper) removeUnilateralLiquidity(ctx sdk.Context, msg *types.MsgRemoveUnilateralLiquidity) (sdk.Coins, error) {
	var (
		targetTokenDenom      string
		counterpartTokenDenom string
	)

	pool, exist := k.GetPoolBySequenceId(ctx, msg.PoolId)
	if exist != true {
		return sdk.Coins{}, sdkerrors.Wrap(types.ErrReservePoolNotExists, fmt.Sprintf("liquidity pool id: %d ", msg.PoolId))
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdk.Coins{}, err
	}

	poolAddr, err := sdk.AccAddressFromBech32(pool.EscrowAddress)
	if err != nil {
		return sdk.Coins{}, err
	}

	balances, err := k.GetPoolBalances(ctx, pool.EscrowAddress)
	if err != nil {
		return sdk.Coins{}, err
	}

	lptDenom := pool.LptDenom
	targetTokenDenom = msg.MinToken.Denom
	counterpartTokenDenom = pool.CounterpartyDenom

	if targetTokenDenom != pool.StandardDenom {
		counterpartTokenDenom = pool.StandardDenom
	}

	targetBalanceAmt := balances.AmountOf(targetTokenDenom)
	counterpartBalanceAmt := balances.AmountOf(counterpartTokenDenom)
	lptBalanceAmt := k.bk.GetSupply(ctx, lptDenom).Amount

	// Calculate Withdrawn Amount
	// t_withdrawn = t_balance * delta_lpt / lpt_balance
	// c_withdrawn = c_balance * delta_lpt / lpt_balance
	//
	// Calculate Swap Amount
	// As `(t_balance - t_withdraw)(c_balance - c_withdraw) = (t_balance - t_withdraw - t_swap) * c_balance`,
	// We get `t_swap = (t_balance - t_withdraw) * c_withdraw / c_balance`
	targetWithdrawnAmt := targetBalanceAmt.Mul(msg.ExactLiquidity).Quo(lptBalanceAmt)
	counterpartWithdrawnAmt := counterpartBalanceAmt.Mul(msg.ExactLiquidity).Quo(lptBalanceAmt)
	targetSwapAmt := targetBalanceAmt.Sub(targetWithdrawnAmt).Mul(counterpartWithdrawnAmt).Quo(counterpartBalanceAmt)

	// deduce with fee
	// target_amt = (target_withdraw + target_swap ) * ( 1 - fee_unilateral)
	// fee_unilateral = numerator / denominator
	deltaFeeUnilateral := sdk.OneDec().Sub(k.GetParams(ctx).UnilateralLiquidityFee)
	numerator := sdk.NewIntFromBigInt(deltaFeeUnilateral.BigInt())
	denominator := sdk.NewIntWithDecimal(1, sdk.Precision)
	targetTokenAmt := targetWithdrawnAmt.Add(targetSwapAmt).Mul(numerator).Quo(denominator)

	if targetTokenAmt.LT(msg.MinToken.Amount) {
		return nil, sdkerrors.Wrap(types.ErrConstraintNotMet, fmt.Sprintf("token withdrawn amount not met, user expected: no less than %s, actual: %s", msg.MinToken.String(), sdk.NewCoin(targetTokenDenom, targetTokenAmt).String()))
	}

	// event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveUnilateralLiquidity,
			sdk.NewAttribute(types.AttributeValueSender, msg.Sender),
			sdk.NewAttribute(types.AttributeValueTokenUnilateral, targetTokenDenom),
		),
	)

	// send lpt and burn lpt
	lptCoins := sdk.NewCoins(sdk.NewCoin(lptDenom, msg.ExactLiquidity))
	if err := k.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, lptCoins); err != nil {
		return nil, err
	}
	if err := k.bk.BurnCoins(ctx, types.ModuleName, lptCoins); err != nil {
		return nil, err
	}

	// send withdraw coins
	coins := sdk.NewCoins(sdk.NewCoin(targetTokenDenom, targetTokenAmt))

	return coins, k.bk.SendCoins(ctx, poolAddr, sender, coins)
}

// GetParams gets the parameters for the coinswap module.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var swapParams types.Params
	k.paramSpace.GetParamSet(ctx, &swapParams)
	return swapParams
}

// SetParams sets the parameters for the coinswap module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// SetStandardDenom sets the standard denom for the coinswap module.
func (k Keeper) SetStandardDenom(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	denomWrap := gogotypes.StringValue{Value: denom}
	bz := k.cdc.MustMarshal(&denomWrap)
	store.Set(types.KeyStandardDenom, bz)
}

// GetStandardDenom returns the standard denom of the coinswap module.
func (k Keeper) GetStandardDenom(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyStandardDenom)

	var denomWrap = gogotypes.StringValue{}
	k.cdc.MustUnmarshal(bz, &denomWrap)
	return denomWrap.Value
}
