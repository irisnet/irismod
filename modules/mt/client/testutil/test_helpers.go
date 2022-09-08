package testutil

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	mtcli "github.com/irisnet/irismod/modules/mt/client/cli"
)

// IssueDenomExec creates a redelegate message.
func IssueDenomExec(clientCtx client.Context, from string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdIssueDenom(), args)
}

func BurnMTExec(clientCtx client.Context, from string, denomID string, mtID string, amount string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		mtID,
		amount,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdBurnMT(), args)
}

func MintMTExec(clientCtx client.Context, from string, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdMintMT(), args)
}

func EditMTExec(clientCtx client.Context, from string, denomID string, mtID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		mtID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdEditMT(), args)
}

func TransferMTExec(clientCtx client.Context, from string, recipient string, denomID string, mtID string, amount string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		from,
		recipient,
		denomID,
		mtID,
		amount,
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdTransferMT(), args)
}

func QueryDenomExec(clientCtx client.Context, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdQueryDenom(), args)
}

func QueryDenomsExec(clientCtx client.Context, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdQueryDenoms(), args)
}

func QueryMTsExec(clientCtx client.Context, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdQueryMTs(), args)
}

func QueryMTExec(clientCtx client.Context, denomID string, mtID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denomID,
		mtID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdQueryMT(), args)
}

func QueryBlancesExec(clientCtx client.Context, from string, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		from,
		denomID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdQueryBalances(), args)
}

func TransferDenomExec(clientCtx client.Context, from string, recipient string, denomID string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		from,
		recipient,
		denomID,
	}

	args = append(args, extraArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, mtcli.GetCmdTransferDenom(), args)
}
