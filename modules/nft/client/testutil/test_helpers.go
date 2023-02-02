package testutil

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	nftcli "github.com/irisnet/irismod/modules/nft/client/cli"
	"github.com/irisnet/irismod/simapp"
)

// IssueDenomExec creates a redelegate message.
func IssueDenomExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denom string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denom,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, nftcli.GetCmdIssueDenom(), args)
}

func BurnNFTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	tokenID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, nftcli.GetCmdBurnNFT(), args)
}

func MintNFTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	tokenID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, nftcli.GetCmdMintNFT(), args)
}

func EditNFTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	denomID string,
	tokenID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, nftcli.GetCmdEditNFT(), args)
}

func TransferNFTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	from string,
	recipient string,
	denomID string,
	tokenID string,
	extraArgs ...string,
) *simapp.ResponseTx {
	args := []string{
		recipient,
		denomID,
		tokenID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}
	args = append(args, extraArgs...)

	return network.ExecTxCmdWithResult(t, clientCtx, nftcli.GetCmdTransferNFT(), args)
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
		recipient,
		denomID,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}

	args = append(args, extraArgs...)
	return network.ExecTxCmdWithResult(t, clientCtx, nftcli.GetCmdTransferDenom(), args)
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

	network.ExecQueryCmd(t, clientCtx, nftcli.GetCmdQueryDenom(), args, resp)
}

func QueryCollectionExec(t *testing.T,
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

	network.ExecQueryCmd(t, clientCtx, nftcli.GetCmdQueryCollection(), args, resp)
}

func QueryDenomsExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	resp proto.Message,
	extraArgs ...string) {
	args := []string{
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, nftcli.GetCmdQueryDenoms(), args, resp)
}

func QuerySupplyExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	denom string,
	resp proto.Message,
	extraArgs ...string) {
	args := []string{
		denom,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, nftcli.GetCmdQuerySupply(), args, resp)
}

func QueryOwnerExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	address string,
	resp proto.Message,
	extraArgs ...string) {
	args := []string{
		address,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, nftcli.GetCmdQueryOwner(), args, resp)
}

func QueryNFTExec(t *testing.T,
	network simapp.Network,
	clientCtx client.Context,
	denomID string,
	tokenID string,
	resp proto.Message,
	extraArgs ...string) {
	args := []string{
		denomID,
		tokenID,
		fmt.Sprintf("--%s=json", cli.OutputFlag),
	}
	args = append(args, extraArgs...)

	network.ExecQueryCmd(t, clientCtx, nftcli.GetCmdQueryNFT(), args, resp)
}
