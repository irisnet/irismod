package keeper

import (
	"context"

	"github.com/irisnet/irismod/modules/rental/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) User(context.Context, *types.QueryUserRequest) (*types.QueryUserResponse, error) {
	panic("Fixme")
}
func (k Keeper) Expires(context.Context, *types.QueryExpiresRequest) (*types.QueryExpiresResponse, error) {
	panic("Fixme")
}

func (k Keeper) HasUser(context.Context, *types.QueryHasUserRequest) (*types.QueryHasUserResponse, error) {
	panic("Fixme")
}
