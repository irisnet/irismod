package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tidwall/gjson"

	"mods.irisnet.org/modules/service/types"
)

// GetExchangedPrice gets the exchanged price for the specified consumer and binding
// Note: ensure that the binding is valid
func (k Keeper) GetExchangedPrice(
	ctx sdk.Context, consumer sdk.AccAddress, binding types.ServiceBinding,
) (
	sdk.Coins, string, error,
) {
	provider, _ := sdk.AccAddressFromBech32(binding.Provider)
	pricing := k.GetPricing(ctx, binding.ServiceName, provider)

	// get discounts
	discountByTime := types.GetDiscountByTime(pricing, ctx.BlockTime())
	discountByVolume := types.GetDiscountByVolume(
		pricing, k.GetRequestVolume(ctx, consumer, binding.ServiceName, provider),
	)

	baseDenom := k.BaseDenom(ctx)
	rawDenom := pricing.Price.GetDenomByIndex(0)

	rawPrice := pricing.Price.AmountOf(rawDenom)
	price := math.LegacyNewDecFromInt(rawPrice).Mul(discountByTime).Mul(discountByVolume)

	realPrice := price
	if baseDenom != rawDenom {
		rate, err := k.GetExchangeRate(ctx, rawDenom, baseDenom)
		if err != nil {
			return nil, rawDenom, err
		}

		realPrice = price.Mul(rate)
	}

	return sdk.NewCoins(sdk.NewCoin(baseDenom, realPrice.TruncateInt())), rawDenom, nil
}

// GetExchangeRate retrieves the exchange rate of the given pair by the oracle module service
func (k Keeper) GetExchangeRate(ctx sdk.Context, quoteDenom, baseDenom string) (math.LegacyDec, error) {
	exchangeRateSvc, exist := k.GetModuleServiceByModuleName(types.RegisterModuleName)
	if !exist {
		return math.LegacyDec{}, errorsmod.Wrapf(types.ErrInvalidModuleService, "module service does not exist: %s", types.RegisterModuleName)
	}

	inputBody := fmt.Sprintf(`{"pair":"%s-%s"}`, quoteDenom, baseDenom)
	input := fmt.Sprintf(`{"header":{},"body":%s`, inputBody)
	if err := types.ValidateRequestInputBody(types.OraclePriceSchemas, inputBody); err != nil {
		return math.LegacyDec{}, err
	}

	result, output := exchangeRateSvc.ReuquestService(ctx, input)
	if code, msg := CheckResult(result); code != types.ResultOK {
		return math.LegacyDec{}, errorsmod.Wrapf(types.ErrInvalidModuleService, msg)
	}

	outputBody := gjson.Get(output, types.PATH_BODY).String()
	if err := types.ValidateResponseOutputBody(types.OraclePriceSchemas, outputBody); err != nil {
		return math.LegacyDec{}, err
	}

	rate, err := GetExchangeRateFromJSON(outputBody)
	if err != nil {
		return math.LegacyDec{}, err
	}

	if rate.IsZero() {
		return math.LegacyDec{}, errorsmod.Wrapf(types.ErrInvalidResponseOutputBody, "rate can not be zero")
	}

	return rate, nil
}

func CheckResult(jsonStr string) (code types.ResultCode, msg string) {
	code = types.ResultCode(gjson.Get(jsonStr, "code").Uint())
	msg = gjson.Get(jsonStr, "message").String()
	return
}

func GetExchangeRateFromJSON(jsonStr string) (math.LegacyDec, error) {
	result := gjson.Get(jsonStr, types.OraclePriceValueJSONPath)
	return math.LegacyNewDecFromStr(result.String())
}
