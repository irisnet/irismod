package types

const (
	// ModuleName is the name of the module
	ModuleName = "nft"

	// StoreKey is the default store key for NFT
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the NFT store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for the NFT module
	RouterKey = ModuleName
)

// RoyaltyKey
var (
	PrefixFeeDenominator = []byte{0x01}
	PrefixDefaultRoyalty = []byte{0x02}
	PrefixTokenRoyalty   = []byte{0x03}
	delimiter            = []byte("/")
)

func KeyFeeDenominator(classId string) []byte {
	key := append(PrefixFeeDenominator, delimiter...)
	return append(key, []byte(classId)...)
}

func KeyDefaultRoyalty(classId string) []byte {
	key := append(PrefixDefaultRoyalty, delimiter...)
	return append(key, []byte(classId)...)
}

func KeyTokenRoyalty(classId string, tokenId string) []byte {
	key := append(PrefixTokenRoyalty, delimiter...)
	if len(classId) > 0 {
		key = append(key, []byte(classId)...)
		key = append(key, delimiter...)
	}
	if len(classId) > 0 && len(tokenId) > 0 {
		key = append(key, []byte(tokenId)...)
	}
	return key
}
