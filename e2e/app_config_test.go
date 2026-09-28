package e2e

import (
	"testing"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

func TestAuthPreBlockerConfigured(t *testing.T) {
	for _, module := range appConfigDefinition.Modules {
		if module.Name != "runtime" {
			continue
		}

		var config runtimev1alpha1.Module
		require.NoError(t, proto.Unmarshal(module.Config.Value, &config))
		require.Equal(t, []string{upgradetypes.ModuleName, authtypes.ModuleName}, config.PreBlockers)
		return
	}

	require.Fail(t, "runtime module config not found")
}
