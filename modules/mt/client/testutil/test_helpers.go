package testutil

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	mtcli "github.com/irisnet/irismod/modules/mt/client/cli"
	"github.com/irisnet/irismod/simapp"
)

// IssueDenomExec creates a redelegate message.
func IssueDenomExec(
	t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, mtcli.GetCmdIssueDenom(), args)
}

func BurnMTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	mtID string,
	amount string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denomID,
		mtID,
		amount,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, mtcli.GetCmdBurnMT(), args)
}

func MintMTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, mtcli.GetCmdMintMT(), args)
}

func EditMTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	mtID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denomID,
		mtID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, mtcli.GetCmdEditMT(), args)
}

func TransferMTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	recipient string,
	denomID string,
	mtID string,
	amount string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		from,
		recipient,
		denomID,
		mtID,
		amount,
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, mtcli.GetCmdTransferMT(), args)
}

func QueryDenomExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	denomID string,
	resp proto.Message,
	extraArgs ...string,
) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, mtcli.GetCmdQueryDenom(), args, resp)
}

func QueryDenomsExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	resp proto.Message,
	extraArgs ...string,
) {
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, mtcli.GetCmdQueryDenoms(), args, resp)
}

func QueryMTsExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	denomID string,
	resp proto.Message,
	extraArgs ...string,
) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, mtcli.GetCmdQueryMTs(), args, resp)
}

func QueryMTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	denomID string,
	mtID string,
	resp proto.Message,
	extraArgs ...string,
) {
	args := []string{
		denomID,
		mtID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, mtcli.GetCmdQueryMT(), args, resp)
}

func QueryBlancesExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	resp proto.Message,
	extraArgs ...string,
) {
	args := []string{
		from,
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, mtcli.GetCmdQueryBalances(), args, resp)
}

func TransferDenomExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	recipient string,
	denomID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		from,
		recipient,
		denomID,
	}

	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, mtcli.GetCmdTransferDenom(), args)
}
