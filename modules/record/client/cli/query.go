package cli

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/irisnet/irismod/modules/record/types"
)

// GetQueryCmd returns the cli query commands for the record module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the record module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdQueryRecord(),
	)
	return queryCmd
}

// GetCmdQueryRecord implements the query record command.
func GetCmdQueryRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record [record-id]",
		Short: "Query a record",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			recordID, err := hex.DecodeString(args[0])
			if err != nil {
				return errors.New("invalid record id, must be hex encoded string")
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Record(context.Background(), &types.QueryRecordRequest{
				RecordId: recordID,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Record)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
