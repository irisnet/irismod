package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cosmos/cosmos-sdk/codec"
)

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
func ValidateGenesis(gs GenesisState) error {
	for _, v := range gs.RenterInfos {
		if err := nfttypes.ValidateClassID(v.ClassId); err != nil {
			return sdkerrors.Wrapf(ErrInvalidClassID, "Invalid class id (%s)", v.ClassId)
		}

		if err := nfttypes.ValidateNFTID(v.NftId); err != nil {
			return sdkerrors.Wrapf(ErrInvalidNftID, "Invalid nft id (%s)", v.NftId)

		}

		_, err := sdk.AccAddressFromBech32(v.User)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid user address (%s)", v.User)
		}
	}

	return nil
}
