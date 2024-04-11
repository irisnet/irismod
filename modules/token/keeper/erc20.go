package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/irisnet/irismod/contracts"
	"github.com/irisnet/irismod/modules/token/types"
)

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
	_, err = k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", name)
	}
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
	abi := contracts.ERC20TokenContract.ABI
	res, err := k.CallEVM(ctx, abi, k.moduleAddress(), contract, true, "mint", to, amount)
	if err != nil {
		return err
	}

	if res.Failed() {
		return errorsmod.Wrapf(types.ErrVMExecution, "failed to mint %s", amount)
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
	abi := contracts.ERC20TokenContract.ABI
	res, err := k.CallEVM(ctx, abi, k.moduleAddress(), contract, true, "burn", from, amount)
	if err != nil {
		return err
	}

	if res.Failed() {
		return errorsmod.Wrapf(types.ErrVMExecution, "failed to burn %s", amount)
	}
	return nil
}


func (k Keeper) moduleAddress() common.Address {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	return common.BytesToAddress(moduleAddr.Bytes())
}
