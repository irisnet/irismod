package keeper

import (
	"context"

	etherminttypes "github.com/evmos/ethermint/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/irisnet/irismod/modules/erc721-converter/types"
)

// TokenPairs returns all registered pairs
func (k Keeper) TokenPairs(c context.Context, req *types.QueryTokenPairsRequest) (*types.QueryTokenPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var pairs []types.TokenPair
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)

	pageRes, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var pair types.TokenPair
		if err := k.cdc.Unmarshal(value, &pair); err != nil {
			return err
		}
		pairs = append(pairs, pair)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryTokenPairsResponse{
		TokenPairs: pairs,
		Pagination: pageRes,
	}, nil
}

// TokenPair returns a given registered token pair
func (k Keeper) TokenPair(c context.Context, req *types.QueryTokenPairRequest) (*types.QueryTokenPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// check if the token is a hex address, if not, check if it is a valid SDK
	// denom
	if err := etherminttypes.ValidateAddress(req.Token); err != nil {
		if err := sdk.ValidateDenom(req.Token); err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"invalid format for token %s, should be either hex ('0x...') cosmos denom", req.Token,
			)
		}
	}

	id := k.GetTokenPairID(ctx, req.Token)

	if len(id) == 0 {
		return nil, status.Errorf(codes.NotFound, "token pair with token '%s'", req.Token)
	}

	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "token pair with token '%s'", req.Token)
	}

	return &types.QueryTokenPairResponse{TokenPair: pair}, nil
}
