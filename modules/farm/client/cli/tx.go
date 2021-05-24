package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
		Use:     "create",
		Short:   "Create a new farm pool",
		Example: "iris tx farm create <Farm Pool Name> [flags]",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			description, _ := cmd.Flags().GetString(FlagDescription)
			lpTokenDenom, _ := cmd.Flags().GetString(FlagLPTokenDenom)
			beginHeight, err := cmd.Flags().GetInt64(FlagBeginHeight)
			if err != nil {
				return err
			}
			destructible, _ := cmd.Flags().GetBool(FlagDestructible)

			rewardPerBlockStr, _ := cmd.Flags().GetString(FlagRewardPerBlock)
			rewardPerBlock, err := sdk.ParseCoinsNormalized(rewardPerBlockStr)
			if err != nil {
				return err
			}

			totalRewardStr, _ := cmd.Flags().GetString(FlagTotalReward)
			totalReward, err := sdk.ParseCoinsNormalized(totalRewardStr)
			if err != nil {
				return err
			}

			msg := types.MsgCreatePool{
				Name:           args[0],
				Description:    description,
				LpTokenDenom:   lpTokenDenom,
				BeginHeight:    uint64(beginHeight),
				RewardPerBlock: rewardPerBlock,
				TotalReward:    totalReward,
				Destructible:   destructible,
				Creator:        clientCtx.GetFromAddress().String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	cmd.Flags().AddFlagSet(FsCreateFarmPool)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdDestroyFarmPool implements the destroy a farm pool command.
func GetCmdDestroyFarmPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "destroy",
		Short:   "Destroy a new farm pool",
		Example: "iris tx farm destroy <Farm Pool Name> [flags]",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.MsgDestroyPool{
				PoolName: args[0],
				Creator:  clientCtx.GetFromAddress().String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdAppendReward implements the append some reward for farm pool command.
func GetCmdAppendReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "append",
		Short:   "Append some reward for farm pool",
		Example: "iris tx farm append <Farm Pool Name> <reward> [flags]",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reward, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgAppendReward{
				PoolName: args[0],
				Amount:   reward,
				Creator:  clientCtx.GetFromAddress().String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdStake implements the staking lp token to farm pool command.
func GetCmdStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stake",
		Short:   "Stake some lp token to farm pool",
		Example: "iris tx farm stake <Farm Pool Name> <lp token> [flags]",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgStake{
				PoolName: args[0],
				Amount:   amount,
				Sender:   clientCtx.GetFromAddress().String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdUnstake implements the unstaking some lp token from farm pool command.
func GetCmdUnstake() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unstake",
		Short:   "Unstake some lp token from farm pool",
		Example: "iris tx farm unstake <Farm Pool Name> <lp token> [flags]",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgUnstake{
				PoolName: args[0],
				Amount:   amount,
				Sender:   clientCtx.GetFromAddress().String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdHarvest implements the withdrawing some reward from the farm pool.
func GetCmdHarvest() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "harvest",
		Short:   "withdraw some reward from the farm pool",
		Example: "iris tx farm harvest <Farm Pool Name>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgHarvest{
				PoolName: args[0],
				Sender:   clientCtx.GetFromAddress().String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
