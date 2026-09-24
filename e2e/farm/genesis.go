package farm

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/network"

	"mods.irisnet.org/e2e"
	tokentypes "mods.irisnet.org/modules/token/types/v1"
)

func addTestTokenToGenesis(t *testing.T, cfg *network.Config) {
	t.Helper()
	e2e.AddTestTokenToGenesis(t, cfg, tokentypes.Token{
		Symbol:        "kitty",
		Name:          "Kitty Token",
		Scale:         0,
		MinUnit:       "kitty",
		InitialSupply: 100000000,
		MaxSupply:     200000000,
		Mintable:      true,
	})
}
