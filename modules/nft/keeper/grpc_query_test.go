package keeper_test

import (
	gocontext "context"
	"math/big"

	sdkmath "cosmossdk.io/math"

	"github.com/irisnet/irismod/modules/nft/types"
)

func (suite *KeeperSuite) TestSupply() {
	err := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Supply(gocontext.Background(), &types.QuerySupplyRequest{
		DenomId: denomID,
		Owner:   address.String(),
	})

	suite.NoError(err)
	suite.Equal(1, int(response.Amount))
}

func (suite *KeeperSuite) TestOwner() {
	err := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.NFTsOfOwner(gocontext.Background(), &types.QueryNFTsOfOwnerRequest{
		DenomId: denomID,
		Owner:   address.String(),
	})

	suite.NoError(err)
	suite.NotNil(response.Owner)
	suite.Contains(response.Owner.IDCollections[0].TokenIds, tokenID)
}

func (suite *KeeperSuite) TestCollection() {
	err := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Collection(gocontext.Background(), &types.QueryCollectionRequest{
		DenomId: denomID,
	})

	suite.NoError(err)
	suite.NotNil(response.Collection)
	suite.Len(response.Collection.NFTs, 1)
	suite.Equal(response.Collection.NFTs[0].Id, tokenID)
}

func (suite *KeeperSuite) TestDenom() {
	err := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Denom(gocontext.Background(), &types.QueryDenomRequest{
		DenomId: denomID,
	})

	suite.NoError(err)
	suite.NotNil(response.Denom)
	suite.Equal(response.Denom.Id, denomID)
}

func (suite *KeeperSuite) TestDenoms() {
	err := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Denoms(gocontext.Background(), &types.QueryDenomsRequest{})

	suite.NoError(err)
	suite.NotEmpty(response.Denoms)
	suite.Equal(response.Denoms[0].Id, denomID)
}

func (suite *KeeperSuite) TestNFT() {
	err := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.NFT(gocontext.Background(), &types.QueryNFTRequest{
		DenomId: denomID,
		TokenId: tokenID,
	})

	suite.NoError(err)
	suite.NotEmpty(response.NFT)
	suite.Equal(response.NFT.Id, tokenID)
}

func (suite *KeeperSuite) TestFeeDenominator() {
	response, err := suite.keeper.FeeDenominator(suite.ctx, &types.MsgFeeDenominatorRequest{})
	suite.NoError(err)
	suite.Equal(response.RoyaltyFraction, sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(10000)))
}

func (suite *KeeperSuite) TestDefaultRoyaltyInfo() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000)), address)
	suite.NoError(err)

	response, err := suite.keeper.DefaultRoyaltyInfo(suite.ctx, &types.MsgDefaultRoyaltyInfoRequest{
		DenomId: denomID,
	})
	suite.NoError(err)

	suite.Equal(response.Receiver, address.String())
	suite.Equal(response.RoyaltyFraction, sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000)))

}

func (suite *KeeperSuite) TestTokenRoyaltyInfo() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000)), address)
	suite.NoError(err)

	err1 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err1)

	err2 := suite.keeper.SaveTokenRoyalty(suite.ctx, denomID, tokenID,
		address.String(), sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000)), address)
	suite.NoError(err2)

	response, err := suite.keeper.TokenRoyaltyInfo(suite.ctx, &types.MsgTokenRoyaltyInfoRequest{
		DenomId: denomID,
		NftId:   tokenID,
	})
	suite.NoError(err)

	suite.Equal(response.Receiver, address.String())
	suite.Equal(response.RoyaltyFraction, sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000)))

}

func (suite *KeeperSuite) TestRoyaltyInfo() {
	err := suite.keeper.SaveDefaultRoyalty(suite.ctx, denomID,
		address.String(), sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(100)), address)
	suite.NoError(err)

	err1 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID, tokenNm, tokenURI, tokenURIHash, tokenData, address)
	suite.NoError(err1)

	err2 := suite.keeper.SaveTokenRoyalty(suite.ctx, denomID, tokenID,
		address.String(), sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(10)), address)
	suite.NoError(err2)

	response1, err3 := suite.keeper.RoyaltyInfo(suite.ctx, &types.MsgRoyaltyInfoRequest{
		DenomId:   denomID,
		NftId:     tokenID,
		SalePrice: sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000000)),
	})
	suite.NoError(err3)
	suite.Equal(response1.RoyaltyAmount, sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000)))

	err4 := suite.keeper.SaveNFT(suite.ctx, denomID, tokenID2, tokenNm2, tokenURI2, tokenURIHash2, "", address)
	suite.NoError(err4)

	response2, err5 := suite.keeper.RoyaltyInfo(suite.ctx, &types.MsgRoyaltyInfoRequest{
		DenomId:   denomID,
		NftId:     tokenID2,
		SalePrice: sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(1000000)),
	})
	suite.NoError(err5)
	suite.Equal(response2.RoyaltyAmount, sdkmath.NewUintFromBigInt(new(big.Int).SetUint64(10000)))

}
