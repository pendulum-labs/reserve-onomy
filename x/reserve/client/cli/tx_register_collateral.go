package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"reserve/x/reserve/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CmdRegisterCollateralProposal implements the command to submit a register-collateral proposal.
func CmdRegisterCollateralProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-collateral metadata-path",
		Args:  cobra.ExactArgs(2),
		Short: "Submit a register collateral proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a register collateral proposal.
Example:
$ %s tx gov submit-proposal register-collateral metadata-path --title="Test Proposal" --description="My awesome proposal" --deposit="10000000000000000000aonex"`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			path := args[1]

			metadataFile, err := os.Open(path)
			if err != nil {
				return err
			}

			byteMetadata, err := io.ReadAll(metadataFile)
			if err != nil {
				return err
			}

			var metadata banktypes.Metadata

			err = json.Unmarshal(byteMetadata, &metadata)
			if err != nil {
				return err
			}

			proposalFlags, err := parseProposalFlags(cmd.Flags())
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(proposalFlags.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewCreateDenomProposal(from, proposalFlags.Title, proposalFlags.Description, metadata, rate)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addProposalFlags(cmd)

	return cmd
}
