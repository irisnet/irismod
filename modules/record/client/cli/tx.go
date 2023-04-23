package cli

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/irisnet/irismod/utils"
	"github.com/spf13/cobra"
	"github.com/tjfoc/gmsm/x509"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/irisnet/irismod/modules/record/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tjfoc/gmsm/sm2"
)

// NewTxCmd returns the transaction commands for the record module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Record transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdCreateRecord(),
		GetCmdGrantRecord(),
		GetCmdVerifyRecord(),
	)
	return txCmd
}

// GetCmdCreateRecord implements the create record command.
func GetCmdCreateRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [digest] [digest-algo]",
		Short: "Create a new record",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress().String()

			uri, err := cmd.Flags().GetString(FlagURI)
			if err != nil {
				return err
			}
			meta, err := cmd.Flags().GetString(FlagMeta)
			if err != nil {
				return err
			}
			encrypt, err := cmd.Flags().GetBool(Encrypt)
			if err != nil {
				return err
			}
			if encrypt {
				key := utils.RandStr(16)
				data, err := utils.AesEncrypt([]byte(meta), key)
				if err != nil {
					return err
				}

				fromInfo, err := clientCtx.Keyring.Key(clientCtx.GetFromName())
				if err != nil {
					return err
				}

				keyData, err := sm2.Encrypt(sm2.Decompress(fromInfo.GetPubKey().Bytes()), key, rand.Reader, 0)
				if err != nil {
					return err
				}

				meta = hex.EncodeToString(keyData) + "," + data
			}

			content := types.Content{
				Digest:     args[0],
				DigestAlgo: args[1],
				URI:        uri,
				Meta:       meta,
			}

			msg := types.NewMsgCreateRecord([]types.Content{content}, from)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsCreateRecord)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdGrantRecord implements the create record command.
func GetCmdGrantRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [record-id] [pubKey]",
		Short: "grant pubKey record",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress().String()
			recordID, err := hex.DecodeString(args[0])
			if err != nil {
				return errors.New("invalid record id, must be hex encoded string")
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Record(
				context.Background(),
				&types.QueryRecordRequest{RecordId: tmbytes.HexBytes(recordID).String()},
			)
			if err != nil {
				return err
			}

			if res.GetRecord().Creator != from {
				return errors.New("this record the user does not have permission to grant")
			}

			datas := strings.Split(res.Record.Contents[0].Meta, ",")
			if len(datas) != 2 {
				return errors.New("this record not require grant")
			}

			ks := keyring.NewUnsafe(clientCtx.Keyring)
			priv, err := ks.UnsafeExportPrivKeyHex(clientCtx.GetFromName())

			privateKey, err := x509.ReadPrivateKeyFromHex(priv)
			if err != nil {
				return err
			}
			sm2AesKey, err := hex.DecodeString(datas[0])
			if err != nil {
				return err
			}

			aesKey, err := sm2.Decrypt(privateKey, []byte(sm2AesKey), 0)
			if err != nil {
				return err
			}

			pubkeyHex, err := hex.DecodeString(args[1])
			if err != nil {
				return err
			}

			pubkey := sm2.Decompress(pubkeyHex)

			keyData, err := sm2.Encrypt(pubkey, aesKey, rand.Reader, 0)
			if err != nil {
				return err
			}
			keyStr := hex.EncodeToString(keyData)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), types.NewMsgGrantRecord(args[0], args[1], keyStr, from))
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdVerifyRecord implements the create record command.
func GetCmdVerifyRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify [record-id] [fake-hash]",
		Short: "verify record hash",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recordID, err := hex.DecodeString(args[0])
			if err != nil {
				return errors.New("invalid record id, must be hex encoded string")
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Record(
				context.Background(),
				&types.QueryRecordRequest{RecordId: tmbytes.HexBytes(recordID).String()},
			)
			if err != nil {
				return err
			}

			if res.Record.Contents[0].Digest != args[1] {
				return errors.New("illegal record hash")
			}

			return clientCtx.PrintProto(res.Record)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
