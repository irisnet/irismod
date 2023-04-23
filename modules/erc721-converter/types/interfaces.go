package types

import (
	"context"

	"github.com/irisnet/irismod/modules/nft/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/x/evm/statedb"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	irsmodnfttypes "github.com/irisnet/irismod/modules/nft/types"
)

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
	EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)
	ApplyMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
}

// NFTKeeper defines the expected interface needed to retrieve the NFT denom.
type NFTKeeper interface {
	GetClass(ctx sdk.Context, classID string) (class nfttypes.Class, found bool)
	SaveClass(ctx sdk.Context, class nfttypes.Class) error
	HasClass(ctx sdk.Context, classID string) bool
	Mint(ctx sdk.Context, token nfttypes.NFT, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, nftID string) error
	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error
	GetNFT(ctx sdk.Context, classID string, nftID string) (nft nfttypes.NFT, found bool)
	HasNFT(ctx sdk.Context, classID string, nftID string) bool
	GetOwner(ctx sdk.Context, classID string, nftID string) sdk.AccAddress
}

// IRISModNFTKeeper defines the expected interface needed to retrieve the IRSMod NFT denom.
type IRISModNFTKeeper interface {
	SaveDenom(ctx sdk.Context, id, name, schema, symbol string, creator sdk.AccAddress, mintRestricted, updateRestricted bool, description, uri, uriHash, data string) error
	HasDenom(ctx sdk.Context, denomID string) bool
	GetDenomInfo(ctx sdk.Context, denomID string) (*irsmodnfttypes.Denom, error)

	SaveNFT(ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData string, receiver sdk.AccAddress) error
	RemoveNFT(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error
	TransferOwnership(ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData string, srcOwner, dstOwner sdk.AccAddress) error
	HasNFT(ctx sdk.Context, denomID, tokenID string) bool
	GetNFT(ctx sdk.Context, denomID, tokenID string) (nft exported.NFT, err error)

	Authorize(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error
}

type ERC721Keeper interface {
	GetTokenPairID(ctx sdk.Context, token string) []byte
	GetTokenPair(ctx sdk.Context, id []byte) (TokenPair, bool)
	ConvertERC721(ctx context.Context, msg *MsgConvertERC721) (*MsgConvertERC721Response, error)
}
