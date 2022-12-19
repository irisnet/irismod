package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/x/nft"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/simapp"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	rental "github.com/irisnet/irismod/modules/rental/types"
)

const (
	testClassID          = "flower"
	testClassName        = "Crypto Flower"
	testClassSymbol      = "flower"
	testClassDescription = "Crypto Flower"
	testClassURI         = "class uri"
	testClassURIHash     = "ae702cefd6b6a65fe2f991ad6d9969ed"
	testID               = "iris"
	testURI              = "iris uri"
	testURIHash          = "229bfd3c1b431c14a526497873897108"

	testAccNumber = 2
	testOwnerIdx  = 0
	testUserIdx   = 1

	testExpires = "999999999999999"
)

type TestSuite struct {
	suite.Suite

	app   *simapp.SimApp
	ctx   sdk.Context
	addrs []sdk.AccAddress

	queryClient rental.QueryClient
}

func (s *TestSuite) SetupTest() {
	app := simapp.Setup(s.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	rental.RegisterQueryServer(queryHelper, app.RentalKeeper)
	queryClient := rental.NewQueryClient(queryHelper)

	s.app = app
	s.ctx = ctx
	s.queryClient = queryClient
	s.addrs = simapp.AddTestAddrsIncremental(app, ctx, testAccNumber, sdk.NewInt(30000000))

	s.prepareNFT()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestRent() {
	// s.app.RentalKeeper.Rent()
}

// prepareNFT setups nft for rental testing
func (s *TestSuite) prepareNFT() {
	testClass := nft.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}

	err := s.app.NFTKeeper.NFTkeeper().SaveClass(s.ctx, testClass)
	s.Require().NoError(err)

	testNft := nft.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
		UriHash: testURIHash,
	}

	err = s.app.NFTKeeper.NFTkeeper().Mint(s.ctx, testNft, s.addrs[testOwnerIdx])
	s.Require().NoError(err)
}
