package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"reserve/x/reserve/types"
)

var _ = strconv.Itoa(0)

func CmdGetVaultByUid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-vault-by-uid",
		Short: "Query get_vault_by_uid",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetVaultByUidRequest{}

			res, err := queryClient.GetVaultByUid(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
