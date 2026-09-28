package e2e

import (
	"testing"

	cmtconfig "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/store"
	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/require"

	"mods.irisnet.org/simapp"
)

func TestNetworkBlockStoreLateRead(t *testing.T) {
	cfg, err := simapp.NewConfig(defaultDepinjectOptions())
	require.NoError(t, err)
	val := network.Validator{
		Ctx:       &server.Context{Config: cmtconfig.DefaultConfig().SetRoot(t.TempDir())},
		AppConfig: serverconfig.DefaultConfig(),
	}
	// SDK network.startInProcess constructs the app before opening CometBFT databases.
	app := cfg.AppConstructor(val)
	t.Cleanup(func() { require.NoError(t, app.Close()) })
	db, err := cmtconfig.DefaultDBProvider(&cmtconfig.DBContext{
		ID: "blockstore", Config: val.Ctx.Config,
	})
	require.NoError(t, err)
	blockStore := store.NewBlockStore(db)
	t.Cleanup(func() { _ = blockStore.Close() })
	commit := &cmttypes.Commit{
		Height:     7,
		BlockID:    cmttypes.BlockID{Hash: tmhash.Sum([]byte("block"))},
		Signatures: []cmttypes.CommitSig{cmttypes.NewCommitSigAbsent()},
	}
	require.NoError(t, blockStore.SaveSeenCommit(commit.Height, commit))
	require.NoError(t, blockStore.Close())

	// Reproduce the shutdown interleaving: queryMaj23Routine can read a commit
	// after Node.OnStop closes the store. Do not read first and populate its cache.
	require.NotPanics(t, func() {
		require.Equal(t, commit, blockStore.LoadSeenCommit(commit.Height))
	})
}
