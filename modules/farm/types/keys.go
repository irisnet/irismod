package types

// nolint
const (
	// module name
	ModuleName = "farm"

	// StoreKey is the default store key for farm
	StoreKey = ModuleName

	// RouterKey is the message route for farm
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the farm store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the farm querier
	QueryRecord = "farm"
)

var (
	FarmPoolKey        = []byte{0x01} // key for farm pool
	FarmPoolRuleKey    = []byte{0x02} // key for farm pool reward rule
	FarmerKey          = []byte{0x03} // key for farmer
	FarmPoolExpiredKey = []byte{0x04} // key for expired farm pool
)
