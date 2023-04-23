package cli

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	etherminttypes "github.com/evmos/ethermint/types"

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
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress()

			receiver := cliCtx.GetFromAddress()
			if len(args) == 3 {
				receiver, err = sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}
			}

			msg := types.NewConvertNFTMsg(args[0], args[1], sender, receiver)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewConvertERC721Cmd returns a CLI command handler for creating a MsgConvertERC721 transaction
func NewConvertERC721Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-erc721 [contract_address] [token-id] [recipient]",
		Short: "Converts an ERC721 token to an NFT",
		Long:  "Converts an ERC721 token to an NFT",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contract := args[0]
			if err := etherminttypes.ValidateAddress(contract); err != nil {
				return fmt.Errorf("invalid ERC721 contract address %w", err)
			}

			tokenId, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount %s", args[1])
			}

			from := common.BytesToAddress(cliCtx.GetFromAddress().Bytes())
			receiver := cliCtx.GetFromAddress()
			if len(args) == 3 {
				receiver, err = sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}
			}

			msg := &types.MsgConvertERC721{
				ContractAddress: contract,
				TokenId:         tokenId,
				Sender:          from.Hex(),
				Receiver:        receiver.String(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
