package keeper

import (
	"context"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/htlc/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) HTLC(c context.Context, request *types.QueryHTLCRequest) (*types.QueryHTLCResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	hashLock, err := hex.DecodeString(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid hash lock %s", request.Id)
	}

	htlc, found := k.GetHTLC(ctx, hashLock)
	if !found {
		return nil, status.Errorf(codes.NotFound, "HTLC %s not found", request.Id)
	}

	return &types.QueryHTLCResponse{Htlc: &htlc}, nil
}

func (k Keeper) AssetSupply(context.Context, *types.QueryAssetSupplyRequest) (*types.QueryAssetSupplyResponse, error) {
	// TODO
	return nil, nil
}

func (k Keeper) AssetSupplies(context.Context, *types.QueryAssetSuppliesRequest) (*types.QueryAssetSuppliesResponse, error) {
	// TODO
	return nil, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
