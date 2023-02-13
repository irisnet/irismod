package cli_test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	randomcli "github.com/irisnet/irismod/modules/random/client/cli"
	randomtestutil "github.com/irisnet/irismod/modules/random/client/testutil"
	randomtypes "github.com/irisnet/irismod/modules/random/types"
	servicecli "github.com/irisnet/irismod/modules/service/client/cli"
	servicetestutil "github.com/irisnet/irismod/modules/service/client/testutil"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
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

	var err error
	s.cfg = cfg
	s.network, err = network.New(s.T(), s.T().TempDir(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestRandom() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	// ---------------------------------------------------------------------------
	serviceDeposit := fmt.Sprintf("50000%s", s.cfg.BondDenom)
	servicePrices := fmt.Sprintf(`{"price": "50%s"}`, s.cfg.BondDenom)
	qos := int64(3)
	options := "{}"
	provider := val.Address

	from := val.Address
	blockInterval := 4
	oracle := true
	serviceFeeCap := fmt.Sprintf("50%s", s.cfg.BondDenom)

	respResult := `{"code":200,"message":""}`
	seedStr := "ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12ABCDEF12"
	respOutput := fmt.Sprintf(`{"header":{},"body":{"seed":"%s"}}`, seedStr)

	// ------bind random service-------------
	args := []string{
		fmt.Sprintf("--%s=%s", servicecli.FlagServiceName, randomtypes.ServiceName),
		fmt.Sprintf("--%s=%s", servicecli.FlagDeposit, serviceDeposit),
		fmt.Sprintf("--%s=%s", servicecli.FlagPricing, servicePrices),
		fmt.Sprintf("--%s=%d", servicecli.FlagQoS, qos),
		fmt.Sprintf("--%s=%s", servicecli.FlagOptions, options),
		fmt.Sprintf("--%s=%s", servicecli.FlagProvider, provider),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)
	bz, err := servicetestutil.BindServiceExec(clientCtx, provider.String(), args...)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	// ------test GetCmdRequestRandom()-------------
	args = []string{
		fmt.Sprintf("--%s=%s", randomcli.FlagServiceFeeCap, serviceFeeCap),
		fmt.Sprintf("--%s=%t", randomcli.FlagOracle, oracle),
		fmt.Sprintf("--%s=%d", randomcli.FlagBlockInterval, blockInterval),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})
	expectedCode = uint32(0)

	bz, err = randomtestutil.RequestRandomExec(clientCtx, from.String(), args...)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)
	requestID := gjson.Get(txResp.RawLog, "0.events.1.attributes.0.value").String()
	requestHeight := gjson.Get(txResp.RawLog, "0.events.1.attributes.2.value").Int()

	// ------test GetCmdQueryRandomRequestQueue()-------------
	respType = proto.Message(&randomtypes.QueryRandomRequestQueueResponse{})
	bz, err = randomtestutil.QueryRandomRequestQueueExec(clientCtx, fmt.Sprintf("%d", requestHeight))
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	qrrResp := respType.(*randomtypes.QueryRandomRequestQueueResponse)
	s.Require().NoError(err)
	s.Require().Len(qrrResp.Requests, 1)

	// ------get service request-------------
	requestHeight = requestHeight + 1
	_, err = s.network.WaitForHeightWithTimeout(requestHeight, time.Duration(int64(blockInterval+2)*int64(s.cfg.TimeoutCommit)))
	s.Require().NoError(err)

	blockResult, err := clientCtx.Client.BlockResults(context.Background(), &requestHeight)
	s.Require().NoError(err)
	var requestId string
	for _, event := range blockResult.EndBlockEvents {
		if event.Type == servicetypes.EventTypeNewBatchRequestProvider {
			var found bool
			var requestIds []string
			var requestsBz []byte
			for _, attribute := range event.Attributes {
				if string(attribute.Key) == servicetypes.AttributeKeyRequests {
					requestsBz = []byte(attribute.GetValue())
					found = true
				}
			}
			s.Require().True(found)
			if found {
				err := json.Unmarshal(requestsBz, &requestIds)
				s.Require().NoError(err)
			}
			s.Require().Len(requestIds, 1)
			requestId = requestIds[0]
		}
	}
	s.Require().NotNil(requestId)

	// ------respond service request-------------
	args = []string{
		fmt.Sprintf("--%s=%s", servicecli.FlagRequestID, requestId),
		fmt.Sprintf("--%s=%s", servicecli.FlagResult, respResult),
		fmt.Sprintf("--%s=%s", servicecli.FlagData, respOutput),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
	respType = proto.Message(&sdk.TxResponse{})
	expectedCode = uint32(0)
	bz, err = servicetestutil.RespondServiceExec(clientCtx, provider.String(), args...)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)
	generateHeight := txResp.Height

	// ------test GetCmdQueryRandom()-------------
	respType = proto.Message(&randomtypes.Random{})
	bz, err = randomtestutil.QueryRandomExec(clientCtx, requestID)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	randomResp := respType.(*randomtypes.Random)
	s.Require().NoError(err)
	s.Require().NotNil(randomResp.Value)

	generateBLock, err := clientCtx.Client.Block(context.Background(), &generateHeight)
	s.Require().NoError(err)
	seed, err := hex.DecodeString(seedStr)
	s.Require().NoError(err)
	random := randomtypes.MakePRNG(generateBLock.Block.LastBlockID.Hash, generateBLock.Block.Header.Time.Unix(), from, seed, true).GetRand().FloatString(randomtypes.RandPrec)
	s.Require().Equal(random, randomResp.Value)
}
