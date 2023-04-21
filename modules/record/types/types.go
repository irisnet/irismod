package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/bytes"
)

// NewRecord constructs a new Record instance
func NewRecord(txHash bytes.HexBytes, contents []Content, creator sdk.AccAddress) Record {
	return Record{
		TxHash:   txHash.String(),
		Contents: contents,
		Creator:  creator.String(),
	}
}

func NewGrantRecord(txHash bytes.HexBytes, id, key, pubkey, creator string) GrantRecord {
	return GrantRecord{
		TxHash:  txHash.String(),
		Id:      id,
		Key:     key,
		Pubkey:  pubkey,
		Creator: creator,
	}
}
