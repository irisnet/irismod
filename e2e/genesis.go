package e2e

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tokentypes "mods.irisnet.org/modules/token/types"
	tokenv1 "mods.irisnet.org/modules/token/types/v1"
)

const testValidatorMnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

// AddTestTokenToGenesis adds a token owned by the first validator to test genesis.
func AddTestTokenToGenesis(t *testing.T, cfg *network.Config, token tokenv1.Token) {
	t.Helper()

	cfg.Mnemonics = []string{testValidatorMnemonic}
	record, err := keyring.NewInMemory(cfg.Codec).NewAccount(
		"validator",
		testValidatorMnemonic,
		keyring.DefaultBIP39Passphrase,
		sdk.GetConfig().GetFullBIP44Path(),
		hd.Secp256k1,
	)
	require.NoError(t, err)
	owner, err := record.GetAddress()
	require.NoError(t, err)
	token.Owner = owner.String()

	var genesis tokenv1.GenesisState
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[tokentypes.ModuleName], &genesis)
	genesis.Tokens = append(genesis.Tokens, token)
	cfg.GenesisState[tokentypes.ModuleName] = cfg.Codec.MustMarshalJSON(&genesis)
}
