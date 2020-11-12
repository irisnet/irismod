package rest_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectype "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"
	tokencli "github.com/irisnet/irismod/modules/token/client/cli"
	tokentestutil "github.com/irisnet/irismod/modules/token/client/testutil"
	"github.com/irisnet/irismod/simapp"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := simapp.NewConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestCoinswap() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	// ---------------------------------------------------------------------------

	from := val.Address
	symbol := "kitty"
	name := "Kitty Token"
	minUnit := "kitty"
	scale := 0
	initialSupply := int64(100000000)
	maxSupply := int64(200000000)
	mintable := true
	baseURL := val.APIAddress
	uniKitty := "uni%3Akitty"

	//------test GetCmdIssueToken()-------------
	args := []string{
		fmt.Sprintf("--%s=%s", tokencli.FlagSymbol, symbol),
		fmt.Sprintf("--%s=%s", tokencli.FlagName, name),
		fmt.Sprintf("--%s=%s", tokencli.FlagMinUnit, minUnit),
		fmt.Sprintf("--%s=%d", tokencli.FlagScale, scale),
		fmt.Sprintf("--%s=%d", tokencli.FlagInitialSupply, initialSupply),
		fmt.Sprintf("--%s=%d", tokencli.FlagMaxSupply, maxSupply),
		fmt.Sprintf("--%s=%t", tokencli.FlagMintable, mintable),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)
	bz, err := tokentestutil.IssueTokenExec(clientCtx, from.String(), args...)

	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
	out, err := simapp.QueryBalancesExec(clientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), respType))
	balances := respType.(*banktypes.QueryAllBalancesResponse)
	initKittyAmount := balances.Balances[0].Amount
	initStakeAmount := balances.Balances[2].Amount

	var account authtypes.AccountI
	respType = proto.Message(&codectype.Any{})
	out, err = simapp.QueryAccountExec(clientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), respType))
	err = clientCtx.InterfaceRegistry.UnpackAny(respType.(*codectype.Any), &account)
	s.Require().NoError(err)

	coinswaptypes.RegisterLegacyAminoCodec(clientCtx.LegacyAmino)
	coinswaptypes.RegisterInterfaces(clientCtx.InterfaceRegistry)

	status, err := clientCtx.Client.Status(context.Background())
	s.Require().NoError(err)
	deadline := status.SyncInfo.LatestBlockTime.Add(time.Minute)

	txConfig := legacytx.StdTxConfig{Cdc: s.cfg.LegacyAmino}
	msgAddLiquidity := &coinswaptypes.MsgAddLiquidity{
		MaxToken:         sdk.NewCoin(symbol, sdk.NewInt(1000)),
		ExactStandardAmt: sdk.NewInt(1000),
		MinLiquidity:     sdk.NewInt(1000),
		Deadline:         deadline.Unix(),
		Sender:           from.String(),
	}

	// prepare txBuilder with msg
	txBuilder := txConfig.NewTxBuilder()
	feeAmount := sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
	err = txBuilder.SetMsgs(msgAddLiquidity)
	s.Require().NoError(err)
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(1000000)

	// setup txFactory
	txFactory := tx.Factory{}.
		WithChainID(val.ClientCtx.ChainID).
		WithKeybase(val.ClientCtx.Keyring).
		WithTxConfig(txConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON).
		WithSequence(account.GetSequence())

	// sign Tx (offline mode so we can manually set sequence number)
	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, true)
	s.Require().NoError(err)

	stdTx := txBuilder.GetTx().(legacytx.StdTx)
	req := authrest.BroadcastReq{
		Tx:   stdTx,
		Mode: "block",
	}
	reqBz, err := val.ClientCtx.LegacyAmino.MarshalJSON(req)
	s.Require().NoError(err)
	_, err = rest.PostRequest(fmt.Sprintf("%s/txs", baseURL), "application/json", reqBz)

	respType = proto.Message(&banktypes.QueryAllBalancesResponse{})
	out, err = simapp.QueryBalancesExec(clientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), respType))
	balances = respType.(*banktypes.QueryAllBalancesResponse)
	s.Require().Equal(initKittyAmount.Int64()- 1000 , balances.Balances[0].Amount.Int64())
	s.Require().Equal(initStakeAmount.Int64()- 1010 , balances.Balances[2].Amount.Int64())

	url := fmt.Sprintf("%s/irismod/coinswap/liquidities/%s", baseURL, uniKitty)
	println(url)
	resp, err := rest.GetRequest(url)
	respType = proto.Message(&coinswaptypes.QueryLiquidityResponse{})
	s.Require().NoError(err)
	println(string(resp))
	//s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(resp, respType))
	//liquidityResp := respType.(*coinswaptypes.QueryLiquidityResponse)
	//s.Require().Equal(1, len(liquidityResp.Liquidity))
}
