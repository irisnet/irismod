package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irisnet/irismod/modules/mt/exported"
)

var _ exported.MT = BaseMT{}

// NewBaseMT creates a new MT instance
func NewBaseMT(id, name string, owner sdk.AccAddress, uri, uriHash, data string) BaseMT {
	return BaseMT{
		Id:      id,
		Name:    name,
		Owner:   owner.String(),
		URI:     uri,
		UriHash: uriHash,
		Data:    data,
	}
}

// GetID return the id of BaseMT
func (bmt BaseMT) GetID() string {
	return bmt.Id
}

// GetName return the name of BaseMT
func (bmt BaseMT) GetName() string {
	return bmt.Name
}

// GetOwner return the owner of BaseMT
func (bmt BaseMT) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(bmt.Owner)
	return owner
}

// GetURI return the URI of BaseMT
func (bmt BaseMT) GetURI() string {
	return bmt.URI
}

// GetURIHash return the UriHash of BaseMT
func (bmt BaseMT) GetURIHash() string {
	return bmt.UriHash
}

// GetData return the Data of BaseMT
func (bmt BaseMT) GetData() string {
	return bmt.Data
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
