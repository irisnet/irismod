package testutil

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/tidwall/gjson"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/rest"
	sdk "github.com/cosmos/cosmos-sdk/types"

	recordcli "github.com/irisnet/irismod/modules/record/client/cli"
	recordtypes "github.com/irisnet/irismod/modules/record/types"
)

func (s *IntegrationTestSuite) TestQueryRecordGRPC() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	// ---------------------------------------------------------------------------

	from := val.Address
	digest := "digest"
	digestAlgo := "digest-algo"
	uri := "https://example.abc"
	meta := "meta data"

	args := []string{
		fmt.Sprintf("--%s=%s", recordcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", recordcli.FlagMeta, meta),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)

	bz, err := MsgCreateRecordExec(clientCtx, from.String(), digest, digestAlgo, args...)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	recordID := gjson.Get(txResp.RawLog, "0.events.0.attributes.1.value").String()

	// ---------------------------------------------------------------------------

	baseURL := val.APIAddress
	url := fmt.Sprintf("%s/irismod/record/records/%s", baseURL, recordID)

	respType = proto.Message(&recordtypes.QueryRecordResponse{})
	expectedContents := []recordtypes.Content{{
		Digest:     digest,
		DigestAlgo: digestAlgo,
		URI:        uri,
		Meta:       meta,
	}}

	resp, err := rest.GetRequest(url)
	s.Require().NoError(err)
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(resp, respType))
	record := respType.(*recordtypes.QueryRecordResponse).Record
	s.Require().Equal(expectedContents, record.Contents)
}
