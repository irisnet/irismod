package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/irisnet/irismod/modules/token/types/v1"
)

// queryTokenFees retrieves the fees of issuance and minting for the specified symbol
func queryTokenFees(cliCtx client.Context, symbol string) (v1.QueryFeesResponse, error) {
	queryClient := v1.NewQueryClient(cliCtx)

	resp, err := queryClient.Fees(context.Background(), &v1.QueryFeesRequest{Symbol: symbol})
	if err != nil {
		return v1.QueryFeesResponse{}, err
	}

	return *resp, err
}

// queryToken query token information
func queryToken(cliCtx client.Context, denom string) (v1.TokenI, error) {
	queryClient := v1.NewQueryClient(cliCtx)

	resp, err := queryClient.Token(context.Background(), &v1.QueryTokenRequest{
		Denom: denom,
	})
	if err != nil {
		return nil, err
	}

	var evi v1.TokenI
	err = cliCtx.InterfaceRegistry.UnpackAny(resp.Token, &evi)
	if err != nil {
		return nil, err
	}

	return evi, err
}

func parseCoin(cliCtx client.Context, denom string) (sdk.Coin, v1.TokenI, error) {
	decCoin, err := sdk.ParseDecCoin(denom)
	if err != nil {
		return sdk.Coin{}, nil, err
	}

	token, err := queryToken(cliCtx, decCoin.Denom)
	if err != nil {
		return sdk.Coin{}, nil, err
	}

	coin, err := token.ToMinCoin(decCoin)
	if err != nil {
		return sdk.Coin{}, nil, err
	}
	return coin, token, err
}
