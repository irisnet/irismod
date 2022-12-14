package keeper

import (
	"context"

	"github.com/irisnet/irismod/modules/rental/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) SetUser(goCtx context.Context, msg *types.MsgSetUser) (*types.MsgSetUserResponse, error) {
	// todo: set user
	// todo: event
	return &types.MsgSetUserResponse{}, nil
}
