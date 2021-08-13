package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState is the constructor function for GenesisState
func NewGenesisState(params Params, denom string) *GenesisState {
	return &GenesisState{
		Params:        params,
		StandardDenom: denom,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
		StandardDenom: sdk.DefaultBondDenom,
		Sequence:      1,
	}
}

// ValidateGenesis validates the given genesis state
func ValidateGenesis(data GenesisState) error {
	if err := sdk.ValidateDenom(data.StandardDenom); err != nil {
		return err
	}

	var poolIds = make(map[string]bool, len(data.Pool))
	var lptDenoms = make(map[string]bool, len(data.Pool))
	var maxSequence = uint64(1)
	for _, pool := range data.Pool {
		if poolIds[pool.Id] {
			return fmt.Errorf("duplicate pool: %s", pool.Id)
		}
		if lptDenoms[pool.LptDenom] {
			return fmt.Errorf("duplicate lptDenom: %s", pool.LptDenom)
		}
		//validate the liquidity pool token denom
		seq, err := ParseLptDenom(pool.LptDenom)
		if err != nil {
			return err
		}

		if seq > maxSequence {
			maxSequence = seq
		}

		//validate the token denom
		if err := sdk.ValidateDenom(pool.CounterpartyDenom); err != nil {
			return err
		}

		//validate the token denom
		if err := sdk.ValidateDenom(pool.StandardDenom); err != nil {
			return err
		}

		//validate the address
		if _, err := sdk.AccAddressFromBech32(pool.EscrowAddress); err != nil {
			return err
		}
	}
	if maxSequence != data.Sequence {
		return fmt.Errorf("invalid sequence: %d", data.Sequence)
	}
	return data.Params.Validate()
}
