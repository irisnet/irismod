package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewIDCollection creates a new IDCollection instance
func NewIDCollection(denomId string, balances ...Balance) IDCollection {
	return IDCollection{
		DenomId:  denomId,
		Balances: balances,
	}
}

// Supply return the amount of the denom
func (idc IDCollection) Supply() int {
	return len(idc.Balances)
}

// AddBalance adds an tokenID to the idCollection
func (idc IDCollection) AddBalance(balances Balance) IDCollection {
	idc.Balances = append(idc.Balances, balances)
	return idc
}

//// AddBalance adds an tokenID to the idCollection
//func (b Balance) AddAmount(amount uint64) Balance {
//	b.Amount = append(b.Amount, amount)
//	return b
//}

// ----------------------------------------------------------------------------
// Balances is an array of ID Balances
type Balances []Balance

// Add adds an ID to the idCollection
func (b Balances) Add(mtId string, amount uint64) Balances {
	//for i, mt := range b {
	//	if mt.MtId == mtId {
	//		b[i] = mt.AddAmount(amount)
	//		return b
	//	}
	//}
	return append(b, Balance{
		MtId:   mtId,
		Amount: amount,
	})
}

// String follows stringer interface
func (b Balances) String() string {
	if len(b) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, idCollection := range b {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(idCollection.String())
	}
	return buf.String()
}

// ----------------------------------------------------------------------------
// IDCollections is an array of ID Collections
type IDCollections []IDCollection

// Add adds an ID to the idCollection
func (idcs IDCollections) Add(denomID string, balances Balance) IDCollections {
	for i, idc := range idcs {
		if idc.DenomId == denomID {
			idcs[i] = idc.AddBalance(balances)
			return idcs
		}
	}
	return append(idcs, IDCollection{
		DenomId:  denomID,
		Balances: []Balance{balances},
	})
}

// String follows stringer interface
func (idcs IDCollections) String() string {
	if len(idcs) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, idCollection := range idcs {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(idCollection.String())
	}
	return buf.String()
}

// Owner of non fungible tokens
//type Owner struct {
//	Address       sdk.AccAddress `json:"address" yaml:"address"`
//	IDCollections IDCollections  `json:"id_collections" yaml:"id_collections"`
//}

// NewOwner creates a new Owner
func NewOwner(owner sdk.AccAddress, idCollections ...IDCollection) Owner {
	return Owner{
		Address:       owner.String(),
		IdCollections: idCollections,
	}
}

type Owners []Owner

// NewOwner creates a new Owner
func NewOwners(owner ...Owner) Owners {
	return append([]Owner{}, owner...)
}

// String follows stringer interface
func (owners Owners) String() string {
	var buf bytes.Buffer
	for _, owner := range owners {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(owner.String())
	}
	return buf.String()
}
