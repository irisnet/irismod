package farm

import (
	"fmt"
	"testing"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	farmcli "mods.irisnet.org/modules/farm/client/cli"
	farmtypes "mods.irisnet.org/modules/farm/types"
	"mods.irisnet.org/simapp"
)

// CreateFarmPoolExec creates a redelegate message.
func CreateFarmPoolExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, creator),
	}
	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, farmcli.GetCmdCreateFarmPool(), args)
}

// QueryFarmPoolsExec queries farm pools
func QueryFarmPoolsExec(
	t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	extraArgs ...string,
) *farmtypes.QueryFarmPoolsResponse {
	t.Helper()
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	response := &farmtypes.QueryFarmPoolsResponse{}
	network.ExecQueryCmd(t, clientCtx, farmcli.GetCmdQueryFarmPools(), args, response)
	return response
}

// QueryFarmPoolExec queries farm pool
func QueryFarmPoolExec(
	t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	poolID string,
	extraArgs ...string,
) *farmtypes.QueryFarmPoolResponse {
	t.Helper()
	args := []string{
		poolID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)
	response := &farmtypes.QueryFarmPoolResponse{}
	network.ExecQueryCmd(t, clientCtx, farmcli.GetCmdQueryFarmPool(), args, response)
	return response
}

// AppendRewardExec creates a redelegate message.
func AppendRewardExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator,
	poolID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		poolID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, creator),
	}
	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, farmcli.GetCmdAdjustPool(), args)
}

// StakeExec creates a redelegate message.
func StakeExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator,
	poolID,
	lpToken string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		poolID,
		lpToken,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, creator),
	}
	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, farmcli.GetCmdStake(), args)
}

// UnstakeExec creates a redelegate message.
func UnstakeExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator,
	poolID,
	lpToken string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		poolID,
		lpToken,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, creator),
	}
	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, farmcli.GetCmdUnstake(), args)
}

// HarvestExec creates a redelegate message.
func HarvestExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator,
	poolID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		poolID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, creator),
	}
	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, farmcli.GetCmdHarvest(), args)
}

// DestroyExec creates a redelegate message.
func DestroyExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator,
	poolID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		poolID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, creator),
	}
	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, farmcli.GetCmdDestroyFarmPool(), args)
}

// QueryFarmerExec creates a redelegate message.
func QueryFarmerExec(t *testing.T, network simapp.Network, clientCtx client.Context,
	creator string,
	extraArgs ...string,
) *farmtypes.QueryFarmerResponse {
	t.Helper()
	args := []string{
		creator,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)
	response := &farmtypes.QueryFarmerResponse{}
	network.ExecQueryCmd(t, clientCtx, farmcli.GetCmdQueryFarmer(), args, response)
	return response
}
