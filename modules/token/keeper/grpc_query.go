package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/irisnet/irismod/modules/token/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Token(c context.Context, req *types.QueryTokenRequest) (*types.QueryTokenResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	token, err := k.GetToken(ctx, strings.ToLower(req.Denom))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "token %s not found", req.Denom)
	}
	msg, ok := token.(proto.Message)
	if !ok {
		return nil, status.Errorf(codes.Internal, "can't protomarshal %T", token)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTokenResponse{Token: any}, nil
}

func (k Keeper) Tokens(c context.Context, req *types.QueryTokensRequest) (*types.QueryTokensResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var owner sdk.AccAddress
	var err error
	if len(req.Owner) > 0 {
		owner, err = sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid owner address (%s)", err))
		}
	}

	var tokens []types.TokenI
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)
	if owner == nil {
		tokenStore := prefix.NewStore(store, types.PrefixTokenForSymbol)
		pageRes, err = query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
			var token types.Token
			err := k.cdc.UnmarshalBinaryBare(value, &token)
			if err != nil {
				return err
			}
			tokens = append(tokens, &token)
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}
	} else {
		tokenStore := prefix.NewStore(store, types.KeyTokens(owner, ""))
		pageRes, err = query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
			var symbol gogotypes.StringValue
			err := k.cdc.UnmarshalBinaryBare(value, &symbol)
			if err != nil {
				return err
			}
			token, err := k.GetToken(ctx, symbol.Value)
			if err == nil {
				tokens = append(tokens, token)
			}
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}
	}
	result := make([]*codectypes.Any, len(tokens))
	for i, token := range tokens {
		msg, ok := token.(proto.Message)
		if !ok {
			return nil, status.Errorf(codes.Internal, "%T does not implement proto.Message", token)
		}

		var err error
		if result[i], err = codectypes.NewAnyWithValue(msg); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryTokensResponse{Tokens: result, Pagination: pageRes}, nil
}

func (k Keeper) Fees(c context.Context, req *types.QueryFeesRequest) (*types.QueryFeesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := types.CheckSymbol(req.Symbol); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	symbol := strings.ToLower(req.Symbol)
	issueFee := k.GetTokenIssueFee(ctx, symbol)
	mintFee := k.GetTokenMintFee(ctx, symbol)

	return &types.QueryFeesResponse{
		Exist:    k.HasToken(ctx, symbol),
		IssueFee: issueFee,
		MintFee:  mintFee,
	}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParamSet(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
