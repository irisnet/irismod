package coinswap

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	"mods.irisnet.org/e2e"
	coinswaptypes "mods.irisnet.org/modules/coinswap/types"
	tokenv1 "mods.irisnet.org/modules/token/types/v1"
	"mods.irisnet.org/simapp"
)

// QueryTestSuite is a suite of end-to-end tests for the nft module
type QueryTestSuite struct {
	e2e.TestSuite
}

// SetupSuite creates a new network for integration tests
func (s *QueryTestSuite) SetupSuite() {
	sdk.SetCoinDenomRegex(func() string {
		return `[a-zA-Z][a-zA-Z0-9/\-]{2,127}`
	})
	s.SetupSuiteWithModifyConfigFn(func(cfg *network.Config) {
		e2e.AddTestTokenToGenesis(s.T(), cfg, tokenv1.Token{
			Symbol:        "kitty",
			Name:          "Kitty Token",
			Scale:         0,
			MinUnit:       "kitty",
			InitialSupply: 100000000,
			MaxSupply:     200000000,
			Mintable:      true,
		})
	})
}

// TestCoinswap tests all query command in the nft module
func (s *QueryTestSuite) TestCoinswap() {
	val := s.Validators[0]
	clientCtx := val.ClientCtx
	// ---------------------------------------------------------------------------

	from := val.Address
	symbol := "kitty"
	minUnit := "kitty"
	initialSupply := uint64(100000000)
	baseURL := val.APIAddress
	lptDenom := "lpt-1"

	// Fund the owner from the genesis token without using the disabled issue path.
	msgMintToken := &tokenv1.MsgMintToken{
		Coin:     sdk.NewCoin(minUnit, math.NewIntFromUint64(initialSupply)),
		Receiver: from.String(),
		Owner:    from.String(),
	}
	txResult := s.BlockSendMsgs(s.T(), msgMintToken)
	s.Require().Equal(uint32(0), txResult.Code, "send mint token msg failed")

	// _ = tokentestutil.IssueTokenExec(s.T(), s.Network, clientCtx, from.String(), args...)

	balances := simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("100000000", balances.AmountOf(symbol).String())
	s.Require().Equal("399998689", balances.AmountOf(sdk.DefaultBondDenom).String())

	// test add liquidity (poor not exist)
	status, err := clientCtx.Client.Status(context.Background())
	s.Require().NoError(err)
	deadline := status.SyncInfo.LatestBlockTime.Add(time.Minute)

	msgAddLiquidity := &coinswaptypes.MsgAddLiquidity{
		MaxToken:         sdk.NewCoin(symbol, math.NewInt(1000)),
		ExactStandardAmt: math.NewInt(1000),
		MinLiquidity:     math.NewInt(1000),
		Deadline:         deadline.Unix(),
		Sender:           from.String(),
	}
	s.SendMsgs(s.T(), msgAddLiquidity)

	balances = simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("99999000", balances.AmountOf(symbol).String())
	s.Require().Equal("399992679", balances.AmountOf(sdk.DefaultBondDenom).String())
	s.Require().Equal("1000", balances.AmountOf(lptDenom).String())

	queryPoolResponse := proto.Message(&coinswaptypes.QueryLiquidityPoolResponse{})
	url := fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
	resp, err := testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

	queryPool := queryPoolResponse.(*coinswaptypes.QueryLiquidityPoolResponse)
	s.Require().Equal("1000", queryPool.Pool.Standard.Amount.String())
	s.Require().Equal("1000", queryPool.Pool.Token.Amount.String())
	s.Require().Equal("1000", queryPool.Pool.Lpt.Amount.String())

	// test add liquidity (poor exist)
	status, err = clientCtx.Client.Status(context.Background())
	s.Require().NoError(err)
	deadline = status.SyncInfo.LatestBlockTime.Add(time.Minute)

	msgAddLiquidity = &coinswaptypes.MsgAddLiquidity{
		MaxToken:         sdk.NewCoin(symbol, math.NewInt(2001)),
		ExactStandardAmt: math.NewInt(2000),
		MinLiquidity:     math.NewInt(2000),
		Deadline:         deadline.Unix(),
		Sender:           from.String(),
	}
	s.SendMsgs(s.T(), msgAddLiquidity)

	balances = simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("99996999", balances.AmountOf(symbol).String())
	s.Require().Equal("399990669", balances.AmountOf(sdk.DefaultBondDenom).String())
	s.Require().Equal("3000", balances.AmountOf(lptDenom).String())

	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
	resp, err = testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

	s.Require().Equal("3000", queryPool.Pool.Standard.Amount.String())
	s.Require().Equal("3001", queryPool.Pool.Token.Amount.String())
	s.Require().Equal("3000", queryPool.Pool.Lpt.Amount.String())

	// test sell order
	msgSellOrder := &coinswaptypes.MsgSwapOrder{
		Input: coinswaptypes.Input{
			Address: from.String(),
			Coin:    sdk.NewCoin(symbol, math.NewInt(1000)),
		},
		Output: coinswaptypes.Output{
			Address: from.String(),
			Coin:    sdk.NewInt64Coin(s.BondDenom, 748),
		},
		Deadline:   deadline.Unix(),
		IsBuyOrder: false,
	}
	s.SendMsgs(s.T(), msgSellOrder)

	balances = simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("99995999", balances.AmountOf(symbol).String())
	s.Require().Equal("399991407", balances.AmountOf(sdk.DefaultBondDenom).String())
	s.Require().Equal("3000", balances.AmountOf(lptDenom).String())

	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
	resp, err = testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

	s.Require().Equal("2252", queryPool.Pool.Standard.Amount.String())
	s.Require().Equal("4001", queryPool.Pool.Token.Amount.String())
	s.Require().Equal("3000", queryPool.Pool.Lpt.Amount.String())

	// test buy order
	msgBuyOrder := &coinswaptypes.MsgSwapOrder{
		Input: coinswaptypes.Input{
			Address: from.String(),
			Coin:    sdk.NewInt64Coin(s.BondDenom, 753),
		},
		Output: coinswaptypes.Output{
			Address: from.String(),
			Coin:    sdk.NewCoin(symbol, math.NewInt(1000)),
		},
		Deadline:   deadline.Unix(),
		IsBuyOrder: true,
	}
	s.SendMsgs(s.T(), msgBuyOrder)

	balances = simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("99996999", balances.AmountOf(symbol).String())
	s.Require().Equal("399990644", balances.AmountOf(sdk.DefaultBondDenom).String())
	s.Require().Equal("3000", balances.AmountOf(lptDenom).String())

	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
	resp, err = testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

	s.Require().Equal("3005", queryPool.Pool.Standard.Amount.String())
	s.Require().Equal("3001", queryPool.Pool.Token.Amount.String())
	s.Require().Equal("3000", queryPool.Pool.Lpt.Amount.String())

	// Test remove liquidity (remove part)
	msgRemoveLiquidity := &coinswaptypes.MsgRemoveLiquidity{
		WithdrawLiquidity: sdk.NewCoin(lptDenom, math.NewInt(2000)),
		MinToken:          math.NewInt(2000),
		MinStandardAmt:    math.NewInt(2000),
		Deadline:          deadline.Unix(),
		Sender:            from.String(),
	}

	// prepare txBuilder with msg
	s.SendMsgs(s.T(), msgRemoveLiquidity)

	balances = simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("99998999", balances.AmountOf(symbol).String())
	s.Require().Equal("399992637", balances.AmountOf(sdk.DefaultBondDenom).String())
	s.Require().Equal("1000", balances.AmountOf(lptDenom).String())

	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
	resp, err = testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

	s.Require().Equal("1002", queryPool.Pool.Standard.Amount.String())
	s.Require().Equal("1001", queryPool.Pool.Token.Amount.String())
	s.Require().Equal("1000", queryPool.Pool.Lpt.Amount.String())

	// Test remove liquidity (remove all)
	msgRemoveLiquidity = &coinswaptypes.MsgRemoveLiquidity{
		WithdrawLiquidity: sdk.NewCoin(lptDenom, math.NewInt(1000)),
		MinToken:          math.NewInt(1000),
		MinStandardAmt:    math.NewInt(1000),
		Deadline:          deadline.Unix(),
		Sender:            from.String(),
	}

	// prepare txBuilder with msg
	s.SendMsgs(s.T(), msgRemoveLiquidity)

	balances = simapp.QueryBalancesExec(s.T(), clientCtx, from.String())
	s.Require().Equal("100000000", balances.AmountOf(symbol).String())
	s.Require().Equal("399993629", balances.AmountOf(sdk.DefaultBondDenom).String())
	s.Require().Equal("0", balances.AmountOf(lptDenom).String())

	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
	resp, err = testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

	s.Require().Equal("0", queryPool.Pool.Standard.Amount.String())
	s.Require().Equal("0", queryPool.Pool.Token.Amount.String())
	s.Require().Equal("0", queryPool.Pool.Lpt.Amount.String())

	queryPoolsResponse := proto.Message(&coinswaptypes.QueryLiquidityPoolsResponse{})
	url = fmt.Sprintf("%s/irismod/coinswap/pools", baseURL)
	resp, err = testutil.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolsResponse))

	queryPools := queryPoolsResponse.(*coinswaptypes.QueryLiquidityPoolsResponse)
	s.Require().Len(queryPools.Pools, 1)
}
