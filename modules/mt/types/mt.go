package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/mt/exported"
)

var _ exported.MT = MT{}

// NewMT creates a new MT instance
func NewMT(id string, supply uint64, owner sdk.AccAddress, data []byte) MT {
	return MT{
		Id:     id,
		Supply: supply,
		Data:   data,
		Owner:  owner.String(),
	}
}

// GetID return the id of MT
func (bmt MT) GetID() string {
	return bmt.Id
}

func (bmt MT) GetSupply() uint64 {
	return bmt.Supply
}

// GetData return the Data of MT
func (bmt MT) GetData() []byte {
	return bmt.Data
}

// GetOwner return the owner of MT
func (bmt MT) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(bmt.Owner)
	return owner
}

// ----------------------------------------------------------------------------
// MT

// MTs define a list of MT
type MTs []exported.MT

// NewMTs creates a new set of MTs
func NewMTs(mts ...exported.MT) MTs {
	if len(mts) == 0 {
		return MTs{}
	}
	return MTs(mts)
}
