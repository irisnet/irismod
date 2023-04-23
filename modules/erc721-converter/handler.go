package erc721converter

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// NewHandler defines the erc20 module handler instance
func NewHandler(server types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgConvertNFT:
			res, err := server.ConvertNFT(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgConvertERC721:
			res, err := server.ConvertERC721(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			err := errorsmod.Wrapf(errortypes.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, err
		}
	}
}
