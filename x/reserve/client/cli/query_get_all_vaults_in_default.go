package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"reserve/x/reserve/types"
)

var _ = strconv.Itoa(0)

func CmdGetAllVaultsInDefault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-vaults-in-default",
		Short: "Query get_all_vaults_in_default",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAllVaultsInDefaultRequest{}

			res, err := queryClient.GetAllVaultsInDefault(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
