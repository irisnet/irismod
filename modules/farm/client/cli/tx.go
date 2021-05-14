package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/irisnet/irismod/modules/farm/types"
)

// NewTxCmd returns the transaction commands for the farm module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Record transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdCreateFarmPool(),
		GetCmdDestroyFarmPool(),
		GetCmdAppendReward(),
		GetCmdStake(),
		GetCmdUnstake(),
		GetCmdHarvest(),
	)
	return txCmd
}

// GetCmdCreateFarmPool implements the create a new farm pool command.
func GetCmdCreateFarmPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iris tx farm create <Farm Pool Name> [flags]",
		Short: "Create a new farm pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}
	cmd.Flags().AddFlagSet(FsCreateFarmPool)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdDestroyFarmPool implements the destroy a farm pool command.
func GetCmdDestroyFarmPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iris tx farm destroy <Farm Pool Name> [flags]",
		Short: "Destroy a new farm pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdAppendReward implements the append some reward for farm pool command.
func GetCmdAppendReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iris tx farm append <Farm Pool Name> <reward> [flags]",
		Short: "Append some reward for farm pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdStake implements the staking lp token to farm pool command.
func GetCmdStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iris tx farm stake <Farm Pool Name> <lp token> [flags]",
		Short: "Stake some lp token to farm pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdUnstake implements the unstaking some lp token from farm pool command.
func GetCmdUnstake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iris tx farm unstake <Farm Pool Name> <lp token> [flags]",
		Short: "Unstake some lp token from farm pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdHarvest implements the withdrawing some reward from the farm pool.
func GetCmdHarvest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iris tx farm harvest <Farm Pool Name>",
		Short: "withdraw some reward from the farm pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//TODO
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
