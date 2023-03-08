package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/irisnet/irismod/modules/token/types"
)

type Keeper struct {
	storeKey         storetypes.StoreKey
	cdc              codec.Codec
	bankKeeper       types.BankKeeper
	paramSpace       paramstypes.Subspace
	blockedAddrs     map[string]bool
	feeCollectorName string
	registry         types.SwapRegistry
}

func NewKeeper(
	cdc codec.Codec,
	key storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
	bankKeeper types.BankKeeper,
	blockedAddrs map[string]bool,
	feeCollectorName string,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:         key,
		cdc:              cdc,
		paramSpace:       paramSpace,
		bankKeeper:       bankKeeper,
		feeCollectorName: feeCollectorName,
		blockedAddrs:     blockedAddrs,
		registry:         make(types.SwapRegistry),
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irismod/%s", types.ModuleName))
}

// IssueToken issues a new token
func (k Keeper) IssueToken(
	ctx sdk.Context,
	symbol string,
	name string,
	minUnit string,
	scale uint32,
	initialSupply uint64,
	maxSupply uint64,
	mintable bool,
	owner sdk.AccAddress,
) error {
	token := types.NewToken(
		symbol, name, minUnit, scale, initialSupply,
		maxSupply, mintable, owner,
	)

	if err := k.AddToken(ctx, token); err != nil {
		return err
	}

	precision := sdkmath.NewIntWithDecimal(1, int(token.Scale))
	initialCoin := sdk.NewCoin(
		token.MinUnit,
		sdk.NewIntFromUint64(token.InitialSupply).Mul(precision),
	)

	mintCoins := sdk.NewCoins(initialCoin)

	// mint coins into module account
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return err
	}

	// sent coins to owner's account
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, mintCoins)
}

// EditToken edits the specified token
func (k Keeper) EditToken(
	ctx sdk.Context,
	symbol string,
	name string,
	maxSupply uint64,
	mintable types.Bool,
	owner sdk.AccAddress,
) error {
	// get the destination token
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}

	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", owner, symbol)
	}

	if maxSupply > 0 {
		issuedAmt := k.getTokenSupply(ctx, token.MinUnit)
		issuedMainUnitAmt := issuedAmt.Quo(sdkmath.NewIntWithDecimal(1, int(token.Scale)))

		if sdk.NewIntFromUint64(maxSupply).LT(issuedMainUnitAmt) {
			return sdkerrors.Wrapf(types.ErrInvalidMaxSupply, "max supply must not be less than %s", issuedMainUnitAmt)
		}

		token.MaxSupply = maxSupply
	}

	if name != types.DoNotModify {
		token.Name = name

		metadata, _ := k.bankKeeper.GetDenomMetaData(ctx, token.MinUnit)
		metadata.Description = name

		k.bankKeeper.SetDenomMetaData(ctx, metadata)
	}

	if mintable != types.Nil {
		token.Mintable = mintable.ToBool()
	}

	k.setToken(ctx, token)

	return nil
}

// TransferTokenOwner transfers the owner of the specified token to a new one
func (k Keeper) TransferTokenOwner(
	ctx sdk.Context,
	symbol string,
	srcOwner sdk.AccAddress,
	dstOwner sdk.AccAddress,
) error {
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}

	if srcOwner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", srcOwner, symbol)
	}

	token.Owner = dstOwner.String()

	// update token
	k.setToken(ctx, token)

	// reset all indices
	k.resetStoreKeyForQueryToken(ctx, token.Symbol, srcOwner, dstOwner)

	return nil
}

// MintToken mints the specified amount of token to the specified recipient
// NOTE: empty owner means that the external caller is responsible to manage the token authority
func (k Keeper) MintToken(
	ctx sdk.Context,
	symbol string,
	amount uint64,
	recipient sdk.AccAddress,
	owner sdk.AccAddress,
) error {
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}

	if owner.String() != token.Owner {
		return sdkerrors.Wrapf(types.ErrInvalidOwner, "the address %s is not the owner of the token %s", owner, symbol)
	}

	if !token.Mintable {
		return sdkerrors.Wrapf(types.ErrNotMintable, "%s", symbol)
	}

	supply := k.getTokenSupply(ctx, token.MinUnit)
	precision := sdkmath.NewIntWithDecimal(1, int(token.Scale))
	mintableAmt := sdk.NewIntFromUint64(token.MaxSupply).Mul(precision).Sub(supply)
	mintableMainAmt := mintableAmt.Quo(precision).Uint64()

	if amount > mintableMainAmt {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			"the amount exceeds the mintable token amount; expected (0, %d], got %d",
			mintableMainAmt, amount,
		)
	}

	mintCoin := sdk.NewCoin(token.MinUnit, sdk.NewIntFromUint64(amount).Mul(precision))
	mintCoins := sdk.NewCoins(mintCoin)

	// mint coins
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return err
	}

	if recipient.Empty() {
		recipient = owner
	}

	// sent coins to the recipient account
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, mintCoins)
}

// BurnToken burns the specified amount of token
func (k Keeper) BurnToken(
	ctx sdk.Context,
	symbol string,
	amount uint64,
	owner sdk.AccAddress,
) error {
	token, err := k.getTokenBySymbol(ctx, symbol)
	if err != nil {
		return err
	}

	precision := sdkmath.NewIntWithDecimal(1, int(token.Scale))
	burnCoin := sdk.NewCoin(token.GetMinUnit(), sdk.NewIntFromUint64(amount).Mul(precision))
	burnCoins := sdk.NewCoins(burnCoin)

	// burn coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, burnCoins); err != nil {
		return err
	}

	k.AddBurnCoin(ctx, burnCoin)

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
}

// SwapFeeToken swap the fee token
func (k Keeper) SwapFeeToken(
	ctx sdk.Context,
	feePaid sdk.Coin,
	sender sdk.AccAddress,
	recipient sdk.AccAddress,
) (sdk.Coin, error) {
	burnedCoin, mintedCoin, err := k.calcFeeTokenMinted(ctx, feePaid)
	if err != nil {
		return sdk.Coin{}, err
	}
	burnedCoins := sdk.NewCoins(burnedCoin)
	// burn coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, burnedCoins); err != nil {
		return sdk.Coin{}, err
	}
	//k.AddBurnCoin(ctx, burnedCoin)
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnedCoins); err != nil {
		return sdk.Coin{}, err
	}

	// mint coins
	mintedCoins := sdk.NewCoins(mintedCoin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintedCoins); err != nil {
		return sdk.Coin{}, err
	}

	if recipient == nil {
		recipient = sender
	}
	return mintedCoin, k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, mintedCoins)
}

func (k Keeper) WithSwapRegistry(registry types.SwapRegistry) Keeper {
	k.registry = registry
	return k
}

func (k Keeper) calcFeeTokenMinted(ctx sdk.Context, feePaid sdk.Coin) (burned sdk.Coin, minted sdk.Coin, err error) {
	tokenBurned, err := k.GetToken(ctx, feePaid.Denom)
	if err != nil {
		return burned, minted, err
	}

	swapParams, ok := k.registry[tokenBurned.GetMinUnit()]
	if !ok {
		return burned, minted, types.ErrInvalidSwap
	}

	tokenMinted, err := k.GetToken(ctx, swapParams.MinUnit)
	if err != nil {
		return burned, minted, err
	}

	coinBurned, err := tokenBurned.ToMinCoin(sdk.NewDecCoinFromCoin(feePaid))
	if err != nil {
		return burned, minted, err
	}

	var multiple sdk.Dec

	if tokenMinted.GetScale() >= tokenBurned.GetScale() {
		multiple = sdk.NewDecFromInt(sdkmath.NewIntWithDecimal(1, int(tokenMinted.GetScale()-tokenBurned.GetScale())))
	} else {
		multiple = sdk.NewDecWithPrec(1, int64(tokenBurned.GetScale()-tokenMinted.GetScale()))
	}

	amountMinted := multiple.MulInt(coinBurned.Amount).Mul(swapParams.Ratio).TruncateInt()
	return coinBurned, sdk.NewCoin(swapParams.MinUnit, amountMinted), nil
}
