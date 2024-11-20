package v1

import (
	"math"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestValidateParams(t *testing.T) {
	defaultToken := GetNativeToken()
	tests := []struct {
		testCase string
		Params
		expectPass bool
	}{
		{
			"Minimum value",
			Params{
				TokenTaxRate:      sdkmath.LegacyZeroDec(),
				MintTokenFeeRatio: sdkmath.LegacyZeroDec(),
				IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, sdkmath.ZeroInt()),
			},
			true,
		}, {
			"Maximum value",
			Params{
				TokenTaxRate:      sdkmath.LegacyNewDec(1),
				MintTokenFeeRatio: sdkmath.LegacyNewDec(1),
				IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, sdkmath.NewInt(math.MaxInt64)),
			},
			true,
		}, {
			"TokenTaxRate less than the maximum",
			Params{
				TokenTaxRate:      sdkmath.LegacyNewDecWithPrec(-1, 1),
				MintTokenFeeRatio: sdkmath.LegacyNewDec(0),
				IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, sdkmath.NewInt(1)),
			},
			false,
		}, {
			"MintTokenFeeRatio less than the maximum",
			Params{
				TokenTaxRate:      sdkmath.LegacyNewDec(0),
				MintTokenFeeRatio: sdkmath.LegacyNewDecWithPrec(-1, 1),
				IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, sdkmath.NewInt(1)),
			},
			false,
		}, {
			"TokenTaxRate greater than the maximum",
			Params{
				TokenTaxRate:      sdkmath.LegacyNewDecWithPrec(11, 1),
				MintTokenFeeRatio: sdkmath.LegacyNewDec(1),
				IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, sdkmath.NewInt(1)),
			},
			false,
		}, {
			"MintTokenFeeRatio greater than the maximum",
			Params{
				TokenTaxRate:      sdkmath.LegacyNewDec(1),
				MintTokenFeeRatio: sdkmath.LegacyNewDecWithPrec(11, 1),
				IssueTokenBaseFee: sdk.NewCoin(defaultToken.Symbol, sdkmath.NewInt(1)),
			},
			false,
		}, {
			"IssueTokenBaseFee is negative",
			Params{
				TokenTaxRate:      sdkmath.LegacyNewDec(1),
				MintTokenFeeRatio: sdkmath.LegacyNewDec(1),
				IssueTokenBaseFee: sdk.Coin{Denom: defaultToken.Symbol, Amount: sdkmath.NewInt(-1)},
			},
			false,
		},
	}

	for _, tc := range tests {
		if tc.expectPass {
			require.Nil(t, tc.Params.Validate(), "test: %v", tc.testCase)
		} else {
			require.NotNil(t, tc.Params.Validate(), "test: %v", tc.testCase)
		}
	}
}
