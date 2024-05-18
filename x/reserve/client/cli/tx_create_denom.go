package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"reserve/x/reserve/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type proposalFlags struct {
	Title        string
	Description  string
	Deposit      string
	MetadataPath string
}

// CmdCreateDenomProposal implements the command to submit a create-denom proposal.
func CmdCreateDenomProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom rate",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a create denom proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a create denom proposal.
Example:
$ %s tx gov submit-proposal create-denom rate --title="Test Proposal" --description="My awesome proposal" --deposit="10000000000000000000aonex" --metadata-path "./metadata.json"`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			rateString := args[0]
			rateStringSplit := strings.Split(rateString, ",")

			rateNumerator, err := sdk.ParseUint(rateStringSplit[0])
			if err != nil {
				return err
			}

			rateDenominator, err := sdk.ParseUint(rateStringSplit[1])
			if err != nil {
				return err
			}

			rate := []sdk.Uint{rateNumerator, rateDenominator}

			proposalFlags, err := parseSubmitCreateDenomProposalFlags(cmd.Flags())
			if err != nil {
				return err
			}

			path := proposalFlags.MetadataPath

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

func parseSubmitCreateDenomProposalFlags(fs *pflag.FlagSet) (*proposalFlags, error) {
	title, err := fs.GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}
	description, err := fs.GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	deposit, err := fs.GetString(govcli.FlagDeposit)
	if err != nil {
		return nil, err
	}

	path, err := fs.GetString("metadata-path")
	if err != nil {
		return nil, err
	}

	return &proposalFlags{
		Title:        title,
		Description:  description,
		Deposit:      deposit,
		MetadataPath: path,
	}, nil
}

func addProposalFlags(cmd *cobra.Command) {
	cmd.Flags().String(govcli.FlagTitle, "", "The proposal title")
	cmd.Flags().String(govcli.FlagDescription, "", "The proposal description")
	cmd.Flags().String(govcli.FlagDeposit, "", "The proposal deposit")
	cmd.Flags().String("metadata-path", "", "The path to metadata json")
}
