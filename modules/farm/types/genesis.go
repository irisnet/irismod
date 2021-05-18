package types

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(params Params, pools []FarmPool, farmInfos []FarmInfo) *GenesisState {
	return &GenesisState{
		params, pools, farmInfos,
	}
}

// DefaultGenesisState gets the default genesis state for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// ValidateGenesis validates the provided farm genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	//TODO
	return nil
}
