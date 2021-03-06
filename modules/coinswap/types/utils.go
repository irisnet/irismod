package types

import (
	"fmt"
	"strings"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetReservePoolAddr returns the pool address for the provided liquidity denomination.
func GetReservePoolAddr(uniDenom string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(uniDenom)))
}

// GetTokenPairByDenom returns the token pair for the provided denominations
func GetTokenPairByDenom(inputDenom, outputDenom string) string {
	return fmt.Sprintf("%s-%s", outputDenom, inputDenom)
}

// GetUniDenomFromDenom returns the uni denom for the provided denomination.
func GetUniDenomFromDenom(denom string) string {
	return fmt.Sprintf(FormatUniDenom, denom)
}

// GetCoinDenomFromUniDenom returns the token denom by uni denom
func GetCoinDenomFromUniDenom(uniDenom string) (string, error) {
	if err := ValidateUniDenom(uniDenom); err != nil {
		return "", err
	}
	return strings.TrimPrefix(uniDenom, FormatUniABSPrefix), nil
}
