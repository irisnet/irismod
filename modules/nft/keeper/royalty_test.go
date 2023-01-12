package keeper_test

import (
	"math/big"

	"github.com/irisnet/irismod/modules/nft/types"

	sdkmath "cosmossdk.io/math"
)

var (
	denomDefaultFraction = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000))
	tokenFraction        = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(100))

	token1SalePrice     = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000000))
	token1RoyaltyAmount = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(10000))
	token1Fraction      = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(100))

	token2SalePrice     = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000000))
	token2RoyaltyAmount = sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(100000))
)

func (suite *KeeperSuite) TestSaveDefaultRoyalty() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)
	receiver, fraction, err := suite.keeper.GetDefaultRoyaltyInfo(suite.ctx, denomID)
	suite.NoError(err)
	suite.Equal(fraction, denomDefaultFraction)
	suite.Equal(receiver, address.String())
}

func (suite *KeeperSuite) TestRemoveDefaultRoyalty() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)
	receiver, fraction, err := suite.keeper.GetDefaultRoyaltyInfo(suite.ctx, denomID)
	suite.NoError(err)
	suite.Equal(fraction, denomDefaultFraction)
	suite.Equal(receiver, address.String())

	// remove royalty info from denom
	err1 := suite.keeper.RemoveDefaultRoyalty(suite.ctx, denomID, address)
	suite.NoError(err1)

	// royalty info is null
	receiver1, fraction1, err := suite.keeper.GetDefaultRoyaltyInfo(suite.ctx, denomID)
	suite.Equal(err, types.ErrNotEnabledRoyalty)
	suite.Equal(fraction1, sdkmath.Uint{})
	suite.Equal(receiver1, "")

}

func (suite *KeeperSuite) TestSaveTokenRoyalty() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)

	err1 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err1)

	err2 := suite.keeper.SaveTokenRoyalty(suite.ctx, denomID, tokenID,
		address.String(), tokenFraction, address)
	suite.NoError(err2)

	receiver, fraction, err := suite.keeper.GetTokenRoyaltyInfo(
		suite.ctx,
		denomID,
		tokenID,
	)
	suite.NoError(err)

	suite.Equal(receiver, address.String())
	suite.Equal(fraction, tokenFraction)

}

func (suite *KeeperSuite) TestRemoveTokenRoyalty() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)

	err1 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err1)

	err2 := suite.keeper.SaveTokenRoyalty(suite.ctx, denomID, tokenID,
		address.String(), tokenFraction, address)
	suite.NoError(err2)

	err3 := suite.keeper.RemoveTokenRoyalty(suite.ctx, denomID, tokenID, address)
	suite.NoError(err3)

	receiver, fraction, err := suite.keeper.GetTokenRoyaltyInfo(
		suite.ctx,
		denomID,
		tokenID,
	)
	suite.Equal(err, types.ErrNullTokenRoyaltyInfo)

	suite.Equal(receiver, "")
	suite.Equal(fraction, sdkmath.Uint{})

}

func (suite *KeeperSuite) TestGetRoyaltyInfo() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)

	err1 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err1)

	err2 := suite.keeper.SaveTokenRoyalty(suite.ctx, denomID, tokenID,
		address.String(), token1Fraction, address)
	suite.NoError(err2)

	receiver1, royaltyAmount1, err3 := suite.keeper.GetRoyaltyInfo(suite.ctx,
		denomID,
		tokenID,
		token1SalePrice,
	)
	suite.NoError(err3)

	suite.Equal(receiver1, address.String())
	suite.Equal(royaltyAmount1, token1RoyaltyAmount)

	err4 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID2, tokenNm2, tokenURI2, tokenURIHash2, "", address)
	suite.NoError(err4)

	receiver2, royaltyAmount2, err5 := suite.keeper.GetRoyaltyInfo(suite.ctx,
		denomID,
		tokenID2,
		token2SalePrice,
	)
	suite.NoError(err5)
	suite.Equal(receiver2, address.String())
	suite.Equal(royaltyAmount2, token2RoyaltyAmount)

}

func (suite *KeeperSuite) TestGetDefaultRoyaltyInfo() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)

	receiver, royaltyFraction, err := suite.keeper.GetDefaultRoyaltyInfo(suite.ctx, denomID)
	suite.NoError(err)

	suite.Equal(receiver, address.String())
	suite.Equal(royaltyFraction, denomDefaultFraction)

}

func (suite *KeeperSuite) TestGetTokenRoyaltyInfo() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)

	err1 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err1)

	err2 := suite.keeper.SaveTokenRoyalty(suite.ctx, denomID, tokenID,
		address.String(), tokenFraction, address)
	suite.NoError(err2)

	receiver, royaltyFraction, err := suite.keeper.GetTokenRoyaltyInfo(suite.ctx, denomID, tokenID)
	suite.NoError(err)

	suite.Equal(receiver, address.String())
	suite.Equal(royaltyFraction, tokenFraction)
}

func (suite *KeeperSuite) TestIsNotEnabledRoyalty() {

	// royalty should be enabled = false
	enabled := suite.keeper.IsNotEnabledRoyalty(suite.ctx, denomID)
	suite.Equal(enabled, false)

	// set royalty
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), denomDefaultFraction, address)
	suite.NoError(err)

	enabled1 := suite.keeper.IsNotEnabledRoyalty(suite.ctx, denomID)
	suite.Equal(enabled1, true)
}
