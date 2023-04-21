package keeper

import (
	"context"
	"encoding/hex"

	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/record/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the record MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) CreateRecord(goCtx context.Context, msg *types.MsgCreateRecord) (*types.MsgCreateRecordResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.NewRecord(tmhash.Sum(ctx.TxBytes()), msg.Contents, creator)
	recordId := m.Keeper.AddRecord(ctx, record)

	hexID := hex.EncodeToString(recordId)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateRecord,
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyRecordID, hex.EncodeToString(recordId)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreateRecordResponse{
		Id: hexID,
	}, nil
}

func (m msgServer) GrantRecord(goCtx context.Context, msg *types.MsgGrantRecord) (*types.MsgGrantRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	grantRecord := types.NewGrantRecord(tmhash.Sum(ctx.TxBytes()), msg.Id, msg.Key, msg.Pubkey, msg.Creator)
	grantRecordId := m.Keeper.AddGrantRecord(ctx, grantRecord)

	hexID := hex.EncodeToString(grantRecordId)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeGrantRecord,

			sdk.NewAttribute(types.AttributeKeyGrantRecordID, hexID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgGrantRecordResponse{
		Id: hexID,
	}, nil
}
