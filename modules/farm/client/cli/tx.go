package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

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
		Example: fmt.Sprintf("$ %s tx farm create <Farm Pool Name> [flags]", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			description, _ := cmd.Flags().GetString(FlagDescription)
			lpTokenDenom, _ := cmd.Flags().GetString(FlagLPTokenDenom)
			startHeight, err := cmd.Flags().GetInt64(FlagStartHeight)
			if err != nil {
				return err
			}
			editable, _ := cmd.Flags().GetBool(FlagEditable)

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
				StartHeight:    uint64(startHeight),
				RewardPerBlock: rewardPerBlock,
				TotalReward:    totalReward,
				Editable:       editable,
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
		Example: fmt.Sprintf("$ %s tx farm destroy <Farm Pool Name> [flags]", version.AppName),
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
		Use:     "adjust",
		Short:   "Adjust farm pool parameters",
		Example: fmt.Sprintf("$ %s tx farm adjust <pool-name> [flags]", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var additionalReward, rewardPerBlock sdk.Coins
			if cmd.Flags().Changed(FlagRewardPerBlock) {
				rewardPerBlockStr, _ := cmd.Flags().GetString(FlagRewardPerBlock)
				rewardPerBlock, err = sdk.ParseCoinsNormalized(rewardPerBlockStr)
				if err != nil {
					return err
				}
			}
			if cmd.Flags().Changed(FlagAdditionalReward) {
				additionalRewardStr, _ := cmd.Flags().GetString(FlagAdditionalReward)
				additionalReward, err = sdk.ParseCoinsNormalized(additionalRewardStr)
				if err != nil {
					return err
				}
			}
			msg := types.MsgAdjustPool{
				PoolName:         args[0],
				AdditionalReward: additionalReward,
				RewardPerBlock:   rewardPerBlock,
				Creator:          clientCtx.GetFromAddress().String(),
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
		Example: fmt.Sprintf("$ %s tx farm stake <Farm Pool Name> <lp token> [flags]", version.AppName),
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
		Example: fmt.Sprintf("$ %s tx farm unstake <Farm Pool Name> <lp token> [flags]", version.AppName),
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
		Example: fmt.Sprintf("$ %s tx farm harvest <Farm Pool Name>", version.AppName),
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
