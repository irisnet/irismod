package keeper_test

import (
	"encoding/json"
	"github.com/irisnet/irismod/modules/nft/types"
)

var (
	rentalDenomId   = "rentalDenomId"
	rentalDenomName = "rentalDenomName"
	rentalSchema    = "rental schema"
	rentalSymbol    = "ren"
	rentalNftId     = "rentalNftId"
	rentalNftName   = "rentalNftName"

	rentalCreator = CreateTestAddrs(4)[3]
	rentalRenter  = CreateTestAddrs(5)[4]

	rentalUserData = types.DenomUserData{
		RentalEnabled: true,
		UserData:      "",
	}
)

func (suite *KeeperSuite) SetupRentalTest() {

	data, err := json.Marshal(&rentalUserData)
	suite.NoError(err)

	err = suite.keeper.SaveDenom(suite.ctx,
		rentalDenomId,
		rentalDenomName,
		rentalSchema,
		rentalSymbol,
		rentalCreator,
		false,
		false,
		"",
		"",
		"",
		string(data),
	)
	suite.NoError(err)

	err = suite.keeper.SaveNFT(suite.ctx,
		rentalDenomId,
		rentalNftId,
		rentalNftName,
		"",
		"",
		"",
		rentalCreator,
	)
	suite.NoError(err)
}

func (suite *KeeperSuite) TestSetUser() {
	expiry := suite.ctx.BlockTime().Unix() + 10

	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			ClassId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)
}
