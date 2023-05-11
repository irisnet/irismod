package cli

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/irisnet/irismod/utils"
	"github.com/tjfoc/gmsm/sm2"
	"strings"

	"github.com/spf13/cobra"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/irisnet/irismod/modules/record/types"
)

// GetQueryCmd returns the cli query commands for the record module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the record module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdQueryRecord(),
	)
	return queryCmd
}

// GetCmdQueryRecord implements the query record command.
func GetCmdQueryRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record [record-id]",
		Short: "Query a record",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
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

			decrypt, err := cmd.Flags().GetBool(Decrypt)
			if err != nil {
				return err
			}
			encryption, err := cmd.Flags().GetString(Encryption)
			if err != nil {
				return err
			}
			if decrypt {
				datas := strings.Split(res.Record.Contents[0].Meta, ",")
				if len(datas) != 2 {
					return errors.New("failed to decrypt")
				}
				from, err := cmd.Flags().GetString(From)
				if err != nil {
					return err
				}

				ks := keyring.NewUnsafe(clientCtx.Keyring)
				priv, err := ks.UnsafeExportPrivKeyHex(from)

				privateKey, err := utils.GetPrivateFromHex(priv)
				if err != nil {
					return err
				}
				fromInfo, err := clientCtx.Keyring.Key(from)
				address := fromInfo.GetAddress().String()
				var key []byte
				//判断是否是拥有着获取aesKey
				if address == res.Record.Creator {
					sm2AesKey, err := hex.DecodeString(datas[0])
					if err != nil {
						return err
					}
					key, err = sm2.Decrypt(privateKey, sm2AesKey)
					if err != nil {
						return err
					}
				} else {
					str := sm2.Compress(&privateKey.PublicKey)
					grantRecordID := args[0] + hex.EncodeToString(str)
					grantRes, err := queryClient.GrantRecord(
						context.Background(),
						&types.QueryGrantRecordRequest{GrantRecordId: tmbytes.HexBytes(grantRecordID).String()},
					)
					if err != nil {
						return err
					}
					if grantRes.GrantRecord.Key == "" {
						return errors.New("failed to decrypt")
					}
					sm2AesKey, err := hex.DecodeString(grantRes.GrantRecord.Key)
					if err != nil {
						return err
					}

					key, err = sm2.Decrypt(privateKey, sm2AesKey)
					if err != nil {
						return err
					}
				}

				for k, v := range res.Record.Contents {
					datas := strings.Split(v.Meta, ",")
					if len(datas) != 2 {
						return errors.New("failed to decrypt")
					}
					str, err := utils.SymmetricDecrypt(datas[1], key, encryption)
					if err != nil {
						return err
					}
					res.Record.Contents[k].Meta = string(str)
				}

			}

			return clientCtx.PrintProto(res.Record)
		},
	}
	cmd.Flags().AddFlagSet(FsQureyRecord)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
