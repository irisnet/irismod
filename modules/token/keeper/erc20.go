package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/irisnet/irismod/contracts"
	"github.com/irisnet/irismod/modules/token/types"
	v1 "github.com/irisnet/irismod/modules/token/types/v1"
)

// SwapToERC20 executes a swap from a native token to an ERC20 token
//
// Parameters:
//   - ctx: the context
//   - sender: the sender of the amount
//   - receiver: the receiver of the erc20 token
//   - amount:  the amount to be swapped
//
// Returns:
//   - error: error if any.
func (k Keeper) SwapToERC20(
	ctx sdk.Context,
	sender sdk.AccAddress,
	receiver common.Address,
	amount sdk.Coin,
) error {
	receiverAcc := k.accountKeeper.GetAccount(ctx, sdk.AccAddress(receiver.Bytes()))
	if receiverAcc != nil {
		if !k.evmKeeper.SupportedKey(receiverAcc.GetPubKey()) {
			return errorsmod.Wrapf(types.ErrUnsupportedKey, "key %s", receiverAcc.GetPubKey())
		}
	}

	token, err := k.getTokenByMinUnit(ctx, amount.Denom)
	if err != nil {
		return err
	}
	if len(token.Contract) == 0 {
		return errorsmod.Wrapf(types.ErrERC20NotDeployed, "token: %s not deployed", amount.Denom)
	}
	contract := common.HexToAddress(token.Contract)

	amt := sdk.NewCoins(amount)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amt); err != nil {
		return err
	}

	if err := k.MintERC20(ctx, contract, receiver, amount.Amount.Uint64()); err != nil {
		return err
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, amt); err != nil {
		return err
	}

	ctx.EventManager().EmitTypedEvent(&v1.EventSwapToERC20{
		Amount:   amount,
		Sender:   sender.String(),
		Receiver: receiver.String(),
	})
	return nil
}

// SwapFromERC20 executes a swap from an ERC20 token to a native token.
//
// Parameters:
//
//	ctx - the context in which the swap is executed
//	sender - the address of the sender
//	receiver - the address of the receiver
//	contract - the address of the ERC20 contract
//	amount - the amount of tokens to swap
//
// Return type: error
func (k Keeper) SwapFromERC20(
	ctx sdk.Context,
	sender common.Address,
	receiver sdk.AccAddress,
	contract common.Address,
	amount *big.Int,
) error {
	token, err := k.getTokenByContract(ctx, contract.String())
	if err != nil {
		return err
	}

	balance := k.BalanceOf(ctx, contract, sender)
	if r := balance.Cmp(amount); r < 0 {
		return errorsmod.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"balance: %d, swap: %d",
			balance,
			amount,
		)
	}
	if err := k.BurnERC20(ctx, contract, sender, amount.Uint64()); err != nil {
		return err
	}

	amt := sdk.NewCoins(sdk.NewCoin(token.MinUnit, sdkmath.NewIntFromBigInt(amount)))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amt); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, amt); err != nil {
		return err
	}

	ctx.EventManager().EmitTypedEvent(&v1.EventSwapFromERC20{
		Contract: contract.String(),
		Amount:   amount.Int64(),
		Sender:   sender.String(),
		Receiver: receiver.String(),
	})
	return nil
}

// DeployERC20 deploys an ERC20 token contract.
//
// Parameters:
//   - ctx: the context
//   - name: the name of the token
//   - symbol: the symbol of the token
//   - scale: the scale of the token
//
// Returns:
//   - Address: the contract address.
//   - error: error if any.
func (k Keeper) DeployERC20(
	ctx sdk.Context,
	name string,
	symbol string,
	minUnit string,
	scale int8,
) (common.Address, error) {
	contractArgs, err := contracts.ERC20TokenContract.ABI.Pack(
		"",
		name,
		symbol,
		scale,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "erc20 metadata is invalid %s: %s", name, err.Error())
	}
	deployer := k.moduleAddress()

	data := make([]byte, len(contracts.ERC20TokenContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.ERC20TokenContract.Bin)], contracts.ERC20TokenContract.Bin)
	copy(data[len(contracts.ERC20TokenContract.Bin):], contractArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	contractAddr := crypto.CreateAddress(deployer, nonce)
	result, err := k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", name)
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(types.ErrVMExecution, "failed to deploy contract for %s, reason: %s", name, result.Revert())
	}

	ctx.EventManager().EmitTypedEvent(&v1.EventDeployERC20{
		Symbol:   symbol,
		Name:     name,
		Scale:    uint32(scale),
		MinUnit:  minUnit,
		Contract: contractAddr.String(),
	})
	return contractAddr, nil
}

// BalanceOf retrieves the balance of a specific account in the contract.
//
// Parameters:
//   - ctx: the sdk.Context for the function
//   - contract: the address of the contract
//   - account: the address of the account to retrieve the balance for
//
// Returns:
//   - *big.Int: the balance of the specified account
func (k Keeper) BalanceOf(
	ctx sdk.Context,
	contract, account common.Address,
) *big.Int {
	abi := contracts.ERC20TokenContract.ABI
	res, err := k.CallEVM(ctx, abi, k.moduleAddress(), contract, false, "balanceOf", account)
	if err != nil {
		return nil
	}

	unpacked, err := abi.Unpack("balanceOf", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return nil
	}

	balance, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil
	}

	return balance
}

// MintERC20 mints ERC20 tokens to an account.
//
// Parameters:
//   - ctx: the sdk.Context for the function
//   - contract: the address of the contract
//   - to: the address of the receiver
//   - amount: the amount to mint
//
// Returns:
//   - err : error if any
func (k Keeper) MintERC20(
	ctx sdk.Context,
	contract, to common.Address,
	amount uint64,
) error {
	balanceBefore := k.BalanceOf(ctx, contract, to)
	abi := contracts.ERC20TokenContract.ABI
	res, err := k.CallEVM(ctx, abi, k.moduleAddress(), contract, true, "mint", to, amount)
	if err != nil {
		return err
	}

	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to mint contract: %s, reason: %s",
			contract.String(),
			res.Revert(),
		)
	}

	balanceAfter := k.BalanceOf(ctx, contract, to)
	expectBalance := big.NewInt(0).Add(balanceBefore, big.NewInt(int64(amount)))
	if r := expectBalance.Cmp(balanceAfter); r != 0 {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to mint contract: %s, expect %d, actual %d, ",
			contract.String(),
			expectBalance.Int64(),
			balanceAfter.Int64(),
		)
	}
	return nil
}

// BurnERC20 burns a specific amount of ERC20 tokens from a given contract and address.
//
// Parameters:
//   - ctx: the context in which the transaction is executed
//   - contract: the contract address of the ERC20 token
//   - from: the address from which the tokens are burned
//   - amount: the amount of tokens to burn
//
// Returns an error.
func (k Keeper) BurnERC20(
	ctx sdk.Context,
	contract, from common.Address,
	amount uint64,
) error {
	balanceBefore := k.BalanceOf(ctx, contract, from)
	abi := contracts.ERC20TokenContract.ABI
	res, err := k.CallEVM(ctx, abi, k.moduleAddress(), contract, true, "burn", from, amount)
	if err != nil {
		return err
	}

	if res.Failed() {
		return errorsmod.Wrapf(types.ErrVMExecution, "failed to burn %d", amount)
	}

	balanceAfter := k.BalanceOf(ctx, contract, from)
	expectBalance := big.NewInt(0).Sub(balanceBefore, big.NewInt(int64(amount)))
	if r := expectBalance.Cmp(balanceAfter); r != 0 {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to burn contract: %s, expect %d, actual %d, ",
			contract.String(),
			expectBalance.Int64(),
			balanceAfter.Int64(),
		)
	}
	return nil
}

func (k Keeper) moduleAddress() common.Address {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	return common.BytesToAddress(moduleAddr.Bytes())
}
