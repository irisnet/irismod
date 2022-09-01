package v2

var (
	PrefixNFT   = []byte{0x01}
	PrefixDenom = []byte{0x04} // key for denom of the nft

	delimiter = []byte("/")
)

// KeyDenom gets the storeKey by the denom id
func KeyDenom(id string) []byte {
	key := append(PrefixDenom, delimiter...)
	return append(key, []byte(id)...)
}

// KeyNFT gets the key of nft stored by an denom and id
func KeyNFT(denomID, tokenID string) []byte {
	key := append(PrefixNFT, delimiter...)
	if len(denomID) > 0 {
		key = append(key, []byte(denomID)...)
		key = append(key, delimiter...)
	}

	if len(denomID) > 0 && len(tokenID) > 0 {
		key = append(key, []byte(tokenID)...)
	}
	return key
}
