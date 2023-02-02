package cli_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftcli "github.com/irisnet/irismod/modules/nft/client/cli"
	nfttestutil "github.com/irisnet/irismod/modules/nft/client/testutil"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
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

func (s *IntegrationTestSuite) TestNft() {
	val := s.network.Validators[0]
	val2 := s.network.Validators[1]
	clientCtx := val.ClientCtx
	expectedCode := uint32(0)

	// ---------------------------------------------------------------------------

	from := val.Address
	tokenName := "Kitty Token"
	uri := "uri"
	uriHash := "uriHash"
	description := "description"
	data := "{\"key1\":\"value1\",\"key2\":\"value2\"}"
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
		fmt.Sprintf("--%s=%s", nftcli.FlagDenomName, denomName),
		fmt.Sprintf("--%s=%s", nftcli.FlagSchema, schema),
		fmt.Sprintf("--%s=%s", nftcli.FlagSymbol, symbol),
		fmt.Sprintf("--%s=%s", nftcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", nftcli.FlagURIHash, uriHash),
		fmt.Sprintf("--%s=%s", nftcli.FlagDescription, description),
		fmt.Sprintf("--%s=%s", nftcli.FlagData, data),
		fmt.Sprintf("--%s=%t", nftcli.FlagMintRestricted, mintRestricted),
		fmt.Sprintf("--%s=%t", nftcli.FlagUpdateRestricted, updateRestricted),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	txResult := nfttestutil.IssueDenomExec(s.T(),
		s.network,
		clientCtx, from.String(), denom, args...)
	s.Require().Equal(expectedCode, txResult.Code)
	denomID := gjson.Get(txResult.Log, "0.events.0.attributes.0.value").String()

	//------test GetCmdQueryDenom()-------------
	queryDenomResponse := &nfttypes.Denom{}
	nfttestutil.QueryDenomExec(s.T(), s.network, clientCtx, denomID, queryDenomResponse)
	s.Require().Equal(denomName, queryDenomResponse.Name)
	s.Require().Equal(schema, queryDenomResponse.Schema)
	s.Require().Equal(symbol, queryDenomResponse.Symbol)
	s.Require().Equal(uri, queryDenomResponse.Uri)
	s.Require().Equal(uriHash, queryDenomResponse.UriHash)
	s.Require().Equal(description, queryDenomResponse.Description)
	s.Require().Equal(data, queryDenomResponse.Data)
	s.Require().Equal(mintRestricted, queryDenomResponse.MintRestricted)
	s.Require().Equal(updateRestricted, queryDenomResponse.UpdateRestricted)

	//------test GetCmdQueryDenoms()-------------
	queryDenomsResponse := &nfttypes.QueryDenomsResponse{}
	nfttestutil.QueryDenomsExec(s.T(), s.network, clientCtx, queryDenomsResponse)
	s.Require().Equal(1, len(queryDenomsResponse.Denoms))
	s.Require().Equal(denomID, queryDenomsResponse.Denoms[0].Id)

	//------test GetCmdMintNFT()-------------
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagData, data),
		fmt.Sprintf("--%s=%s", nftcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", nftcli.FlagURIHash, uriHash),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, tokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	txResult = nfttestutil.MintNFTExec(s.T(),
		s.network,
		clientCtx, from.String(), denomID, tokenID, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	//------test GetCmdQuerySupply()-------------
	querySupplyResponse := &nfttypes.QuerySupplyResponse{}
	nfttestutil.QuerySupplyExec(s.T(), s.network, clientCtx, denomID, querySupplyResponse)
	s.Require().Equal(uint64(1), querySupplyResponse.Amount)

	//------test GetCmdQueryNFT()-------------
	queryNFTResponse := &nfttypes.BaseNFT{}
	nfttestutil.QueryNFTExec(s.T(), s.network, clientCtx, denomID, tokenID, queryNFTResponse)
	s.Require().Equal(tokenID, queryNFTResponse.Id)
	s.Require().Equal(tokenName, queryNFTResponse.Name)
	s.Require().Equal(uri, queryNFTResponse.URI)
	s.Require().Equal(uriHash, queryNFTResponse.UriHash)
	s.Require().Equal(data, queryNFTResponse.Data)
	s.Require().Equal(from.String(), queryNFTResponse.Owner)

	//------test GetCmdQueryOwner()-------------
	queryNFTsOfOwnerResponse := &nfttypes.QueryNFTsOfOwnerResponse{}
	nfttestutil.QueryOwnerExec(s.T(), s.network, clientCtx, from.String(), queryNFTsOfOwnerResponse)
	s.Require().Equal(from.String(), queryNFTsOfOwnerResponse.Owner.Address)
	s.Require().Equal(denom, queryNFTsOfOwnerResponse.Owner.IDCollections[0].DenomId)
	s.Require().Equal(tokenID, queryNFTsOfOwnerResponse.Owner.IDCollections[0].TokenIds[0])

	//------test GetCmdQueryCollection()-------------
	queryCollectionResponse := &nfttypes.QueryCollectionResponse{}
	nfttestutil.QueryCollectionExec(s.T(), s.network, clientCtx, denomID, queryCollectionResponse)
	s.Require().Equal(1, len(queryCollectionResponse.Collection.NFTs))

	//------test GetCmdEditNFT()-------------
	newTokenData := "{\"key1\":\"value1\",\"key2\":\"value2\"}"
	newTokenURI := "newuri"
	newTokenURIHash := "newuriHash"
	newTokenName := "new Kitty Token"
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagData, newTokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagURI, newTokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagURIHash, newTokenURIHash),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	txResult = nfttestutil.EditNFTExec(s.T(),
		s.network,
		clientCtx, from.String(), denomID, tokenID, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	queryNFTResponse = &nfttypes.BaseNFT{}
	nfttestutil.QueryNFTExec(s.T(), s.network, clientCtx, denomID, tokenID, queryNFTResponse)
	s.Require().Equal(newTokenName, queryNFTResponse.Name)
	s.Require().Equal(newTokenURI, queryNFTResponse.URI)
	s.Require().Equal(newTokenURIHash, queryNFTResponse.UriHash)
	s.Require().Equal(newTokenData, queryNFTResponse.Data)

	//------test GetCmdTransferNFT()-------------
	recipient := sdk.AccAddress(crypto.AddressHash([]byte("dgsbl")))

	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagData, data),
		fmt.Sprintf("--%s=%s", nftcli.FlagURI, uri),
		fmt.Sprintf("--%s=%s", nftcli.FlagURIHash, uriHash),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, tokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	txResult = nfttestutil.TransferNFTExec(s.T(),
		s.network,
		clientCtx, from.String(), recipient.String(), denomID, tokenID, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	queryNFTResponse = &nfttypes.BaseNFT{}
	nfttestutil.QueryNFTExec(s.T(), s.network, clientCtx, denomID, tokenID, queryNFTResponse)
	s.Require().Equal(tokenID, queryNFTResponse.Id)
	s.Require().Equal(tokenName, queryNFTResponse.Name)
	s.Require().Equal(uri, queryNFTResponse.URI)
	s.Require().Equal(uriHash, queryNFTResponse.UriHash)
	s.Require().Equal(data, queryNFTResponse.Data)
	s.Require().Equal(recipient.String(), queryNFTResponse.Owner)

	//------test GetCmdBurnNFT()-------------
	newTokenID := "dgsbl"
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagData, newTokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagURI, newTokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	txResult = nfttestutil.MintNFTExec(s.T(),
		s.network,
		clientCtx, from.String(), denomID, newTokenID, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	querySupplyResponse = &nfttypes.QuerySupplyResponse{}
	nfttestutil.QuerySupplyExec(s.T(), s.network, clientCtx, denomID, querySupplyResponse)
	s.Require().Equal(uint64(2), querySupplyResponse.Amount)

	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}
	txResult = nfttestutil.BurnNFTExec(s.T(),
		s.network,
		clientCtx, from.String(), denomID, newTokenID, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	querySupplyResponse = &nfttypes.QuerySupplyResponse{}
	nfttestutil.QuerySupplyExec(s.T(), s.network, clientCtx, denomID, querySupplyResponse)
	s.Require().Equal(uint64(1), querySupplyResponse.Amount)

	//------test GetCmdTransferDenom()-------------
	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.network.BondDenom, sdk.NewInt(10))).String()),
	}

	txResult = nfttestutil.TransferDenomExec(s.T(),
		s.network,
		clientCtx, from.String(), val2.Address.String(), denomID, args...)
	s.Require().Equal(expectedCode, txResult.Code)

	queryDenomResponse = &nfttypes.Denom{}
	nfttestutil.QueryDenomExec(s.T(), s.network, clientCtx, denomID, queryDenomResponse)

	s.Require().Equal(val2.Address.String(), queryDenomResponse.Creator)
	s.Require().Equal(denomName, queryDenomResponse.Name)
	s.Require().Equal(schema, queryDenomResponse.Schema)
	s.Require().Equal(symbol, queryDenomResponse.Symbol)
	s.Require().Equal(mintRestricted, queryDenomResponse.MintRestricted)
	s.Require().Equal(updateRestricted, queryDenomResponse.UpdateRestricted)
}
