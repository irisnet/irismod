package simulation

import (
	"encoding/json"
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"mods.irisnet.org/modules/htlc/types"
)

var (
	MinTimeLock = uint64(50)
	MaxTimeLock = uint64(34560)
	Deputy      = sdk.AccAddress(crypto.AddressHash([]byte("Deputy")))
)

const (
	BNB_DENOM   = "htltbnb"
	OTHER_DENOM = "htltinc"
)

func RandomizedGenState(simState *module.SimulationState) {
	htlcGenesis := &types.GenesisState{
		Params: types.Params{
			AssetParams: []types.AssetParam{
				{
					Denom: BNB_DENOM,
					SupplyLimit: types.SupplyLimit{
						Limit:          math.NewInt(350000000000000),
						TimeLimited:    false,
						TimeBasedLimit: math.ZeroInt(),
						TimePeriod:     time.Hour,
					},
					Active:        true,
					DeputyAddress: Deputy.String(),
					FixedFee:      math.NewInt(1000),
					MinSwapAmount: math.NewInt(2000),
					MaxSwapAmount: math.NewInt(1000000000000),
					MinBlockLock:  MinTimeLock,
					MaxBlockLock:  MaxTimeLock,
				},
				{
					Denom: OTHER_DENOM,
					SupplyLimit: types.SupplyLimit{
						Limit:          math.NewInt(100000000000000),
						TimeLimited:    true,
						TimeBasedLimit: math.NewInt(50000000000),
						TimePeriod:     time.Hour,
					},
					Active:        false,
					DeputyAddress: Deputy.String(),
					FixedFee:      math.NewInt(1000),
					MinSwapAmount: math.NewInt(2000),
					MaxSwapAmount: math.NewInt(100000000000),
					MinBlockLock:  MinTimeLock,
					MaxBlockLock:  MaxTimeLock,
				},
			},
		},
		Htlcs: []types.HTLC{},
		Supplies: []types.AssetSupply{
			types.NewAssetSupply(
				sdk.NewCoin("htltbnb", math.ZeroInt()),
				sdk.NewCoin("htltbnb", math.ZeroInt()),
				sdk.NewCoin("htltbnb", math.ZeroInt()),
				sdk.NewCoin("htltbnb", math.ZeroInt()),
				time.Duration(0),
			),
			types.NewAssetSupply(
				sdk.NewCoin("htltinc", math.ZeroInt()),
				sdk.NewCoin("htltinc", math.ZeroInt()),
				sdk.NewCoin("htltinc", math.ZeroInt()),
				sdk.NewCoin("htltinc", math.ZeroInt()),
				time.Duration(0),
			),
		},
		PreviousBlockTime: types.DefaultPreviousBlockTime,
	}

	bz, err := json.MarshalIndent(htlcGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(htlcGenesis)
}
