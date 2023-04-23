package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrERC721Disabled          = errorsmod.Register(ModuleName, 2, "erc721 module is disabled")
	ErrInternalTokenPair       = errorsmod.Register(ModuleName, 3, "internal ethereum token mapping error")
	ErrTokenPairNotFound       = errorsmod.Register(ModuleName, 4, "token pair not found")
	ErrTokenPairAlreadyExists  = errorsmod.Register(ModuleName, 5, "token pair already exists")
	ErrUndefinedOwner          = errorsmod.Register(ModuleName, 6, "undefined owner of contract pair")
	ErrBalanceInvariance       = errorsmod.Register(ModuleName, 7, "post transfer balance invariant failed")
	ErrUnexpectedEvent         = errorsmod.Register(ModuleName, 8, "unexpected event")
	ErrABIPack                 = errorsmod.Register(ModuleName, 9, "contract ABI pack failed")
	ErrABIUnpack               = errorsmod.Register(ModuleName, 10, "contract ABI unpack failed")
	ErrEVMDenom                = errorsmod.Register(ModuleName, 11, "EVM denomination registration")
	ErrEVMCall                 = errorsmod.Register(ModuleName, 12, "EVM call unexpected error")
	ErrERC721TokenPairDisabled = errorsmod.Register(ModuleName, 13, "erc721 token pair is disabled")
	ErrUnauthorized            = errorsmod.Register(ModuleName, 14, "unauthorized address")
	ErrClassNotFound           = errorsmod.Register(ModuleName, 15, "class not found")
)
