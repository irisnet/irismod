package rest_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	tokencli "github.com/irisnet/irismod/modules/token/client/cli"
	tokentestutil "github.com/irisnet/irismod/modules/token/client/testutil"
	coinrest "github.com/irisnet/irismod/modules/coinswap/client/rest"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	codectype "github.com/cosmos/cosmos-sdk/codec/types"
	coinswaptypes "github.com/irisnet/irismod/modules/coinswap/types"

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
}
