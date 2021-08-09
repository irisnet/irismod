package types

const (
	// ModuleName is the name of the module.
	ModuleName = "coinswap"

	// RouterKey is the message route for the coinswap module.
	RouterKey = ModuleName

	// StoreKey is the default store key for the coinswap module.
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the coinswap module.
	QuerierRoute = StoreKey

	// KeyNextPoolSequence is the key used to store the next pool sequence in
	// the keeper.
	KeyNextPoolSequence = "nextPoolSequence"

	// KeyPool is the key used to store the pool information  in
	// the keeper.
	KeyPool = "pool"
)
