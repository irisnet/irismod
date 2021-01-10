package types

// NewGenesisState is the constructor function for GenesisState
func NewGenesisState(params Params, denom string) *GenesisState {
	return &GenesisState{
		Params:        params,
		StandardDenom: denom,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams(), StandardDenom)
}

// ValidateGenesis - placeholder function
func ValidateGenesis(data GenesisState) error {
	return data.Params.Validate()
}
