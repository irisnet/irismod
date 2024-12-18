package keeper

import (
	"context"
	"encoding/hex"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	gogotypes "github.com/cosmos/gogoproto/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mods.irisnet.org/modules/oracle/types"
)

var _ types.QueryServer = Keeper{}

// Feed queries a feed by the feed name
func (k Keeper) Feed(
	c context.Context,
	req *types.QueryFeedRequest,
) (*types.QueryFeedResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	feed, found := k.GetFeed(ctx, req.FeedName)
	if !found {
		return nil, status.Errorf(codes.NotFound, "feed %s not found", req.FeedName)
	}
	feedCtx := BuildFeedContext(ctx, k, feed)
	return &types.QueryFeedResponse{Feed: feedCtx}, nil
}

func (k Keeper) Feeds(
	c context.Context,
	req *types.QueryFeedsRequest,
) (*types.QueryFeedsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var result types.FeedsContext
	var pageRes *query.PageResponse
	var err error
	store := ctx.KVStore(k.storeKey)
	if len(req.State) == 0 {
		feedStore := prefix.NewStore(store, types.GetFeedPrefixKey())
		pageRes, err = query.Paginate(
			feedStore,
			shapePageRequest(req.Pagination),
			func(key, value []byte) error {
				var feed types.Feed
				k.cdc.MustUnmarshal(value, &feed)
				result = append(result, BuildFeedContext(ctx, k, feed))
				return nil
			},
		)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}
	} else {
		state, err := types.RequestContextStateFromString(req.State)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request state")
		}
		feedStore := prefix.NewStore(store, types.GetFeedStatePrefixKey(state))
		pageRes, err = query.Paginate(feedStore, shapePageRequest(req.Pagination), func(key, value []byte) error {
			var feedName gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &feedName)
			if feed, found := k.GetFeed(ctx, feedName.Value); found {
				result = append(result, BuildFeedContext(ctx, k, feed))
			}
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}
	}
	return &types.QueryFeedsResponse{Feeds: result, Pagination: pageRes}, nil
}

func (k Keeper) FeedValue(
	c context.Context,
	req *types.QueryFeedValueRequest,
) (*types.QueryFeedValueResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	result := k.GetFeedValues(ctx, req.FeedName)
	return &types.QueryFeedValueResponse{FeedValues: result}, nil
}

func BuildFeedContext(ctx sdk.Context, k Keeper, feed types.Feed) (feedCtx types.FeedContext) {
	requestContextID, _ := hex.DecodeString(feed.RequestContextID)
	reqCtx, found := k.sk.GetRequestContext(ctx, requestContextID)
	if found {
		feedCtx.Providers = reqCtx.Providers
		feedCtx.ResponseThreshold = reqCtx.ResponseThreshold
		feedCtx.ServiceName = reqCtx.ServiceName
		feedCtx.Input = reqCtx.Input
		feedCtx.RepeatedFrequency = reqCtx.RepeatedFrequency
		feedCtx.ServiceFeeCap = reqCtx.ServiceFeeCap
		feedCtx.Timeout = reqCtx.Timeout
		feedCtx.State = reqCtx.State
	}
	feedCtx.Feed = &feed
	return feedCtx
}
