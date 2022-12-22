package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/irisnet/irismod/modules/nft/types"
	"github.com/spf13/cobra"
)

// Tx Command for rental plugin

// GetCmdRental is the CLI tx command for NFT rental plugin.
func GetCmdRental() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "rental",
		Long:               "NFT rental plugin tx subcommands",
		DisableFlagParsing: true,
		Run:                func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(
		GetCmdRentalSetUser(),
	)

	return cmd
}

// GetCmdRentalSetUser is the CLI command for setting user for a rentable NFT.
func GetCmdRentalSetUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "set-user [denom-id] [nft-id]",
		Long: "Set user and expiry for an NFT.",
		Example: fmt.Sprintf(
			"$ %s tx nft rental set-user <denom-id> <nft-id> "+
				"--user=<user> "+
				"--expiry=<expiry> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress().String()

			user, err := cmd.Flags().GetString(FlagRentalUser)
			if err != nil {
				return err
			}

			expiry, err := cmd.Flags().GetInt64(FlagRentalExpiry)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetUser(
				args[0],
				args[1],
				user,
				expiry,
				sender,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsRental)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// Query Command for rental plugin

// GetCmdQueryRental is the CLI query command for NFT rental plugin.
func GetCmdQueryRental() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "rental",
		Long:               "NFT rental plugin query subcommands",
		DisableFlagParsing: true,
		Run:                func(cmd *cobra.Command, args []string) {},
	}

	cmd.AddCommand(
		GetCmdQueryRentalUserOf(),
		GetCmdQueryRentalUserExpires(),
		GetCmdQueryRentalHasUser(),
	)

	return cmd
}

// GetCmdQueryRentalUserOf is the CLI command for query user of a rentable NFT.
func GetCmdQueryRentalUserOf() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user [denom-id] [nft-id]",
		Long:    "user of a rented nft",
		Example: fmt.Sprintf("$ %s query nft rental user <denom-id> <nft-id>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.UserOf(context.Background(), &types.QueryUserOfRequest{
				DenomId: args[0],
				NftId:   args[1],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryRentalUserExpires is the CLI command for query expiry of a rentable NFT.
func GetCmdQueryRentalUserExpires() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user [denom-id] [nft-id]",
		Long:    "expiry of a rented nft",
		Example: fmt.Sprintf("$ %s query nft rental user <denom-id> <nft-id>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.UserExpires(context.Background(), &types.QueryUserExpiresRequest{
				DenomId: args[0],
				NftId:   args[1],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryRentalHasUser is the CLI command for query does an NFT have a user.
func GetCmdQueryRentalHasUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user [denom-id] [nft-id]",
		Long:    "does an nft have a user",
		Example: fmt.Sprintf("$ %s query nft rental user <denom-id> <nft-id>", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.HasUser(context.Background(), &types.QueryHasUserRequest{
				DenomId: args[0],
				NftId:   args[1],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
