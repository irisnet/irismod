package testutil

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	tokencli "github.com/irisnet/irismod/modules/token/client/cli"
	tokentypes "github.com/irisnet/irismod/modules/token/types"
	"github.com/irisnet/irismod/simapp"
)

func IssueTokenExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, tokencli.GetCmdIssueToken(), args)
}

func EditTokenExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	symbol string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		symbol,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, tokencli.GetCmdEditToken(), args)
}

func MintTokenExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	symbol string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		symbol,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, tokencli.GetCmdMintToken(), args)
}

func BurnTokenExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	symbol string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		symbol,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, tokencli.GetCmdBurnToken(), args)
}

func TransferTokenOwnerExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	symbol string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		symbol,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, tokencli.GetCmdTransferTokenOwner(), args)
}

func QueryTokenExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	denom string,
	extraArgs ...string,
) tokentypes.TokenI {
	args := []string{
		denom,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	respType := proto.Message(&codectypes.Any{})
	network.ExecQueryCmd(t, clientCtx, tokencli.GetCmdQueryToken(), args, respType)

	var token tokentypes.TokenI
	err := clientCtx.InterfaceRegistry.UnpackAny(respType.(*codectypes.Any), &token)
	require.NoError(t, err, "QueryTokenExec failed")
	return token
}

func QueryTokensExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	owner string,
	extraArgs ...string,
) []tokentypes.TokenI {
	args := []string{
		owner,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)
	tokens := []tokentypes.TokenI{}
	buf, err := clitestutil.ExecTestCLICmd(clientCtx, tokencli.GetCmdQueryTokens(), args)
	require.NoError(t, err, "QueryTokensExec failed")
	require.NoError(t, clientCtx.LegacyAmino.UnmarshalJSON(buf.Bytes(), &tokens))
	return tokens
}

func QueryFeeExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	symbol string,
	extraArgs ...string,
) *tokentypes.QueryFeesResponse {
	args := []string{
		symbol,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	response := &tokentypes.QueryFeesResponse{}
	network.ExecQueryCmd(t, clientCtx, tokencli.GetCmdQueryFee(), args, response)
	return response
}

func QueryParamsExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	extraArgs ...string,
) *tokentypes.Params {
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	response := &tokentypes.Params{}
	network.ExecQueryCmd(t, clientCtx, tokencli.GetCmdQueryParams(), args, response)
	return response
}
