package keeper_test

import (
	"math/rand"
	"time"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmtime "github.com/tendermint/tendermint/types/time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/htlc/types"
)

var (
	DenomMap                    = map[int]string{0: "htltbtc", 1: "htlteth", 2: "htltbnb", 3: "htltxrp", 4: "htltdai"}
	TestDeputy                  = sdk.AccAddress(crypto.AddressHash([]byte("TestDeputy")))
	TestUser1                   = sdk.AccAddress(crypto.AddressHash([]byte("TestUser1")))
	TestUser2                   = sdk.AccAddress(crypto.AddressHash([]byte("TestUser2")))
	MinTimeLock          uint64 = 220
	MaxTimeLock          uint64 = 270
	ReceiverOnOtherChain        = "ReceiverOnOtherChain"
	SenderOnOtherChain          = "SenderOnOtherChain"
)

func c(denom string, amount int64) sdk.Coin {
	return sdk.NewInt64Coin(denom, amount)
}

func cs(coins ...sdk.Coin) sdk.Coins {
	return sdk.NewCoins(coins...)
}

func ts(minOffset int) uint64 {
	return uint64(tmtime.Now().Add(time.Duration(minOffset) * time.Minute).Unix())
}

func NewHTLTGenesis(deputyAddress sdk.AccAddress) *types.GenesisState {
	return &types.GenesisState{
		Params: types.Params{
			AssetParams: []types.AssetParam{
				{
					Denom: "htltbnb",
					SupplyLimit: types.SupplyLimit{
						Limit:          sdk.NewInt(350000000000000),
						TimeLimited:    false,
						TimeBasedLimit: sdk.ZeroInt(),
						TimePeriod:     time.Hour,
					},
					Active:        true,
					DeputyAddress: deputyAddress.String(),
					FixedFee:      sdk.NewInt(1000),
					MinSwapAmount: sdk.OneInt(),
					MaxSwapAmount: sdk.NewInt(1000000000000),
					MinBlockLock:  MinTimeLock,
					MaxBlockLock:  MaxTimeLock,
				},
				{
					Denom: "htltinc",
					SupplyLimit: types.SupplyLimit{
						Limit:          sdk.NewInt(100000000000000),
						TimeLimited:    true,
						TimeBasedLimit: sdk.NewInt(50000000000),
						TimePeriod:     time.Hour,
					},
					Active:        false,
					DeputyAddress: deputyAddress.String(),
					FixedFee:      sdk.NewInt(1000),
					MinSwapAmount: sdk.OneInt(),
					MaxSwapAmount: sdk.NewInt(100000000000),
					MinBlockLock:  MinTimeLock,
					MaxBlockLock:  MaxTimeLock,
				},
			},
		},
		Htlcs: []types.HTLC{},
		Supplies: []types.AssetSupply{
			types.NewAssetSupply(
				sdk.NewCoin("htltbnb", sdk.ZeroInt()),
				sdk.NewCoin("htltbnb", sdk.ZeroInt()),
				sdk.NewCoin("htltbnb", sdk.ZeroInt()),
				sdk.NewCoin("htltbnb", sdk.ZeroInt()),
				time.Duration(0),
			),
			types.NewAssetSupply(
				sdk.NewCoin("htltinc", sdk.ZeroInt()),
				sdk.NewCoin("htltinc", sdk.ZeroInt()),
				sdk.NewCoin("htltinc", sdk.ZeroInt()),
				sdk.NewCoin("htltinc", sdk.ZeroInt()),
				time.Duration(0),
			),
		},
		PreviousBlockTime: types.DefaultPreviousBlockTime,
	}
}

// func htlts(ctx sdk.Context, count int) []types.HTLC {
// 	var htlts []types.HTLC
// 	for i := 0; i < count; i++ {
// 		htlt := htlt(ctx, i)
// 		htlts = append(htlts, htlt)
// 	}
// 	return htlts
// }

// func htlt(ctx sdk.Context, index int) types.HTLC {
// 	expireOffset := uint64(200)
// 	timestamp := ts(index)
// 	amount := cs(c("htltbnb", 50000))

// 	secret, _ := generateSecret()
// 	hashLock := types.GetHashLock(secret[:], timestamp)
// 	id := types.GetID(TestUser1, TestUser2, amount, hashLock)

// 	return types.NewHTLC(
// 		id,
// 		TestUser1,
// 		TestUser2,
// 		ReceiverOnOtherChain,
// 		SenderOnOtherChain,
// 		amount,
// 		hashLock,
// 		secret,
// 		timestamp,
// 		uint64(ctx.BlockHeight())+expireOffset,
// 		types.Open,
// 		0,
// 		true,
// 		types.Incoming,
// 	)
// }

// func generateSecret() ([]byte, error) {
// 	bytes := make([]byte, 32)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return []byte{}, err
// 	}
// 	return bytes, nil
// }

// func assetSupplies(count int) []types.AssetSupply {
// 	if count > 5 {
// 		return []types.AssetSupply{}
// 	}

// 	var supplies []types.AssetSupply

// 	for i := 0; i < count; i++ {
// 		supply := assetSupply(DenomMap[i])
// 		supplies = append(supplies, supply)
// 	}
// 	return supplies
// }

// func assetSupply(denom string) types.AssetSupply {
// 	return types.NewAssetSupply(c(denom, 0), c(denom, 0), c(denom, 0), c(denom, 0), time.Duration(0))
// }

// GeneratePrivKeyAddressPairsFromRand generates (deterministically) a total of n secp256k1 private keys and addresses.
func GeneratePrivKeyAddressPairs(n int) (keys []crypto.PrivKey, addrs []sdk.AccAddress) {
	r := rand.New(rand.NewSource(12345)) // make the generation deterministic
	keys = make([]crypto.PrivKey, n)
	addrs = make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		secret := make([]byte, 32)
		if _, err := r.Read(secret); err != nil {
			panic("Could not read randomness")
		}
		keys[i] = secp256k1.GenPrivKeySecp256k1(secret)
		addrs[i] = sdk.AccAddress(keys[i].PubKey().Address())
	}
	return
}

func GenerateRandomSecret() ([]byte, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
