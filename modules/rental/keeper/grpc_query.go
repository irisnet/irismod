package keeper

import (
	"context"

	"github.com/irisnet/irismod/modules/rental/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) UserOf(context.Context, *types.MsgUserOfRequest) (*types.MsgUserOfResponse, error) {
	panic("Fixme")
}
func (k Keeper) UserExpires(context.Context, *types.MsgUserExpiresRequest) (*types.MsgUserExpiresResponse, error) {
	panic("Fixme")
}

func (k Keeper) HaveUser(context.Context, *types.MsgHaveUserRequest) (*types.MsgHaveUserResponse, error) {
	panic("Fixme")
}
