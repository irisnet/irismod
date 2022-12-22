package cli_test

//
//import (
//	"fmt"
//
//	"github.com/cosmos/cosmos-sdk/client/flags"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/gogo/protobuf/proto"
//	"github.com/irisnet/irismod/modules/nft/client/cli"
//	"github.com/irisnet/irismod/modules/nft/client/testutil"
//)
//
//func (s *IntegrationTestSuite) TestRental() {
//	owner := s.network.Validators[0]
//	user := s.network.Validators[1]
//
//	// 1. issue a denom with rental enabled.
//	denomID := "denom"
//	denomName := "denomName"
//	denomSchema := "denomSchema"
//	denomSymbol := "dsb"
//	mintRestricted := false
//	updateRestricted := true
//	denomDescription := "denom description"
//	denomURI := "denomURI"
//	denomURIHash := "denomURIHash"
//	denomData := "denomData"
//	rentable := true
//
//	args := []string{
//		fmt.Sprintf("--%s=%s", cli.FlagDenomName, denomName),
//		fmt.Sprintf("--%s=%s", cli.FlagSchema, denomSchema),
//		fmt.Sprintf("--%s=%s", cli.FlagSymbol, denomSymbol),
//		fmt.Sprintf("--%s=%s", cli.FlagURI, denomURI),
//		fmt.Sprintf("--%s=%s", cli.FlagURIHash, denomURIHash),
//		fmt.Sprintf("--%s=%s", cli.FlagDescription, denomDescription),
//		fmt.Sprintf("--%s=%s", cli.FlagData, denomData),
//		fmt.Sprintf("--%s=%t", cli.FlagMintRestricted, mintRestricted),
//		fmt.Sprintf("--%s=%t", cli.FlagUpdateRestricted, updateRestricted),
//		fmt.Sprintf("--%s=%t", cli.FlagRentable, rentable),
//
//		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
//		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
//		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
//	}
//
//	respType := proto.Message(&sdk.TxResponse{})
//	expectedCode := uint32(0)
//
//	bz, err := testutil.IssueDenomExec(owner.ClientCtx, owner.Address.String(), denomID, args...)
//	s.Require().NoError(err)
//	s.Require().NoError(owner.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
//	txResp := respType.(*sdk.TxResponse)
//	s.Require().Equal(expectedCode, txResp.Code)
//
//	// 2. mint an nft
//	tokenID := "token"
//	tokenName := "tokenName"
//	tokenURI := "tokenURI"
//	tokenURIHash := "tokenURIHash"
//	tokenData := "tokenData"
//
//	args = []string{
//		fmt.Sprintf("--%s=%s", cli.FlagData, tokenData),
//		fmt.Sprintf("--%s=%s", cli.FlagRecipient, owner.Address.String()),
//		fmt.Sprintf("--%s=%s", cli.FlagURI, tokenURI),
//		fmt.Sprintf("--%s=%s", cli.FlagURIHash, tokenURIHash),
//		fmt.Sprintf("--%s=%s", cli.FlagTokenName, tokenName),
//
//		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
//		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
//		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
//	}
//
//	respType = proto.Message(&sdk.TxResponse{})
//
//	bz, err = testutil.MintNFTExec(owner.ClientCtx, owner.Address.String(), denomID, tokenID, args...)
//	s.Require().NoError(err)
//	s.Require().NoError(owner.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
//	txResp = respType.(*sdk.TxResponse)
//	s.Require().Equal(expectedCode, txResp.Code)
//
//	// 3. set user for that nft
//
//	expiry := "2524579200" // 2050-01-01-00:00:00
//
//	args = []string{
//		fmt.Sprintf("--%s=%s", cli.FlagRentalUser, user.Address.String()),
//		fmt.Sprintf("--%s=%s", cli.FlagRentalExpiry, expiry),
//	}
//
//	respType = proto.Message(&sdk.TxResponse{})
//
//	bz, err = testutil.RentalSetUser(owner.ClientCtx, owner.Address.String(), denomID, tokenID, args...)
//	//s.Require().NoError(err)
//	// s.Require().NoError(owner.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
//	owner.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType)
//	fmt.Print(respType)
//	//txResp = respType.(*sdk.TxResponse)
//	//s.Require().Equal(expectedCode, txResp.Code)
//
//	// 4. get rental info
//}
