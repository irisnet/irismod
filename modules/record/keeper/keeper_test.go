package keeper_test

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"mods.irisnet.org/modules/record/keeper"
	"mods.irisnet.org/modules/record/types"
	"mods.irisnet.org/simapp"
)

var testCreator = sdk.AccAddress(tmhash.Sum([]byte("test-creator")))

type KeeperTestSuite struct {
	suite.Suite

	cdc    codec.Codec
	ctx    sdk.Context
	keeper keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	depInjectOptions := simapp.DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{&suite.keeper},
	}

	app := simapp.Setup(suite.T(), false, depInjectOptions)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(false)
	suite.keeper.SetIntraTxCounter(suite.ctx, 0)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestAddRecord() {
	content := types.Content{
		Digest:     "test",
		DigestAlgo: "SHA256",
		URI:        "localhost:1317",
		Meta:       "test",
	}
	testRecord := types.NewRecord([]byte("test"), []types.Content{content}, testCreator)

	recordID := suite.keeper.AddRecord(suite.ctx, testRecord)
	addedRecord, found := suite.keeper.GetRecord(suite.ctx, recordID)
	suite.True(found)
	suite.Equal(testRecord, addedRecord)

	// check IntraTxCounter
	suite.Equal(uint32(1), suite.keeper.GetIntraTxCounter(suite.ctx))

	// add the same record, return different record id
	recordID2 := suite.keeper.AddRecord(suite.ctx, testRecord)
	suite.NotEqual(recordID, recordID2)
	addedRecord2, found := suite.keeper.GetRecord(suite.ctx, recordID2)
	suite.True(found)
	suite.Equal(testRecord, addedRecord2)

	recordsIterator := suite.keeper.RecordsIterator(suite.ctx)
	defer recordsIterator.Close()
	var records []types.Record
	for ; recordsIterator.Valid(); recordsIterator.Next() {
		var record types.Record
		suite.cdc.MustUnmarshal(recordsIterator.Value(), &record)
		records = append(records, record)
	}
	suite.Equal(2, len(records))
	suite.Equal(testRecord, records[0])
	suite.Equal(testRecord, records[1])
}
