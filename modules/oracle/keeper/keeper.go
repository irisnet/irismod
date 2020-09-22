package keeper

import (
	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	servicetypes "github.com/irisnet/irismod/modules/service/types"

	"github.com/irisnet/irismod/modules/oracle/types"
)

// Keeper
type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey

	sk types.ServiceKeeper

	paramSpace paramtypes.Subspace
}

// NewKeeper
func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	sk types.ServiceKeeper,
) Keeper {
	keeper := Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		sk:         sk,
		paramSpace: paramSpace,
	}
	_ = sk.RegisterResponseCallback(types.ModuleName, keeper.HandlerResponse)
	_ = sk.RegisterStateCallback(types.ModuleName, keeper.HandlerStateChanged)
	return keeper
}

//CreateFeed create a stopped feed
func (k Keeper) CreateFeed(ctx sdk.Context, msg *types.MsgCreateFeed) error {
	if _, found := k.GetFeed(ctx, msg.FeedName); found {
		return sdkerrors.Wrapf(types.ErrExistedFeedName, msg.FeedName)
	}

	requestContextID, err := k.sk.CreateRequestContext(
		ctx,
		msg.ServiceName,
		msg.Providers,
		msg.Creator,
		msg.Input,
		msg.ServiceFeeCap,
		msg.Timeout,
		false,
		true,
		msg.RepeatedFrequency,
		-1,
		servicetypes.PAUSED,
		msg.ResponseThreshold,
		types.ModuleName,
	)
	if err != nil {
		return err
	}

	k.SetFeed(ctx, types.Feed{
		FeedName:         msg.FeedName,
		AggregateFunc:    msg.AggregateFunc,
		ValueJsonPath:    msg.ValueJsonPath,
		LatestHistory:    msg.LatestHistory,
		RequestContextID: requestContextID,
		Description:      msg.Description,
		Creator:          msg.Creator,
	})
	k.Enqueue(ctx, msg.FeedName, servicetypes.PAUSED)

	return nil
}

//StartFeed start a stopped feed
func (k Keeper) StartFeed(ctx sdk.Context, msg *types.MsgStartFeed) error {
	feed, found := k.GetFeed(ctx, msg.FeedName)
	if !found {
		return sdkerrors.Wrapf(types.ErrUnknownFeedName, msg.FeedName)
	}

	if !msg.Creator.Equals(feed.Creator) {
		return sdkerrors.Wrapf(types.ErrUnauthorized, msg.Creator.String())
	}

	reqCtx, existed := k.sk.GetRequestContext(ctx, feed.RequestContextID)
	if !existed {
		return sdkerrors.Wrapf(types.ErrUnknownFeedName, msg.FeedName)
	}

	//Can not start feed in "running" state
	if reqCtx.State == servicetypes.RUNNING {
		return sdkerrors.Wrapf(types.ErrInvalidFeedState, msg.FeedName)
	}

	if err := k.sk.StartRequestContext(ctx, feed.RequestContextID, feed.Creator); err != nil {
		return err
	}

	k.dequeueAndEnqueue(ctx, msg.FeedName, servicetypes.PAUSED, servicetypes.RUNNING)
	return nil
}

//PauseFeed pause a running feed
func (k Keeper) PauseFeed(ctx sdk.Context, msg *types.MsgPauseFeed) error {
	feed, found := k.GetFeed(ctx, msg.FeedName)
	if !found {
		return sdkerrors.Wrapf(types.ErrUnknownFeedName, msg.FeedName)
	}

	if !msg.Creator.Equals(feed.Creator) {
		return sdkerrors.Wrapf(types.ErrUnauthorized, msg.Creator.String())
	}

	reqCtx, existed := k.sk.GetRequestContext(ctx, feed.RequestContextID)
	if !existed {
		return sdkerrors.Wrapf(types.ErrUnknownFeedName, msg.FeedName)
	}

	//Can only pause feed in "running" state
	if reqCtx.State != servicetypes.RUNNING {
		return sdkerrors.Wrapf(types.ErrInvalidFeedState, msg.FeedName)
	}

	if err := k.sk.PauseRequestContext(ctx, feed.RequestContextID, feed.Creator); err != nil {
		return err
	}

	k.dequeueAndEnqueue(ctx, msg.FeedName, servicetypes.RUNNING, servicetypes.PAUSED)
	return nil
}

//EditFeed edit a feed
func (k Keeper) EditFeed(ctx sdk.Context, msg *types.MsgEditFeed) error {
	feed, found := k.GetFeed(ctx, msg.FeedName)
	if !found {
		return sdkerrors.Wrapf(types.ErrUnknownFeedName, msg.FeedName)
	}

	if !msg.Creator.Equals(feed.Creator) {
		return sdkerrors.Wrapf(types.ErrUnauthorized, msg.Creator.String())
	}

	if err := k.sk.UpdateRequestContext(
		ctx,
		feed.RequestContextID,
		msg.Providers,
		msg.ResponseThreshold,
		msg.ServiceFeeCap,
		msg.Timeout,
		msg.RepeatedFrequency,
		-1,
		msg.Creator,
	); err != nil {
		return err
	}

	if msg.LatestHistory > 0 {
		cnt := k.getFeedValuesCnt(ctx, feed.FeedName)
		if expectCnt := int(msg.LatestHistory); expectCnt < cnt {
			k.deleteOldestFeedValue(ctx, feed.FeedName, cnt-expectCnt)
		}
		feed.LatestHistory = msg.LatestHistory
	}

	if types.IsModified(msg.Description) {
		feed.Description = strings.TrimSpace(msg.Description)
	}

	k.SetFeed(ctx, feed)
	return nil
}

//HandlerResponse is responsible for processing the data returned from the servicetypes module,
//processed by the aggregate function, and then saved
func (k Keeper) HandlerResponse(ctx sdk.Context,
	requestContextID tmbytes.HexBytes,
	responseOutput []string,
	err error) {
	if len(responseOutput) == 0 || err != nil {
		ctx.Logger().Error(
			"Oracle feed failed", "requestContextID",
			requestContextID.String(), "err", err.Error(),
		)
		return
	}

	feed, found := k.GetFeedByReqCtxID(ctx, requestContextID)
	if !found {
		ctx.Logger().Error(
			"Not existed requestContext", "requestContextID", requestContextID.String(),
		)
		return
	}

	reqCtx, existed := k.sk.GetRequestContext(ctx, requestContextID)
	if !existed {
		ctx.Logger().Error(
			"Not existed requestContext", "requestContextID", requestContextID.String(),
		)
		return
	}

	aggregate, err := types.GetAggregateFunc(feed.AggregateFunc)
	if err != nil {
		ctx.Logger().Error(
			"Not existed aggregateFunc", "aggregateFunc", feed.AggregateFunc,
		)
		return
	}

	var data []types.ArgsType
	for _, jsonStr := range responseOutput {
		result := gjson.Get(jsonStr, feed.ValueJsonPath)
		data = append(data, result)
	}

	result := aggregate(data)
	value := types.FeedValue{
		Data:      result,
		Timestamp: ctx.BlockTime(),
	}
	k.SetFeedValue(ctx, feed.FeedName, reqCtx.BatchCounter, feed.LatestHistory, value)

	bz, _ := json.Marshal(value)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSetFeed,
			sdk.NewAttribute(types.AttributeKeyFeedName, feed.FeedName),
			sdk.NewAttribute(types.AttributeKeyFeedValue, string(bz)),
		),
	)
}

//HandlerStateChanged is responsible for update feed state
func (k Keeper) HandlerStateChanged(ctx sdk.Context, requestContextID tmbytes.HexBytes, _ string) {
	reqCtx, existed := k.sk.GetRequestContext(ctx, requestContextID)
	if !existed {
		ctx.Logger().Error(
			"Not existed requestContext", "requestContextID", requestContextID.String(),
		)
		return
	}

	feed, found := k.GetFeedByReqCtxID(ctx, requestContextID)
	if !found {
		ctx.Logger().Error(
			"Not existed feed", "requestContextID", requestContextID.String(),
		)
		return
	}

	var oldState servicetypes.RequestContextState
	switch reqCtx.State {
	case servicetypes.PAUSED:
		oldState = servicetypes.RUNNING
	case servicetypes.RUNNING:
		oldState = servicetypes.PAUSED
	case servicetypes.COMPLETED:
		ctx.Logger().Error(
			"Feed state invalid", "requestContextID",
			requestContextID.String(), "state", reqCtx.State.String(),
		)
		return
	}

	ctx.Logger().Info(
		"Feed state changed",
		"feed", feed.FeedName,
		"srcState", oldState,
		"dstState", reqCtx.State.String(),
	)
	k.dequeueAndEnqueue(ctx, feed.FeedName, oldState, reqCtx.State)
	return
}

func (k Keeper) GetRequestContext(ctx sdk.Context, requestContextID tmbytes.HexBytes) (servicetypes.RequestContext, bool) {
	return k.sk.GetRequestContext(ctx, requestContextID)
}
