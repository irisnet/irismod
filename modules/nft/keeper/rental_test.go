package keeper_test

import (
	"encoding/json"
	"github.com/irisnet/irismod/modules/nft/types"
)

var (
	rentalDenomId   = "rentalDenomId"
	rentalDenomId2  = "rentalDenomId2"
	rentalDenomName = "rentalDenomName"
	rentalSchema    = "rental schema"
	rentalSymbol    = "ren"
	rentalNftId     = "rentalNftId"
	rentalNftId2    = "rentalNftId2"
	rentalNftId3    = "rentalNftId3"
	rentalNftName   = "rentalNftName"

	rentalCreator = CreateTestAddrs(4)[3]
	rentalRenter  = CreateTestAddrs(5)[4]

	rentalUserData = types.DenomUserdata{
		&types.RentalMetadata{Enabled: true},
	}
	rentalUserData2 = types.DenomUserdata{
		&types.RentalMetadata{Enabled: false},
	}
)

func (suite *KeeperSuite) SetupRentalTest() {

	data, err := json.Marshal(&rentalUserData)
	suite.NoError(err)
	data2, err := json.Marshal(&rentalUserData2)
	suite.NoError(err)

	// rentalDenomId enables the rental feature
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

	// rentalDenomId2  disables the rental feature
	err = suite.keeper.SaveDenom(suite.ctx,
		rentalDenomId2,
		rentalDenomName,
		rentalSchema,
		rentalSymbol,
		rentalCreator,
		false,
		false,
		"",
		"",
		"",
		string(data2),
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

	err = suite.keeper.SaveNFT(suite.ctx,
		rentalDenomId,
		rentalNftId2,
		rentalNftName,
		"",
		"",
		"",
		rentalCreator,
	)
	suite.NoError(err)

	err = suite.keeper.SaveNFT(suite.ctx,
		rentalDenomId2,
		rentalNftId3,
		rentalNftName,
		"",
		"",
		"",
		rentalCreator,
	)
	suite.NoError(err)

}

func (suite *KeeperSuite) TestSetUser() {
	expiry := suite.ctx.BlockTime().Unix() + 100
	expiry2 := expiry - 100

	// able to rent
	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			DenomId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	// unable to rent for invalid expiry
	err = suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			DenomId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry2,
		})
	suite.Error(err)

	// unable to rent for not enabling rental
	err = suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			DenomId: rentalDenomId2,
			NftId:   rentalNftId3,
			Expires: expiry2,
		})
	suite.Error(err)
}

func (suite *KeeperSuite) TestUserOf() {
	expiry := suite.ctx.BlockTime().Unix() + 10

	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			DenomId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	resp, err := suite.keeper.UserOf(suite.ctx, &types.QueryUserOfRequest{
		DenomId: rentalDenomId,
		NftId:   rentalNftId,
	})
	suite.NoError(err)
	suite.Equal(rentalRenter.String(), resp.User)
}

func (suite *KeeperSuite) TestUserExpires() {
	expiry := suite.ctx.BlockTime().Unix() + 10
	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			DenomId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	resp, err := suite.keeper.UserExpires(suite.ctx, &types.QueryUserExpiresRequest{
		DenomId: rentalDenomId,
		NftId:   rentalNftId,
	})
	suite.NoError(err)
	suite.Equal(expiry, resp.Expires)
}

func (suite *KeeperSuite) TestHasUser() {
	expiry := suite.ctx.BlockTime().Unix() + 10

	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			DenomId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	resp, err := suite.keeper.HasUser(suite.ctx, &types.QueryHasUserRequest{
		DenomId: rentalDenomId,
		NftId:   rentalNftId,
	})
	suite.NoError(err)
	suite.Equal(true, resp.HasUser)
}
