package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// Validate performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {
	// Fixme!
	panic("Fixme!")
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(renterInfos []RentalInfo) *GenesisState {
	return &GenesisState{
		RenterInfos: renterInfos,
	}
}

// DefaultGenesisState returns a default rental module genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]RentalInfo{})
}

// GetGenesisStateFromAppState returns rental GenesisState given raw application genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data GenesisState) error {
	panic("Fixme!")
}
