package cli

import (
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
		GetCmdQueryReward(),
	)
	return queryCmd
}

// GetCmdQueryFarmPools implements the query the farm pool.
func GetCmdQueryFarmPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "farm pools --pool-name <Farm Pool Name>",
		Short: "Query a farm",
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}

	cmd.Flags().AddFlagSet(FsQueryFarmPool)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryReward implements the query the farmer reward.
func GetCmdQueryReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "farm reward <Farmer Address> --pool-name <Farm Pool Name>",
		Short: "Query farmer reward",
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}

	cmd.Flags().AddFlagSet(FsQueryFarmPool)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
