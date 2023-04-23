package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all erc20 CLI query commands
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the erc20 module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetTokenPairsCmd(),
		GetTokenPairCmd(),
	)

	return cmd
}

// GetTokenPairsCmd queries all registered token pairs
func GetTokenPairsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-pairs",
		Short: "Gets registered token pairs",
		Long:  "Gets registered token pairs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clienCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clienCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryTokenPairsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TokenPairs(context.Background(), req)
			if err != nil {
				return err
			}

			return clienCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetTokenPairCmd queries a registered token pair
func GetTokenPairCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-pair TOKEN",
		Short: "Get a registered token pair",
		Long:  "Get a registered token pair",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryTokenPairRequest{
				Token: args[0],
			}

			res, err := queryClient.TokenPair(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
