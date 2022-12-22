package testutil

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/irisnet/irismod/modules/nft/client/cli"
)

func RentalSetUser(clientCtx client.Context, from string, denom, token string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denom,
		token,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetCmdRentalSetUser(), args)
}

func QueryRentalUserOf(clientCtx client.Context, denom, token string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denom,
		token,
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetCmdQueryRentalUserOf(), args)
}

func QueryRentalUserExpires(clientCtx client.Context, denom, token string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denom,
		token,
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetCmdQueryRentalUserExpires(), args)
}

func QueryRentalHasUser(clientCtx client.Context, denom, token string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		denom,
		token,
	}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.GetCmdQueryRentalHasUser(), args)
}