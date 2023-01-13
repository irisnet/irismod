package simapp

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Network struct {
	*network.Network
	network.Config
}

type ResponseTx struct {
	abci.ResponseDeliverTx
	Height int64
}

func SetupNetwork(t *testing.T) Network {
	cfg := NewConfig()
	cfg.NumValidators = 1

	network, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err, "SetupNetwork failed")

	_, err = network.WaitForHeight(1)
	require.NoError(t, err)
	return Network{
		Network: network,
		Config:  cfg,
	}
}

func (n Network) ExecTxCmdWithResult(t *testing.T,
	clientCtx client.Context,
	cmd *cobra.Command,
	extraArgs []string,
) *ResponseTx {
	buf, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, extraArgs)
	require.NoError(t, err, "ExecTestCLICmd failed")

	n.WaitForNextBlock()

	respType := proto.Message(&sdk.TxResponse{})
	require.NoError(t, clientCtx.Codec.UnmarshalJSON(buf.Bytes(), respType), buf.String())

	txResp := respType.(*sdk.TxResponse)
	require.Equal(t, uint32(0), txResp.Code)

	txHashBz, err := hex.DecodeString(txResp.TxHash)
	require.NoError(t, err, "query tx failed")

	txResult, err := clientCtx.Client.Tx(context.Background(), txHashBz, false)
	require.NoError(t, err, "query tx failed")
	require.Equal(t, uint32(0), txResult.TxResult.Code, fmt.Sprintf("execute %s failed", cmd.Name()))
	return &ResponseTx{txResult.TxResult, txResult.Height}
}

func (n Network) ExecQueryCmd(t *testing.T,
	clientCtx client.Context,
	cmd *cobra.Command,
	extraArgs []string,
	resp proto.Message,
) {
	buf, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, extraArgs)
	require.NoError(t, err, "ExecTestCLICmd failed")
	require.NoError(t, clientCtx.Codec.UnmarshalJSON(buf.Bytes(), resp), buf.String())
}
