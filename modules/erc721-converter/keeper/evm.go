package keeper

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/irisnet/irismod/modules/erc721-converter/contracts"

	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/evmos/ethermint/server/config"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

func (k Keeper) DeployERC721Contract(ctx sdk.Context, class nfttypes.Class) (common.Address, error) {
	contractArgs, err := contracts.ERC721PresetMinterPauserAutoIdContract.ABI.Pack(
		"",
		class.Name,
		class.Symbol,
		class.Uri,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "class metadata is invalid %s: %s", class.Name, err.Error())
	}
	data := make([]byte, len(contracts.ERC721PresetMinterPauserAutoIdContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.ERC721PresetMinterPauserAutoIdContract.Bin)], contracts.ERC721PresetMinterPauserAutoIdContract.Bin)
	copy(data[len(contracts.ERC721PresetMinterPauserAutoIdContract.Bin):], contractArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.CallEVMWithData(ctx, types.ModuleAddress, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", class.Name)
	}

	return contractAddr, nil
}

// QueryERC721 queries an ERC721 contract
func (k Keeper) QueryERC721(
	ctx sdk.Context,
	contract common.Address,
) (types.ERC721Data, error) {

	var (
		nameRes   types.ERC721StringResponse
		symbolRes types.ERC721StringResponse
	)

	erc721 := contracts.ERC721PresetMinterPauserAutoIdContract.ABI
	// Name
	res, err := k.CallEVM(ctx, erc721, types.ModuleAddress, contract, false, "name")
	if err != nil {
		return types.ERC721Data{}, err
	}

	if err := erc721.UnpackIntoInterface(&nameRes, "name", res.Ret); err != nil {
		return types.ERC721Data{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack name: %s", err.Error(),
		)
	}

	// Symbol
	res, err = k.CallEVM(ctx, erc721, types.ModuleAddress, contract, false, "symbol")
	if err != nil {
		return types.ERC721Data{}, err
	}

	if err := erc721.UnpackIntoInterface(&symbolRes, "symbol", res.Ret); err != nil {
		return types.ERC721Data{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
		)
	}

	return types.NewERC721Data(nameRes.Value, symbolRes.Value), nil
}

// BalanceOf queries an account's balance for a given ERC721 contract
func (k Keeper) BalanceOf(
	ctx sdk.Context,
	abi abi.ABI,
	contract, account common.Address,
) *big.Int {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, "balanceOf", account)
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

// OwnerOf queries an account's owner for a given ERC721 contract
func (k Keeper) OwnerOf(
	ctx sdk.Context,
	abi abi.ABI,
	contract common.Address,
	tokenID *big.Int,
) (common.Address, error) {
	res, err := k.CallEVM(ctx, abi, types.ModuleAddress, contract, false, "ownerOf", tokenID)
	if err != nil {
		return common.Address{}, err
	}

	unpacked, err := abi.Unpack("ownerOf", res.Ret)
	if err != nil || len(unpacked) == 0 {
		return common.Address{}, err
	}

	owner, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, err
	}

	return owner, nil
}

// CallEVM performs a smart contract method call using given args
func (k Keeper) CallEVM(
	ctx sdk.Context,
	abi abi.ABI,
	from, contract common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {
	data, err := abi.Pack(method, args...)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrABIPack,
			errorsmod.Wrap(err, "failed to create transaction data").Error(),
		)
	}

	resp, err := k.CallEVMWithData(ctx, from, &contract, data, commit)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "contract call failed: method '%s', contract '%s'", method, contract)
	}
	return resp, nil
}

// CallEVMWithData performs a smart contract method call using contract data
func (k Keeper) CallEVMWithData(
	ctx sdk.Context,
	from common.Address,
	contract *common.Address,
	data []byte,
	commit bool,
) (*evmtypes.MsgEthereumTxResponse, error) {
	nonce, err := k.accountKeeper.GetSequence(ctx, from.Bytes())
	if err != nil {
		return nil, err
	}

	gasCap := config.DefaultGasCap
	if commit {
		args, err := json.Marshal(evmtypes.TransactionArgs{
			From: &from,
			To:   contract,
			Data: (*hexutil.Bytes)(&data),
		})
		if err != nil {
			return nil, errorsmod.Wrapf(errortypes.ErrJSONMarshal, "failed to marshal tx args: %s", err.Error())
		}

		gasRes, err := k.evmKeeper.EstimateGas(sdk.WrapSDKContext(ctx), &evmtypes.EthCallRequest{
			Args:   args,
			GasCap: config.DefaultGasCap,
		})
		if err != nil {
			return nil, err
		}
		gasCap = gasRes.Gas
	}

	msg := ethtypes.NewMessage(
		from,
		contract,
		nonce,
		big.NewInt(0), // amount
		gasCap,        // gasLimit
		big.NewInt(0), // gasFeeCap
		big.NewInt(0), // gasTipCap
		big.NewInt(0), // gasPrice
		data,
		ethtypes.AccessList{}, // AccessList
		!commit,               // isFake
	)

	res, err := k.evmKeeper.ApplyMessage(ctx, msg, evmtypes.NewNoOpTracer(), commit)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		return nil, errorsmod.Wrap(evmtypes.ErrVMExecution, res.VmError)
	}

	return res, nil
}
