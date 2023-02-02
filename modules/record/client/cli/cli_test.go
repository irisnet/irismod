package cli_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	recordcli "github.com/irisnet/irismod/modules/record/client/cli"
	recordtestutil "github.com/irisnet/irismod/modules/record/client/testutil"
	recordtypes "github.com/irisnet/irismod/modules/record/types"
	"github.com/irisnet/irismod/simapp"
)

type IntegrationTestSuite struct {
	suite.Suite

	network simapp.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.network = simapp.SetupNetwork(s.T())
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestRecord() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	// ---------------------------------------------------------------------------

	from := val.Address
	digest := "digest"
	digestAlgo := "digest-algo"
	uri := "uri"
	meta := "meta"

	args := []string{
		fmt.Sprintf("--%s=%s", recordcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", recordcli.FlagMeta, meta),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	expectedCode := uint32(0)

	txResult := recordtestutil.CreateRecordExec(s.T(),
		s.network,
		clientCtx, from.String(), digest, digestAlgo, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	recordID := gjson.Get(txResult.Log, "0.events.0.attributes.1.value").String()

	// ---------------------------------------------------------------------------

	record := &recordtypes.Record{}
	expectedContents := []recordtypes.Content{{
		Digest:     digest,
		DigestAlgo: digestAlgo,
		URI:        uri,
		Meta:       meta,
	}}

	recordtestutil.QueryRecordExec(s.T(), s.network, clientCtx, recordID, record)
	s.Require().Equal(expectedContents, record.Contents)
}
