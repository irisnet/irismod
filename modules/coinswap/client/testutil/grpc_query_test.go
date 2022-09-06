package testutil_test

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/gogo/protobuf/proto"
// 	"github.com/stretchr/testify/suite"

// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
// 	codectype "github.com/cosmos/cosmos-sdk/codec/types"
// 	"github.com/cosmos/cosmos-sdk/testutil/network"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/rest"
// 	"github.com/cosmos/cosmos-sdk/types/tx"
// 	"github.com/cosmos/cosmos-sdk/types/tx/signing"
// 	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

// 	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"
// 	tokencli "github.com/irisnet/irismod/modules/token/client/cli"
// 	tokentestutil "github.com/irisnet/irismod/modules/token/client/testutil"
// 	"github.com/irisnet/irismod/simapp"
// )

// type IntegrationTestSuite struct {
// 	suite.Suite

// 	cfg     network.Config
// 	network *network.Network
// }

// func (s *IntegrationTestSuite) SetupSuite() {
// 	s.T().Log("setting up integration test suite")

// 	cfg := simapp.NewConfig()
// 	cfg.NumValidators = 1

// 	s.cfg = cfg
// 	s.network = network.New(s.T(), cfg)

// 	_, err := s.network.WaitForHeight(1)
// 	s.Require().NoError(err)

// 	sdk.SetCoinDenomRegex(func() string {
// 		return `[a-zA-Z][a-zA-Z0-9/\-]{2,127}`
// 	})
// }

// func (s *IntegrationTestSuite) TearDownSuite() {
// 	s.T().Log("tearing down integration test suite")
// 	s.network.Cleanup()
// }

// func TestIntegrationTestSuite(t *testing.T) {
// 	suite.Run(t, new(IntegrationTestSuite))
// }

// func (s *IntegrationTestSuite) TestCoinswap() {
// 	val := s.network.Validators[0]
// 	clientCtx := val.ClientCtx
// 	// ---------------------------------------------------------------------------

// 	from := val.Address
// 	symbol := "kitty"
// 	name := "Kitty Token"
// 	minUnit := "kitty"
// 	scale := 0
// 	initialSupply := int64(100000000)
// 	maxSupply := int64(200000000)
// 	mintable := true
// 	baseURL := val.APIAddress
// 	lptDenom := "lpt-1"

// 	//------test GetCmdIssueToken()-------------
// 	args := []string{
// 		fmt.Sprintf("--%s=%s", tokencli.FlagSymbol, symbol),
// 		fmt.Sprintf("--%s=%s", tokencli.FlagName, name),
// 		fmt.Sprintf("--%s=%s", tokencli.FlagMinUnit, minUnit),
// 		fmt.Sprintf("--%s=%d", tokencli.FlagScale, scale),
// 		fmt.Sprintf("--%s=%d", tokencli.FlagInitialSupply, initialSupply),
// 		fmt.Sprintf("--%s=%d", tokencli.FlagMaxSupply, maxSupply),
// 		fmt.Sprintf("--%s=%t", tokencli.FlagMintable, mintable),

// 		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
// 		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
// 		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
// 	}
// 	respType := proto.Message(&sdk.TxResponse{})
// 	expectedCode := uint32(0)
// 	bz, err := tokentestutil.IssueTokenExec(clientCtx, from.String(), args...)

// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
// 	txResp := respType.(*sdk.TxResponse)
// 	s.Require().Equal(expectedCode, txResp.Code)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err := simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))
// 	balances := respType.(*banktypes.QueryAllBalancesResponse)
// 	fmt.Println(balances.Balances)
// 	s.Require().Equal("100000000", balances.Balances.AmountOf(symbol).String())
// 	s.Require().Equal("399986975", balances.Balances.AmountOf(sdk.DefaultBondDenom).String())

// 	var account authtypes.AccountI
// 	respType = proto.Message(&codectype.Any{})
// 	out, err = simapp.QueryAccountExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))
// 	err = clientCtx.InterfaceRegistry.UnpackAny(respType.(*codectype.Any), &account)
// 	s.Require().NoError(err)

// 	// test add liquidity (poor not exist)
// 	status, err := clientCtx.Client.Status(context.Background())
// 	s.Require().NoError(err)
// 	deadline := status.SyncInfo.LatestBlockTime.Add(time.Minute)

// 	msgAddLiquidity := &coinswaptypes.MsgAddLiquidity{
// 		MaxToken:         sdk.NewCoin(symbol, sdk.NewInt(1000)),
// 		ExactStandardAmt: sdk.NewInt(1000),
// 		MinLiquidity:     sdk.NewInt(1000),
// 		Deadline:         deadline.Unix(),
// 		Sender:           from.String(),
// 	}

// 	// prepare txBuilder with msg
// 	txBuilder := val.ClientCtx.TxConfig.NewTxBuilder()
// 	feeAmount := sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
// 	err = txBuilder.SetMsgs(msgAddLiquidity)
// 	s.Require().NoError(err)
// 	txBuilder.SetFeeAmount(feeAmount)
// 	txBuilder.SetGasLimit(1000000)

// 	// setup txFactory
// 	txFactory := clienttx.Factory{}.
// 		WithChainID(val.ClientCtx.ChainID).
// 		WithKeybase(val.ClientCtx.Keyring).
// 		WithTxConfig(val.ClientCtx.TxConfig).
// 		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
// 		WithSequence(account.GetSequence())

// 	// Sign Tx.
// 	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
// 	s.Require().NoError(err)

// 	txBytes, err := val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
// 	s.Require().NoError(err)
// 	req := &tx.BroadcastTxRequest{
// 		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
// 		TxBytes: txBytes,
// 	}

// 	reqBz, err := val.ClientCtx.Codec.MarshalJSON(req)
// 	s.Require().NoError(err)
// 	res, err := rest.PostRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL), "application/json", reqBz)
// 	s.Require().NoError(err)
// 	var result tx.BroadcastTxResponse
// 	err = val.ClientCtx.Codec.UnmarshalJSON(res, &result)
// 	s.Require().NoError(err)
// 	s.Require().Equal(uint32(0), result.TxResponse.Code, "rawlog", result.TxResponse.RawLog)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))

// 	balances = respType.(*banktypes.QueryAllBalancesResponse)
// 	coins := balances.Balances
// 	fmt.Println(coins)
// 	s.Require().Equal("99999000", coins.AmountOf(symbol).String())
// 	s.Require().Equal("399980965", coins.AmountOf(sdk.DefaultBondDenom).String())
// 	s.Require().Equal("1000", coins.AmountOf(lptDenom).String())

// 	queryPoolResponse := proto.Message(&coinswaptypes.QueryLiquidityPoolResponse{})
// 	url := fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
// 	resp, err := rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

// 	queryPool := queryPoolResponse.(*coinswaptypes.QueryLiquidityPoolResponse)
// 	s.Require().Equal("1000", queryPool.Pool.Standard.Amount.String())
// 	s.Require().Equal("1000", queryPool.Pool.Token.Amount.String())
// 	s.Require().Equal("1000", queryPool.Pool.Lpt.Amount.String())

// 	// test add liquidity (poor exist)
// 	status, err = clientCtx.Client.Status(context.Background())
// 	s.Require().NoError(err)
// 	deadline = status.SyncInfo.LatestBlockTime.Add(time.Minute)

// 	msgAddLiquidity = &coinswaptypes.MsgAddLiquidity{
// 		MaxToken:         sdk.NewCoin(symbol, sdk.NewInt(2001)),
// 		ExactStandardAmt: sdk.NewInt(2000),
// 		MinLiquidity:     sdk.NewInt(2000),
// 		Deadline:         deadline.Unix(),
// 		Sender:           from.String(),
// 	}

// 	// prepare txBuilder with msg
// 	txBuilder = val.ClientCtx.TxConfig.NewTxBuilder()
// 	feeAmount = sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
// 	err = txBuilder.SetMsgs(msgAddLiquidity)
// 	s.Require().NoError(err)
// 	txBuilder.SetFeeAmount(feeAmount)
// 	txBuilder.SetGasLimit(1000000)

// 	// setup txFactory
// 	txFactory = clienttx.Factory{}.
// 		WithChainID(val.ClientCtx.ChainID).
// 		WithKeybase(val.ClientCtx.Keyring).
// 		WithTxConfig(val.ClientCtx.TxConfig).
// 		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
// 		WithSequence(account.GetSequence() + 1)

// 	// sign Tx
// 	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
// 	s.Require().NoError(err)

// 	txBytes, err = val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
// 	s.Require().NoError(err)
// 	req = &tx.BroadcastTxRequest{
// 		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
// 		TxBytes: txBytes,
// 	}

// 	reqBz, err = val.ClientCtx.Codec.MarshalJSON(req)
// 	s.Require().NoError(err)
// 	res, err = rest.PostRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL), "application/json", reqBz)
// 	s.Require().NoError(err)
// 	err = val.ClientCtx.Codec.UnmarshalJSON(res, &result)
// 	s.Require().NoError(err)
// 	s.Require().Equal(uint32(0), result.TxResponse.Code, "rawlog", result.TxResponse.RawLog)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))

// 	balances = respType.(*banktypes.QueryAllBalancesResponse)
// 	coins = balances.Balances
// 	fmt.Println(coins)
// 	s.Require().Equal("99996999", coins.AmountOf(symbol).String())
// 	s.Require().Equal("399978955", coins.AmountOf(sdk.DefaultBondDenom).String())
// 	s.Require().Equal("3000", coins.AmountOf(lptDenom).String())

// 	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
// 	resp, err = rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

// 	s.Require().Equal("3000", queryPool.Pool.Standard.Amount.String())
// 	s.Require().Equal("3001", queryPool.Pool.Token.Amount.String())
// 	s.Require().Equal("3000", queryPool.Pool.Lpt.Amount.String())

// 	// test sell order
// 	msgSellOrder := &coinswaptypes.MsgSwapOrder{
// 		Input: coinswaptypes.Input{
// 			Address: from.String(),
// 			Coin:    sdk.NewCoin(symbol, sdk.NewInt(1000)),
// 		},
// 		Output: coinswaptypes.Output{
// 			Address: from.String(),
// 			Coin:    sdk.NewInt64Coin(s.cfg.BondDenom, 748),
// 		},
// 		Deadline:   deadline.Unix(),
// 		IsBuyOrder: false,
// 	}

// 	// prepare txBuilder with msg
// 	txBuilder = val.ClientCtx.TxConfig.NewTxBuilder()
// 	feeAmount = sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
// 	err = txBuilder.SetMsgs(msgSellOrder)
// 	s.Require().NoError(err)
// 	txBuilder.SetFeeAmount(feeAmount)
// 	txBuilder.SetGasLimit(1000000)

// 	// setup txFactory
// 	txFactory = clienttx.Factory{}.
// 		WithChainID(val.ClientCtx.ChainID).
// 		WithKeybase(val.ClientCtx.Keyring).
// 		WithTxConfig(val.ClientCtx.TxConfig).
// 		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
// 		WithSequence(account.GetSequence() + 2)

// 	// sign Tx (offline mode so we can manually set sequence number)
// 	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
// 	s.Require().NoError(err)

// 	txBytes, err = val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
// 	s.Require().NoError(err)
// 	req = &tx.BroadcastTxRequest{
// 		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
// 		TxBytes: txBytes,
// 	}

// 	reqBz, err = val.ClientCtx.Codec.MarshalJSON(req)
// 	s.Require().NoError(err)
// 	_, err = rest.PostRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL), "application/json", reqBz)
// 	s.Require().NoError(err)
// 	err = val.ClientCtx.Codec.UnmarshalJSON(res, &result)
// 	s.Require().NoError(err)
// 	s.Require().Equal(uint32(0), result.TxResponse.Code, "rawlog", result.TxResponse.RawLog)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))

// 	balances = respType.(*banktypes.QueryAllBalancesResponse)
// 	coins = balances.Balances
// 	fmt.Println(coins)
// 	s.Require().Equal("99995999", coins.AmountOf(symbol).String())
// 	s.Require().Equal("399979693", coins.AmountOf(sdk.DefaultBondDenom).String())
// 	s.Require().Equal("3000", coins.AmountOf(lptDenom).String())

// 	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
// 	resp, err = rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

// 	s.Require().Equal("2252", queryPool.Pool.Standard.Amount.String())
// 	s.Require().Equal("4001", queryPool.Pool.Token.Amount.String())
// 	s.Require().Equal("3000", queryPool.Pool.Lpt.Amount.String())

// 	// test buy order
// 	msgBuyOrder := &coinswaptypes.MsgSwapOrder{
// 		Input: coinswaptypes.Input{
// 			Address: from.String(),
// 			Coin:    sdk.NewInt64Coin(s.cfg.BondDenom, 753),
// 		},
// 		Output: coinswaptypes.Output{
// 			Address: from.String(),
// 			Coin:    sdk.NewCoin(symbol, sdk.NewInt(1000)),
// 		},
// 		Deadline:   deadline.Unix(),
// 		IsBuyOrder: true,
// 	}

// 	// prepare txBuilder with msg
// 	txBuilder = val.ClientCtx.TxConfig.NewTxBuilder()
// 	feeAmount = sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
// 	err = txBuilder.SetMsgs(msgBuyOrder)
// 	s.Require().NoError(err)
// 	txBuilder.SetFeeAmount(feeAmount)
// 	txBuilder.SetGasLimit(1000000)

// 	// setup txFactory
// 	txFactory = clienttx.Factory{}.
// 		WithChainID(val.ClientCtx.ChainID).
// 		WithKeybase(val.ClientCtx.Keyring).
// 		WithTxConfig(val.ClientCtx.TxConfig).
// 		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
// 		WithSequence(account.GetSequence() + 3)

// 	// sign Tx (offline mode so we can manually set sequence number)
// 	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
// 	s.Require().NoError(err)

// 	txBytes, err = val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
// 	s.Require().NoError(err)
// 	req = &tx.BroadcastTxRequest{
// 		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
// 		TxBytes: txBytes,
// 	}

// 	reqBz, err = val.ClientCtx.Codec.MarshalJSON(req)
// 	s.Require().NoError(err)
// 	_, err = rest.PostRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL), "application/json", reqBz)
// 	s.Require().NoError(err)
// 	err = val.ClientCtx.Codec.UnmarshalJSON(res, &result)
// 	s.Require().NoError(err)
// 	s.Require().Equal(uint32(0), result.TxResponse.Code, "rawlog", result.TxResponse.RawLog)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))

// 	balances = respType.(*banktypes.QueryAllBalancesResponse)
// 	coins = balances.Balances
// 	fmt.Println(coins)
// 	s.Require().Equal("99996999", coins.AmountOf(symbol).String())
// 	s.Require().Equal("399978930", coins.AmountOf(sdk.DefaultBondDenom).String())
// 	s.Require().Equal("3000", coins.AmountOf(lptDenom).String())

// 	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
// 	resp, err = rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

// 	s.Require().Equal("3005", queryPool.Pool.Standard.Amount.String())
// 	s.Require().Equal("3001", queryPool.Pool.Token.Amount.String())
// 	s.Require().Equal("3000", queryPool.Pool.Lpt.Amount.String())

// 	// Test remove liquidity (remove part)
// 	msgRemoveLiquidity := &coinswaptypes.MsgRemoveLiquidity{
// 		WithdrawLiquidity: sdk.NewCoin(lptDenom, sdk.NewInt(2000)),
// 		MinToken:          sdk.NewInt(2000),
// 		MinStandardAmt:    sdk.NewInt(2000),
// 		Deadline:          deadline.Unix(),
// 		Sender:            from.String(),
// 	}

// 	// prepare txBuilder with msg
// 	txBuilder = val.ClientCtx.TxConfig.NewTxBuilder()
// 	feeAmount = sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
// 	err = txBuilder.SetMsgs(msgRemoveLiquidity)
// 	s.Require().NoError(err)
// 	txBuilder.SetFeeAmount(feeAmount)
// 	txBuilder.SetGasLimit(1000000)

// 	// setup txFactory
// 	txFactory = clienttx.Factory{}.
// 		WithChainID(val.ClientCtx.ChainID).
// 		WithKeybase(val.ClientCtx.Keyring).
// 		WithTxConfig(val.ClientCtx.TxConfig).
// 		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
// 		WithSequence(account.GetSequence() + 4)

// 	// sign Tx (offline mode so we can manually set sequence number)
// 	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
// 	s.Require().NoError(err)

// 	txBytes, err = val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
// 	s.Require().NoError(err)
// 	req = &tx.BroadcastTxRequest{
// 		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
// 		TxBytes: txBytes,
// 	}

// 	reqBz, err = val.ClientCtx.Codec.MarshalJSON(req)
// 	s.Require().NoError(err)
// 	_, err = rest.PostRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL), "application/json", reqBz)
// 	s.Require().NoError(err)
// 	err = val.ClientCtx.Codec.UnmarshalJSON(res, &result)
// 	s.Require().NoError(err)
// 	s.Require().Equal(uint32(0), result.TxResponse.Code, "rawlog", result.TxResponse.RawLog)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))

// 	balances = respType.(*banktypes.QueryAllBalancesResponse)
// 	coins = balances.Balances
// 	fmt.Println(coins)
// 	s.Require().Equal("99998999", coins.AmountOf(symbol).String())
// 	s.Require().Equal("399980923", coins.AmountOf(sdk.DefaultBondDenom).String())
// 	s.Require().Equal("1000", coins.AmountOf(lptDenom).String())

// 	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
// 	resp, err = rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

// 	s.Require().Equal("1002", queryPool.Pool.Standard.Amount.String())
// 	s.Require().Equal("1001", queryPool.Pool.Token.Amount.String())
// 	s.Require().Equal("1000", queryPool.Pool.Lpt.Amount.String())

// 	// Test remove liquidity (remove all)
// 	msgRemoveLiquidity = &coinswaptypes.MsgRemoveLiquidity{
// 		WithdrawLiquidity: sdk.NewCoin(lptDenom, sdk.NewInt(1000)),
// 		MinToken:          sdk.NewInt(1000),
// 		MinStandardAmt:    sdk.NewInt(1000),
// 		Deadline:          deadline.Unix(),
// 		Sender:            from.String(),
// 	}

// 	// prepare txBuilder with msg
// 	txBuilder = val.ClientCtx.TxConfig.NewTxBuilder()
// 	feeAmount = sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
// 	err = txBuilder.SetMsgs(msgRemoveLiquidity)
// 	s.Require().NoError(err)
// 	txBuilder.SetFeeAmount(feeAmount)
// 	txBuilder.SetGasLimit(1000000)

// 	// setup txFactory
// 	txFactory = clienttx.Factory{}.
// 		WithChainID(val.ClientCtx.ChainID).
// 		WithKeybase(val.ClientCtx.Keyring).
// 		WithTxConfig(val.ClientCtx.TxConfig).
// 		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
// 		WithSequence(account.GetSequence() + 5)

// 	// sign Tx (offline mode so we can manually set sequence number)
// 	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, false, true)
// 	s.Require().NoError(err)

// 	txBytes, err = val.ClientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
// 	s.Require().NoError(err)
// 	req = &tx.BroadcastTxRequest{
// 		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
// 		TxBytes: txBytes,
// 	}

// 	reqBz, err = val.ClientCtx.Codec.MarshalJSON(req)
// 	s.Require().NoError(err)
// 	_, err = rest.PostRequest(fmt.Sprintf("%s/cosmos/tx/v1beta1/txs", baseURL), "application/json", reqBz)
// 	s.Require().NoError(err)
// 	err = val.ClientCtx.Codec.UnmarshalJSON(res, &result)
// 	s.Require().NoError(err)
// 	s.Require().Equal(uint32(0), result.TxResponse.Code, "rawlog", result.TxResponse.RawLog)

// 	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
// 	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
// 	s.Require().NoError(err)
// 	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), respType))

// 	balances = respType.(*banktypes.QueryAllBalancesResponse)
// 	coins = balances.Balances
// 	fmt.Println(coins)
// 	s.Require().Equal("100000000", coins.AmountOf(symbol).String())
// 	s.Require().Equal("399981915", coins.AmountOf(sdk.DefaultBondDenom).String())
// 	s.Require().Equal("0", coins.AmountOf(lptDenom).String())

// 	url = fmt.Sprintf("%s/irismod/coinswap/pools/%s", baseURL, lptDenom)
// 	resp, err = rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolResponse))

// 	s.Require().Equal("0", queryPool.Pool.Standard.Amount.String())
// 	s.Require().Equal("0", queryPool.Pool.Token.Amount.String())
// 	s.Require().Equal("0", queryPool.Pool.Lpt.Amount.String())

// 	queryPoolsResponse := proto.Message(&coinswaptypes.QueryLiquidityPoolsResponse{})
// 	url = fmt.Sprintf("%s/irismod/coinswap/pools", baseURL)
// 	resp, err = rest.GetRequest(url)
// 	s.Require().NoError(err)
// 	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, queryPoolsResponse))

// 	queryPools := queryPoolsResponse.(*coinswaptypes.QueryLiquidityPoolsResponse)
// 	s.Require().Len(queryPools.Pools, 1)
// }
