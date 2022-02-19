package cli_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	mtcli "github.com/irisnet/irismod/modules/mt/client/cli"
	mttestutil "github.com/irisnet/irismod/modules/mt/client/testutil"
	mttypes "github.com/irisnet/irismod/modules/mt/types"
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
	cfg.NumValidators = 2

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

func (s *IntegrationTestSuite) TestNft() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]

	// ---------------------------------------------------------------------------

	from := val.Address
	tokenName := "Kitty Token"
	uri := "uri"
	uriHash := "uriHash"
	description := "description"
	data := "data"
	tokenID := "kitty"
	//owner     := "owner"
	denomName := "name"
	denom := "denom"
	schema := "schema"
	symbol := "symbol"
	mintRestricted := true
	updateRestricted := false

	//------test GetCmdIssueDenom()-------------
	args := []string{
		fmt.Sprintf("--%s=%s", mtcli.FlagDenomName, denomName),
		fmt.Sprintf("--%s=%s", mtcli.FlagSchema, schema),
		fmt.Sprintf("--%s=%s", mtcli.FlagSymbol, symbol),
		fmt.Sprintf("--%s=%s", mtcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", mtcli.FlagURIHash, uriHash),
		fmt.Sprintf("--%s=%s", mtcli.FlagDescription, description),
		fmt.Sprintf("--%s=%s", mtcli.FlagData, data),
		fmt.Sprintf("--%s=%t", mtcli.FlagMintRestricted, mintRestricted),
		fmt.Sprintf("--%s=%t", mtcli.FlagUpdateRestricted, updateRestricted),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)

	bz, err := mttestutil.IssueDenomExec(val.ClientCtx, from.String(), denom, args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	denomID := gjson.Get(txResp.RawLog, "0.events.0.attributes.0.value").String()

	//------test GetCmdQueryDenom()-------------
	respType = proto.Message(&mttypes.Denom{})
	bz, err = mttestutil.QueryDenomExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	denomItem := respType.(*mttypes.Denom)
	s.Require().Equal(denomName, denomItem.Name)
	s.Require().Equal(schema, denomItem.Schema)
	s.Require().Equal(symbol, denomItem.Symbol)
	s.Require().Equal(uri, denomItem.Uri)
	s.Require().Equal(uriHash, denomItem.UriHash)
	s.Require().Equal(description, denomItem.Description)
	s.Require().Equal(data, denomItem.Data)
	s.Require().Equal(mintRestricted, denomItem.MintRestricted)
	s.Require().Equal(updateRestricted, denomItem.UpdateRestricted)

	//------test GetCmdQueryDenoms()-------------
	respType = proto.Message(&mttypes.QueryDenomsResponse{})
	bz, err = mttestutil.QueryDenomsExec(val.ClientCtx)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	denomsResp := respType.(*mttypes.QueryDenomsResponse)
	s.Require().Equal(1, len(denomsResp.Denoms))
	s.Require().Equal(denomID, denomsResp.Denoms[0].Id)

	//------test GetCmdMintMT()-------------
	args = []string{
		fmt.Sprintf("--%s=%s", mtcli.FlagData, data),
		fmt.Sprintf("--%s=%s", mtcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", mtcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", mtcli.FlagURIHash, uriHash),
		fmt.Sprintf("--%s=%s", mtcli.FlagTokenName, tokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})

	bz, err = mttestutil.MintMTExec(val.ClientCtx, from.String(), denomID, tokenID, args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	//------test GetCmdQuerySupply()-------------
	respType = proto.Message(&mttypes.QuerySupplyResponse{})
	bz, err = mttestutil.QuerySupplyExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	supplyResp := respType.(*mttypes.QuerySupplyResponse)
	s.Require().Equal(uint64(1), supplyResp.Amount)

	//------test GetCmdQueryMT()-------------
	respType = proto.Message(&mttypes.BaseMT{})
	bz, err = mttestutil.QueryMTExec(val.ClientCtx, denomID, tokenID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	mtItem := respType.(*mttypes.BaseMT)
	s.Require().Equal(tokenID, mtItem.Id)
	s.Require().Equal(tokenName, mtItem.Name)
	s.Require().Equal(uri, mtItem.URI)
	s.Require().Equal(uriHash, mtItem.UriHash)
	s.Require().Equal(data, mtItem.Data)
	s.Require().Equal(from.String(), mtItem.Owner)

	//------test GetCmdQueryOwner()-------------
	respType = proto.Message(&mttypes.QueryOwnerResponse{})
	bz, err = mttestutil.QueryOwnerExec(val.ClientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	ownerResp := respType.(*mttypes.QueryOwnerResponse)
	s.Require().Equal(from.String(), ownerResp.Owner.Address)
	s.Require().Equal(denom, ownerResp.Owner.IDCollections[0].DenomId)
	s.Require().Equal(tokenID, ownerResp.Owner.IDCollections[0].TokenIds[0])

	//------test GetCmdQueryCollection()-------------
	respType = proto.Message(&mttypes.QueryCollectionResponse{})
	bz, err = mttestutil.QueryCollectionExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	collectionItem := respType.(*mttypes.QueryCollectionResponse)
	s.Require().Equal(1, len(collectionItem.Collection.MTs))

	//------test GetCmdEditMT()-------------
	newTokenDate := "newdata"
	newTokenURI := "newuri"
	newTokenURIHash := "newuriHash"
	newTokenName := "new Kitty Token"
	args = []string{
		fmt.Sprintf("--%s=%s", mtcli.FlagData, newTokenDate),
		fmt.Sprintf("--%s=%s", mtcli.FlagURI, newTokenURI),
		fmt.Sprintf("--%s=%s", mtcli.FlagURIHash, newTokenURIHash),
		fmt.Sprintf("--%s=%s", mtcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})

	bz, err = mttestutil.EditMTExec(val.ClientCtx, from.String(), denomID, tokenID, args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&mttypes.BaseMT{})
	bz, err = mttestutil.QueryMTExec(val.ClientCtx, denomID, tokenID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	newNftItem := respType.(*mttypes.BaseMT)
	s.Require().Equal(newTokenName, newNftItem.Name)
	s.Require().Equal(newTokenURI, newNftItem.URI)
	s.Require().Equal(newTokenURIHash, newNftItem.UriHash)
	s.Require().Equal(newTokenDate, newNftItem.Data)

	//------test GetCmdTransferMT()-------------
	recipient := sdk.AccAddress(crypto.AddressHash([]byte("dgsbl")))

	args = []string{
		fmt.Sprintf("--%s=%s", mtcli.FlagData, data),
		fmt.Sprintf("--%s=%s", mtcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", mtcli.FlagURIHash, uriHash),
		fmt.Sprintf("--%s=%s", mtcli.FlagTokenName, tokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})

	bz, err = mttestutil.TransferMTExec(val.ClientCtx, from.String(), recipient.String(), denomID, tokenID, args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&mttypes.BaseMT{})
	bz, err = mttestutil.QueryMTExec(val.ClientCtx, denomID, tokenID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	mtItem = respType.(*mttypes.BaseMT)
	s.Require().Equal(tokenID, mtItem.Id)
	s.Require().Equal(tokenName, mtItem.Name)
	s.Require().Equal(uri, mtItem.URI)
	s.Require().Equal(uriHash, mtItem.UriHash)
	s.Require().Equal(data, mtItem.Data)
	s.Require().Equal(recipient.String(), mtItem.Owner)

	//------test GetCmdBurnMT()-------------
	newTokenID := "dgsbl"
	args = []string{
		fmt.Sprintf("--%s=%s", mtcli.FlagData, newTokenDate),
		fmt.Sprintf("--%s=%s", mtcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", mtcli.FlagURI, newTokenURI),
		fmt.Sprintf("--%s=%s", mtcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})

	bz, err = mttestutil.MintMTExec(val.ClientCtx, from.String(), denomID, newTokenID, args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&mttypes.QuerySupplyResponse{})
	bz, err = mttestutil.QuerySupplyExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	supplyResp = respType.(*mttypes.QuerySupplyResponse)
	s.Require().Equal(uint64(2), supplyResp.Amount)

	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
	respType = proto.Message(&sdk.TxResponse{})
	bz, err = mttestutil.BurnMTExec(val.ClientCtx, from.String(), denomID, newTokenID, args...)
	s.Require().NoError(err)
	s.Require().NoError(val2.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&mttypes.QuerySupplyResponse{})
	bz, err = mttestutil.QuerySupplyExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	supplyResp = respType.(*mttypes.QuerySupplyResponse)
	s.Require().Equal(uint64(1), supplyResp.Amount)

	//------test GetCmdTransferDenom()-------------
	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType = proto.Message(&sdk.TxResponse{})

	bz, err = mttestutil.TransferDenomExec(val.ClientCtx, from.String(), val2.Address.String(), denomID, args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp = respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&mttypes.Denom{})
	bz, err = mttestutil.QueryDenomExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	denomItem2 := respType.(*mttypes.Denom)
	s.Require().Equal(val2.Address.String(), denomItem2.Creator)
	s.Require().Equal(denomName, denomItem2.Name)
	s.Require().Equal(schema, denomItem2.Schema)
	s.Require().Equal(symbol, denomItem2.Symbol)
	s.Require().Equal(mintRestricted, denomItem2.MintRestricted)
	s.Require().Equal(updateRestricted, denomItem2.UpdateRestricted)
}
