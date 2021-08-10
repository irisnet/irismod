package v1110

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	coinswapkeeper "github.com/irisnet/irismod/modules/coinswap/keeper"
	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"
)

func Migrate(ctx sdk.Context, k coinswapkeeper.Keeper, bk bankkeeper.Keeper, ak authkeeper.AccountKeeper) {
	// 1. Query all current liquidity tokens
	var ltpDenoms []string
	for _, coin := range bk.GetSupply(ctx).GetTotal() {
		if strings.HasPrefix(coin.GetDenom(), FormatUniABSPrefix) {
			ltpDenoms = append(ltpDenoms, coin.GetDenom())
		}
	}

	// 2. Create a new liquidity pool based on the results of the first step
	standardDenom := k.GetStandardDenom(ctx)
	var pools = make(map[string]coinswaptypes.Pool, len(ltpDenoms))
	for _, ltpDenom := range ltpDenoms {
		counterpartyDenom := strings.TrimPrefix(ltpDenom, FormatUniABSPrefix)
		pools[ltpDenom] = k.CreatePool(ctx, counterpartyDenom)
		//3. Transfer tokens from the old liquidity to the newly created liquidity pool
		migratePool(ctx, bk, pools[ltpDenom], ltpDenom, standardDenom)
	}

	// 4. Traverse all accounts and modify the old liquidity token to the new liquidity token
	ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		balances := bk.GetAllBalances(ctx, account.GetAddress())
		for _, ltpDenom := range ltpDenoms {
			amount := balances.AmountOf(ltpDenom)
			if sdk.ZeroInt().Equal(amount) {
				return false
			}
			originLptCoin := sdk.NewCoin(ltpDenom, amount)
			migrateProvider(ctx, originLptCoin, bk, pools[ltpDenom], account.GetAddress())
		}
		return false
	})
}

func migrateProvider(ctx sdk.Context,
	originLptCoin sdk.Coin,
	bk bankkeeper.Keeper,
	pool coinswaptypes.Pool,
	provider sdk.AccAddress,
) {
	//1. Burn the old liquidity tokens
	burnCoins := sdk.NewCoins(originLptCoin)
	// send liquidity vouchers to be burned from sender account to module account
	if err := bk.SendCoinsFromAccountToModule(ctx, provider, coinswaptypes.ModuleName, burnCoins); err != nil {
		panic(err)
	}
	// burn liquidity vouchers of reserve pool from module account
	if err := bk.BurnCoins(ctx, coinswaptypes.ModuleName, burnCoins); err != nil {
		panic(err)
	}

	//2. Issue new liquidity tokens
	mintToken := sdk.NewCoin(pool.LptDenom, originLptCoin.Amount)
	mintTokens := sdk.NewCoins(mintToken)
	if err := bk.MintCoins(ctx, coinswaptypes.ModuleName, mintTokens); err != nil {
		panic(err)
	}
	if err := bk.SendCoinsFromModuleToAccount(ctx, coinswaptypes.ModuleName, provider, mintTokens); err != nil {
		panic(err)
	}
}

func migratePool(ctx sdk.Context,
	bk bankkeeper.Keeper,
	pool coinswaptypes.Pool,
	ltpDenom, standardDenom string,
) {
	counterpartyDenom := strings.TrimPrefix(ltpDenom, FormatUniABSPrefix)
	originPoolAddress := GetReservePoolAddr(ltpDenom)

	//Query the amount of the original liquidity pool account
	originPoolBalances := bk.GetAllBalances(ctx, originPoolAddress)
	transferCoins := sdk.NewCoins(
		sdk.NewCoin(standardDenom, originPoolBalances.AmountOf(standardDenom)),
		sdk.NewCoin(counterpartyDenom, originPoolBalances.AmountOf(counterpartyDenom)),
	)

	dstPoolAddress, err := sdk.AccAddressFromBech32(pool.EscrowAddress)
	if err != nil {
		panic(err)
	}

	err = bk.SendCoins(ctx, originPoolAddress, dstPoolAddress, transferCoins)
	if err != nil {
		panic(err)
	}
}
