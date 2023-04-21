package cli

import (
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
			return nil
		},
	}
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
			return nil
		},
	}

	return cmd
}
