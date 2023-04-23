package types

import "fmt"

// NewGenesisState creates a new GenesisState object
func NewGenesisState(tokenPairs []TokenPair) *GenesisState {
	return &GenesisState{
		TokenPairs: tokenPairs,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	seenErc721 := make(map[string]bool)
	seenClass := make(map[string]bool)

	for _, b := range gs.TokenPairs {
		if seenErc721[b.Erc721Address] {
			return fmt.Errorf("token ERC721 contract duplicated on genesis '%s'", b.Erc721Address)
		}
		if seenClass[b.ClassId] {
			return fmt.Errorf("class duplicated on genesis: '%s'", b.ClassId)
		}

		if err := b.Validate(); err != nil {
			return err
		}

		seenErc721[b.Erc721Address] = true
		seenClass[b.ClassId] = true
	}

	return nil
}
