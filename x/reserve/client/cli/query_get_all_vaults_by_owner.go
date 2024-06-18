package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"reserve/x/reserve/types"
)

var _ = strconv.Itoa(0)

func CmdGetAllVaultsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-vaults-by-owner",
		Short: "Query get_all_vaults_by_owner",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAllVaultsByOwnerRequest{}

			res, err := queryClient.GetAllVaultsByOwner(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
