package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cometbft/cometbft/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Buy        TradeType = "buy"
	Sell       TradeType = "sell"
	BuyDouble  TradeType = "buyDouble"
	SellDouble TradeType = "SellDouble"
)

type TradeType string
type TradeFunc func(ctx sdk.Context, input Input, output Output) (sdk.Coin, error)

// GetTradeType returns the trade type based on the given parameters.
//
// Parameters:
// - isBuyOrder: a boolean indicating whether the order is a buy order.
// - standardDenom: the standard denomination.
// - inputDenom: the input denomination.
// - outputDenom: the output denomination.
//
// Returns:
// - TradeType: the trade type based on the given parameters.
func GetTradeType(isBuyOrder bool, standardDenom, inputDenom, outputDenom string) TradeType {
	isDoubleSwap := (inputDenom != standardDenom) && (outputDenom != standardDenom)
	if isBuyOrder && isDoubleSwap {
		return BuyDouble
	}
	if isBuyOrder && !isDoubleSwap {
		return Buy
	}
	if isDoubleSwap {
		return SellDouble
	}
	return Sell
}

// GetReservePoolAddr returns the pool address for the provided liquidity denomination.
func GetReservePoolAddr(lptDenom string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(lptDenom)))
}

// GetTokenPairByDenom returns the token pair for the provided denominations
func GetTokenPairByDenom(inputDenom, outputDenom string) string {
	return fmt.Sprintf("%s-%s", outputDenom, inputDenom)
}

// GetPoolId returns the pool-id by counterpartyDenom.
func GetPoolId(counterpartyDenom string) string {
	return fmt.Sprintf("pool-%s", counterpartyDenom)
}

// GetLptDenom returns the pool coin denom by specified sequence.
func GetLptDenom(sequence uint64) string {
	return fmt.Sprintf(LptTokenFormat, sequence)
}

func ParseLptDenom(lptDenom string) (uint64, error) {
	result := strings.Split(lptDenom, "-")
	if len(result) != 2 {
		return 0, fmt.Errorf("invalid lpt denom: %s", lptDenom)
	}
	return strconv.ParseUint(result[1], 10, 64)
}
