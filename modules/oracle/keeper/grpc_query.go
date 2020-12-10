package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	gogotypes "github.com/gogo/protobuf/types"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/oracle/types"
)

var _ types.QueryServer = Keeper{}

// Feed queries a feed by feed name
func (k Keeper) Feed(c context.Context, req *types.QueryFeedRequest) (*types.QueryFeedResponse, error) {
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

func (k Keeper) Feeds(c context.Context, req *types.QueryFeedsRequest) (*types.QueryFeedsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	state := strings.TrimSpace(req.State)
	var result types.FeedsContext
	var pageRes *query.PageResponse
	var err error
	store := ctx.KVStore(k.storeKey)
	if len(state) == 0 {
		feedStore := prefix.NewStore(store, types.GetFeedPrefixKey())
		pageRes, err = query.Paginate(feedStore, req.Pagination, func(key []byte, value []byte) error {
			var feed types.Feed
			err := k.cdc.UnmarshalBinaryBare(value, &feed)
			if err != nil {
				return err
			}
			result = append(result, BuildFeedContext(ctx, k, feed))
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}
	} else {
		state, err := types.RequestContextStateFromString(req.State)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request state")

		}
		feedStore := prefix.NewStore(store, types.GetFeedStatePrefixKey(state))
		pageRes, err = query.Paginate(feedStore, req.Pagination, func(key []byte, value []byte) error {
			var feedName gogotypes.StringValue
			err := k.cdc.UnmarshalBinaryBare(value, &feedName)
			if err != nil {
				return err
			}
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

func (k Keeper) FeedValue(c context.Context, req *types.QueryFeedValueRequest) (*types.QueryFeedValueResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	result := k.GetFeedValues(ctx, req.FeedName)
	return &types.QueryFeedValueResponse{FeedValues: result}, nil
}
