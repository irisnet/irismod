package keeper_test

import (
	"fmt"
	"github.com/irisnet/irismod/modules/nft/types"
)

func (suite *KeeperSuite) TestUserOf() {
	expiry := suite.ctx.BlockTime().Unix() + 10

	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			ClassId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	resp, err := suite.keeper.UserOf(suite.ctx, &types.QueryUserOfRequest{
		ClassId: rentalDenomId,
		NftId:   rentalNftId,
	})
	suite.NoError(err)
	suite.Equal(rentalRenter.String(), resp.User)
}

func (suite *KeeperSuite) TestUserExpires() {
	expiry := suite.ctx.BlockTime().Unix() + 10
	fmt.Print(expiry)
	err := suite.keeper.Rent(suite.ctx,
		types.RentalInfo{
			User:    rentalRenter.String(),
			ClassId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	resp, err := suite.keeper.UserExpires(suite.ctx, &types.QueryUserExpiresRequest{
		ClassId: rentalDenomId,
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
			ClassId: rentalDenomId,
			NftId:   rentalNftId,
			Expires: expiry,
		})
	suite.NoError(err)

	resp, err := suite.keeper.HasUser(suite.ctx, &types.QueryHasUserRequest{
		ClassId: rentalDenomId,
		NftId:   rentalNftId,
	})
	suite.NoError(err)
	suite.Equal(true, resp.HasUser)
}
