package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/irisnet/irismod/modules/farm/types"
)

// GetQueryCmd returns the cli query commands for the farm module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the farm module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdQueryFarmPools(),
		GetCmdQueryFarmer(),
	)
	return queryCmd
}

// GetCmdQueryFarmPools implements the query the farm pool.
func GetCmdQueryFarmPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "farm pools --pool-name <Farm Pool Name>",
		Short: "Query a farm",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Pools(context.Background(), &types.QueryPoolsRequest{
				Name: args[0],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().AddFlagSet(FsQueryFarmPool)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryFarmer implements the query the farmer reward.
func GetCmdQueryFarmer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "farm farmer <Farmer Address> --pool-name <Farm Pool Name>",
		Short: "Query farmer reward",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolName, err := cmd.Flags().GetString(FlagFarmPool)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Farmer(context.Background(), &types.QueryFarmerRequest{
				Farmer:   args[0],
				PoolName: poolName,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().AddFlagSet(FsQueryFarmPool)
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
