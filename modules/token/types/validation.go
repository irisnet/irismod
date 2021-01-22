package types

import (
	fmt "fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// MaximumMaxSupply limitation for token max supply，1000 billion
	MaximumMaxSupply = uint64(1000000000000)
	// MaximumInitSupply limitation for token initial supply，100 billion
	MaximumInitSupply = uint64(100000000000)
	// MaximumScale limitation for token decimal
	MaximumScale = uint32(9)
	// MinimumSymbolLen limitation for the length of the token's symbol / canonical_symbol
	MinimumSymbolLen = 3
	// MaximumSymbolLen limitation for the length of the token's symbol / canonical_symbol
	MaximumSymbolLen = 64
	// MaximumNameLen limitation for the length of the token's name
	MaximumNameLen = 32
	// MinimumMinUnitLen limitation for the length of the token's min_unit
	MinimumMinUnitLen = 3
	// MaximumMinUnitLen limitation for the length of the token's min_unit
	MaximumMinUnitLen = 64
)

var (
	keywords = strings.Join([]string{
		"peg", "ibc", "swap",
	}, "|")
	keywordsRegex = fmt.Sprintf("^(%s).*", keywords)

	// IsAlphaNumeric only accepts alphanumeric characters
	IsAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	// IsBeginWithAlpha only begin with chars [a-zA-Z]
	IsBeginWithAlpha = regexp.MustCompile(`^[a-zA-Z].*`).MatchString
	// IsBeginWithKeyword define a group of keyword and denom shoule not begin with it
	IsBeginWithKeyword = regexp.MustCompile(keywordsRegex).MatchString
)

// ValidateToken checks if the given token is valid
func ValidateToken(token Token) error {
	_, err := sdk.AccAddressFromBech32(token.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	nameLen := len(token.Name)
	if nameLen == 0 || nameLen > MaximumNameLen {
		return sdkerrors.Wrapf(ErrInvalidName, "invalid token name %s, only accepts length (0, %d]", token.Name, MaximumNameLen)
	}

	if err := ValidateSymbol(token.Symbol); err != nil {
		return err
	}

	if err := ValidateMinUnit(token.MinUnit); err != nil {
		return err
	}

	if err := ValidateInitialSupply(token.InitialSupply); err != nil {
		return err
	}

	if err := ValidateMaxSupply(token.MaxSupply); err != nil {
		return err
	}

	if token.MaxSupply < token.InitialSupply {
		return sdkerrors.Wrapf(ErrInvalidMaxSupply, "invalid token max supply %d, only accepts value [%d, %d]", token.MaxSupply, token.InitialSupply, MaximumMaxSupply)
	}
	return ValidateScale(token.Scale)
}

// ValidateInitialSupply verifies whether the  parameters are legal
func ValidateInitialSupply(initialSupply uint64) error {
	if initialSupply > MaximumInitSupply {
		return sdkerrors.Wrapf(ErrInvalidInitSupply, "invalid token initial supply %d, only accepts value [0, %d]", initialSupply, MaximumInitSupply)
	}
	return nil
}

// ValidateMaxSupply verifies whether the  parameters are legal
func ValidateMaxSupply(maxSupply uint64) error {
	if maxSupply > MaximumMaxSupply {
		return sdkerrors.Wrapf(ErrInvalidMaxSupply, "invalid token max supply %d, maxSupply %d", maxSupply, MaximumMaxSupply)
	}
	return nil
}

// ValidateName verifies whether the  parameters are legal
func ValidateName(name string) error {
	if DoNotModify != name && len(name) > MaximumNameLen {
		return sdkerrors.Wrapf(ErrInvalidName, "invalid token name %s, only accepts length (0, %d]", name, MaximumNameLen)
	}
	return nil
}

// ValidateScale verifies whether the parameters are legal
func ValidateScale(scale uint32) error {
	if scale > MaximumScale {
		return sdkerrors.Wrapf(ErrInvalidScale, "invalid token scale %d, only accepts value [0, %d]", scale, MaximumScale)
	}
	return nil
}

// ValidateMinUnit checks if the given minUnit is valid
func ValidateMinUnit(minUnit string) error {
	minUnitLen := len(strings.TrimSpace(minUnit))
	if minUnitLen < MinimumMinUnitLen || minUnitLen > MaximumMinUnitLen {
		return sdkerrors.Wrapf(ErrInvalidMinUnit, "invalid min_unit %s, only accepts length [%d, %d]", minUnit, MinimumMinUnitLen, MaximumMinUnitLen)
	}

	if !IsBeginWithAlpha(minUnit) || !IsAlphaNumeric(minUnit) {
		return sdkerrors.Wrapf(ErrInvalidSymbol, "invalid min_unit: %s, only accepts alphanumeric characters, and begin with an english letter", minUnit)
	}
	return ValidateKeywords(minUnit)
}

// ValidateSymbol checks if the given symbol is valid
func ValidateSymbol(symbol string) error {
	if len(symbol) < MinimumSymbolLen || len(symbol) > MaximumSymbolLen {
		return sdkerrors.Wrapf(ErrInvalidSymbol, "invalid symbol: %s,  only accepts length [%d, %d]", symbol, MinimumSymbolLen, MaximumSymbolLen)
	}

	if !IsBeginWithAlpha(symbol) || !IsAlphaNumeric(symbol) {
		return sdkerrors.Wrapf(ErrInvalidSymbol, "invalid symbol: %s, only accepts alphanumeric characters, and begin with an english letter", symbol)
	}
	return ValidateKeywords(symbol)
}

// ValidateKeywords checks if the given denom begin with `TokenKeywords`
func ValidateKeywords(denom string) error {
	denom = strings.ToLower(strings.TrimSpace(denom))
	if IsBeginWithKeyword(denom) {
		return sdkerrors.Wrapf(ErrInvalidSymbol, "invalid token: %s, can not begin with keyword: (%s)", denom, keywords)
	}
	return nil
}

// ValidateAmount checks if the given denom begin with `TokenKeywords`
func ValidateAmount(amount uint64) error {
	if amount == 0 || amount > MaximumMaxSupply {
		return sdkerrors.Wrapf(ErrInvalidMaxSupply, "invalid token amount %d, only accepts value (0, %d]", amount, MaximumMaxSupply)
	}
	return nil
}
