package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/irisnet/irismod/modules/oracle/types"
)

// GetQueryCmd returns the cli query commands for the oracle module.
func GetQueryCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdQueryFeed(),
		GetCmdQueryFeeds(),
		GetCmdQueryFeedValue(),
	)
	return txCmd
}

// GetCmdQueryFeed implements the query feed Content definition command
func GetCmdQueryFeed() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "feed [feed-name]",
		Short:   "Query the feed definition",
		Example: fmt.Sprintf("%s q oracle query-feed <feed-name>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Feed(context.Background(), &types.QueryFeedRequest{FeedName: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(&res.Feed)
		},
	}
	cmd.Flags().AddFlagSet(FsQueryFeed)
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryFeed implements the query feed Content definition command
func GetCmdQueryFeeds() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "feeds",
		Short:   "Query a group of feed definition",
		Example: fmt.Sprintf("%s q oracle query-feeds", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Feeds(context.Background(), &types.QueryFeedsRequest{State: viper.GetString(FlagFeedState)})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	cmd.Flags().AddFlagSet(FsQueryFeeds)
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryFeedValue implements the query feed value command
func GetCmdQueryFeedValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "value [feed-name]",
		Short:   "Query the feed result",
		Example: fmt.Sprintf("%s q oracle query-value <feed-name>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.FeedValue(context.Background(), &types.QueryFeedValueRequest{FeedName: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	cmd.Flags().AddFlagSet(FsQueryFeedValue)
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
