package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/irisnet/irismod/modules/coinswap/types"
)

var (
	KeyFee                 = []byte("Fee")
	KeyPoolCreationFee     = []byte("PoolCreationFee")
	KeyTaxRate             = []byte("TaxRate")
	UnilateralLiquidityFee = sdk.NewDecWithPrec(2, 3)
)

type (
	CoinswapKeeper interface {
		GetParams(ctx sdk.Context) types.Params
		SetParams(ctx sdk.Context, params types.Params)
	}

	Params struct {
		Fee             sdk.Dec  `protobuf:"bytes,1,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee"`
		PoolCreationFee sdk.Coin `protobuf:"bytes,2,opt,name=pool_creation_fee,json=poolCreationFee,proto3" json:"pool_creation_fee"`
		TaxRate         sdk.Dec  `protobuf:"bytes,3,opt,name=tax_rate,json=taxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tax_rate"`
	}
)

func Migrate(ctx sdk.Context, k CoinswapKeeper, paramSpace paramstypes.Subspace) error {
	params := GetLegacyParams(ctx, paramSpace)
	newParams := types.Params{
		Fee:                    params.Fee,
		PoolCreationFee:        params.PoolCreationFee,
		TaxRate:                params.TaxRate,
		UnilateralLiquidityFee: UnilateralLiquidityFee,
	}
	k.SetParams(ctx, newParams)
	return nil
}

// GetLegacyParams gets the parameters for the coinswap module.
func GetLegacyParams(ctx sdk.Context, paramSpace paramstypes.Subspace) Params {
	var swapParams Params
	paramSpace.GetParamSet(ctx, &swapParams)
	return swapParams
}

// ParamSetPairs implements paramtypes.KeyValuePairs
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyFee, &p.Fee, nil),
		paramstypes.NewParamSetPair(KeyPoolCreationFee, &p.PoolCreationFee, nil),
		paramstypes.NewParamSetPair(KeyTaxRate, &p.TaxRate, nil),
	}
}
