package record

import (
	"fmt"
	"testing"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/gogoproto/proto"

	recordcli "mods.irisnet.org/modules/record/client/cli"
	"mods.irisnet.org/simapp"
)

// CreateRecordExec creates a record
//
// Parameters:
// - t: *testing.T
// - network: simapp.Network
// - clientCtx: client.Context
// - from: string
// - digest: string
// - digestAlgo: string
// - extraArgs: ...string
// Returns *simapp.ResponseTx
func CreateRecordExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	digest string,
	digestAlgo string,
	extraArgs ...string,
) *simapp.ResponseTx {
	t.Helper()
	args := []string{
		digest,
		digestAlgo,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, recordcli.GetCmdCreateRecord(), args)
}

// QueryRecordExec queries a record.
//
// Parameters:
// - t: *testing.T
// - network: simapp.Network
// - clientCtx: client.Context
// - recordID: string
// - resp: proto.Message
// - extraArgs: ...string
func QueryRecordExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	recordID string,
	resp proto.Message,
	extraArgs ...string,
) {
	t.Helper()
	args := []string{
		recordID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, recordcli.GetCmdQueryRecord(), args, resp)
}
