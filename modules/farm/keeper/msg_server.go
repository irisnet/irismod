package keeper

import (
	"context"

	"github.com/irisnet/irismod/modules/farm/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the farm MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (msgServer) CreatePool(context.Context, *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	return &types.MsgCreatePoolResponse{}, nil
}

func (msgServer) DestroyPool(context.Context, *types.MsgDestroyPool) (*types.MsgDestroyPoolResponse, error) {
	return &types.MsgDestroyPoolResponse{}, nil
}

func (msgServer) AppendReward(context.Context, *types.MsgAppendReward) (*types.MsgAppendRewardResponse, error) {
	return &types.MsgAppendRewardResponse{}, nil
}

func (msgServer) Stake(context.Context, *types.MsgStake) (*types.MsgStakeResponse, error) {
	return &types.MsgStakeResponse{}, nil
}

func (msgServer) Unstake(context.Context, *types.MsgUnstake) (*types.MsgUnstakeResponse, error) {
	return &types.MsgUnstakeResponse{}, nil
}

func (msgServer) Harvest(context.Context, *types.MsgHarvest) (*types.MsgHarvestResponse, error) {
	return &types.MsgHarvestResponse{}, nil
}
