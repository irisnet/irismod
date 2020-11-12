package rest_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectype "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	coinrest "github.com/irisnet/irismod/modules/coinswap/client/rest"
	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"
	tokencli "github.com/irisnet/irismod/modules/token/client/cli"
	tokentestutil "github.com/irisnet/irismod/modules/token/client/testutil"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/irisnet/irismod/simapp"
	"github.com/stretchr/testify/suite"
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

func (s *IntegrationTestSuite) TestHtlc() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	// ---------------------------------------------------------------------------

	from := val.Address
	denomStandard := sdk.DefaultBondDenom
	symbol := "Kitty"
	name := "Kitty Token"
	minUnit := "kitty"
	uniKitty := "uni:kitty"
	scale := 0
	initialSupply := int64(100000000)
	maxSupply := int64(200000000)
	mintable := true
	baseUrl := val.APIAddress

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

	coinType := proto.Message(&sdk.Coin{})
	out, err := simapp.QueryBalancesExec(clientCtx, from.String(), strings.ToLower(symbol))
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), coinType))
	balance := coinType.(*sdk.Coin)
	kittyAmount := balance.Amount.Int64()
	println(kittyAmount)

	coinType = proto.Message(&sdk.Coin{})
	out, err = simapp.QueryBalancesExec(clientCtx, from.String(), denomStandard)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), coinType))
	balance = coinType.(*sdk.Coin)
	stakeAmount := balance.Amount.Int64()
	println(stakeAmount)


	var account authtypes.AccountI
	respType = proto.Message(&codectype.Any{})
	out, err = simapp.QueryAccountExec(clientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), respType))
	err = clientCtx.InterfaceRegistry.UnpackAny(respType.(*codectype.Any), &account)
	s.Require().NoError(err)


	baseReq:=rest.BaseReq{
		From:          from.String(),
		Memo:          "",
		ChainID:       clientCtx.ChainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence()+1,
		Fees:          sdk.NewCoins(),
		GasPrices:     nil,
		Gas:           "10000",
		GasAdjustment: fmt.Sprintf("%f", 1.0),
		Simulate:      clientCtx.Simulate,
	}

	coinReq := coinrest.AddLiquidityReq{
		BaseReq:          baseReq,
		ID:               uniKitty,
		MaxToken:         "2000",
		ExactStandardAmt: "2000",
		MinLiquidity:     "2000",
		Deadline:         "10m0s",
		Sender:           from.String(),
	}

	coinswaptypes.RegisterLegacyAminoCodec(clientCtx.LegacyAmino)
	coinswaptypes.RegisterInterfaces(clientCtx.InterfaceRegistry)
	url := fmt.Sprintf("%s/coinswap/liquidities/%s/deposit", baseUrl, uniKitty)
	reqBz, err := val.ClientCtx.LegacyAmino.MarshalJSON(coinReq)
	s.Require().NoError(err)
	res, err := rest.PostRequest(url, "application/json", reqBz)
	s.Require().NoError(err)
	println(string(res))

	txConfig := legacytx.StdTxConfig{Cdc: s.cfg.LegacyAmino}
	msg := &types.MsgSend{
		FromAddress: val.Address.String(),
		ToAddress:   val.Address.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin(fmt.Sprintf("%stoken", val.Moniker), 100)},
	}

	// prepare txBuilder with msg
	txBuilder := txConfig.NewTxBuilder()
	feeAmount := sdk.Coins{sdk.NewInt64Coin(s.cfg.BondDenom, 10)}
	gasLimit := testdata.NewTestGasLimit()
	txBuilder.SetMsgs(msg)
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetMemo("kitty")

	// setup txFactory
	txFactory := tx.Factory{}.
		WithChainID(val.ClientCtx.ChainID).
		WithKeybase(val.ClientCtx.Keyring).
		WithTxConfig(txConfig).
		WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON).
		WithSequence(account.GetSequence()+1)


	// sign Tx (offline mode so we can manually set sequence number)
	err = authclient.SignTx(txFactory, val.ClientCtx, val.Moniker, txBuilder, true)
	s.Require().NoError(err)

	stdTx := txBuilder.GetTx().(legacytx.StdTx)
	req := authrest.BroadcastReq{
		Tx:   stdTx,
		Mode: "sync",
	}
	reqBz, err = val.ClientCtx.LegacyAmino.MarshalJSON(req)
	s.Require().NoError(err)
	res,err = rest.PostRequest(fmt.Sprintf("%s/txs", val.APIAddress), "application/json", reqBz)
	println(string(res))
}
