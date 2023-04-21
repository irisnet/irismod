package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/irisnet/irismod/modules/erc721-converter/types"
	"github.com/spf13/cobra"
)

// NewTxCmd returns a root CLI command handler for erc721 transaction commands
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "erc721 subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConvertNft(),
		NewConvertERC721Cmd(),
		RegisterDenom(),
		RegisterERC721(),
	)

	return txCmd
}

// NewConvertNft returns a CLI command handler for creating a MsgConvertNft transaction
func NewConvertNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-nft [denom-id] [token-id] [recipient]",
		Short: "Converts an NFT to an ERC721 token",
		Long:  "Converts an NFT to an ERC721 token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

// NewConvertERC721Cmd returns a CLI command handler for creating a MsgConvertERC721 transaction
func NewConvertERC721Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-erc721 [denom-id] [token-id] [recipient]",
		Short: "Converts an ERC721 token to an NFT",
		Long:  "Converts an ERC721 token to an NFT",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

// RegisterDenom returns a CLI command handler for creating a MsgRegisterDenom transaction
func RegisterDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-denom [denom-id] [base-uri] [base-token-uri]",
		Short: "Registers a new denom",
		Long:  "Registers a new denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

// RegisterERC721 returns a CLI command handler for creating a MsgRegisterERC721 token transaction
func RegisterERC721() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-erc721 [denom-id] [contract-address] [base-token-uri]",
		Short: "Registers a new ERC721 token",
		Long:  "Registers a new ERC721 token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
