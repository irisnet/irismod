package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

func (m msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse,
	error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	//check valid begin height
	if ctx.BlockHeight() > int64(msg.BeginHeight) {
		return nil, sdkerrors.Wrapf(
			types.ErrExpiredHeight,
			"The current block height[%d] is greater than BeginHeight[%d]",
			ctx.BlockHeight(),
			msg.BeginHeight,
		)
	}

	//check valid lp token denom
	if supply := m.Keeper.bk.GetSupply(ctx, msg.LpTokenDenom); supply.IsZero() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidLPToken,
			"The lp token denom[%s] is not exist",
			msg.LpTokenDenom,
		)
	}
	if err = m.Keeper.CreatePool(ctx,
		msg.Name,
		msg.Description,
		msg.LpTokenDenom,
		msg.BeginHeight,
		msg.RewardPerBlock.Sort(),
		msg.TotalReward.Sort(),
		msg.Destructible,
		creator,
	); err != nil {
		return nil, err
	}
	return &types.MsgCreatePoolResponse{}, nil
}

func (m msgServer) DestroyPool(goCtx context.Context, msg *types.MsgDestroyPool) (*types.MsgDestroyPoolResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.DestroyPool(ctx, msg.PoolName, creator); err != nil {
		return nil, err
	}
	return &types.MsgDestroyPoolResponse{}, nil
}

func (m msgServer) AppendReward(goCtx context.Context, msg *types.MsgAppendReward) (*types.MsgAppendRewardResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	remaining, err := m.Keeper.AppendReward(ctx, msg.PoolName, msg.Amount, creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgAppendRewardResponse{RemainingReward: remaining}, nil
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
