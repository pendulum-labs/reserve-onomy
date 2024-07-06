package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeCreateDenomProposal defines the type for a CreateDenomProposal.
	ProposalTypeCreateDenomProposal = "CreateDenomProposal"
	// ProposalTypeRegisterCollateralProposal defines the type for a RegisterCollateralProposal.
	ProposalTypeRegisterCollateralProposal = "RegisterCollateralProposal"
)

var (
	_ govtypes.Content = &CreateDenomProposal{}
	_ govtypes.Content = &RegisterCollateralProposal{}
)

func init() { // nolint:gochecknoinits // cosmos sdk style
	govtypes.RegisterProposalType(ProposalTypeCreateDenomProposal)
	govtypes.RegisterProposalTypeCodec(&CreateDenomProposal{}, fmt.Sprintf("%s/%s", ModuleName, ProposalTypeCreateDenomProposal))
}

// NewCreateDenomProposal creates a new create-denom proposal.
func NewCreateDenomProposal(sender sdk.AccAddress, title string, description string, metadata banktypes.Metadata, rate []sdk.Uint, collateralDeposit sdk.Coin, debtInterestRate uint64, bondInterestRate uint64) *CreateDenomProposal {
	return &CreateDenomProposal{sender.String(), title, description, &metadata, rate, collateralDeposit, debtInterestRate, bondInterestRate}
}

// GetTitle returns the title of a create-denom proposal.
func (m *CreateDenomProposal) GetTitle() string { return m.Title }

// GetDescription returns the description of a create-denom proposal.
func (m *CreateDenomProposal) GetDescription() string { return m.Description }

// GetProposer returns the proposer from the create-denom proposal struct.
func (m *CreateDenomProposal) GetProposer() string { return m.Sender }

// ProposalRoute returns the routing key of a create-denom proposal.
func (m *CreateDenomProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the create-denom proposal.
func (m *CreateDenomProposal) ProposalType() string { return ProposalTypeCreateDenomProposal }

// ValidateBasic runs basic stateless validity checks.
func (m *CreateDenomProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}
	if err := sdk.VerifyAddressFormat(sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", err)
	}

	return nil
}

// NewRegisterCollateralProposal creates a new create-denom proposal.
func NewRegisterCollateralProposal(sender sdk.AccAddress, title string, description string, metadata banktypes.Metadata, minCollateralDeposit sdk.Coin, LendingRatio uint64, liquidationRatio uint64) *RegisterCollateralProposal {
	return &RegisterCollateralProposal{sender.String(), title, description, &metadata, minCollateralDeposit, LendingRatio, liquidationRatio}
}

// GetTitle returns the title of a register-collateral proposal.
func (m *RegisterCollateralProposal) GetTitle() string { return m.Title }

// GetDescription returns the description of a register-collateral proposal.
func (m *RegisterCollateralProposal) GetDescription() string { return m.Description }

// GetProposer returns the proposer from the register-collateral struct.
func (m *RegisterCollateralProposal) GetProposer() string { return m.Sender }

// ProposalRoute returns the routing key of a register-collateral proposal.
func (m *RegisterCollateralProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the register-collateral proposal.
func (m *RegisterCollateralProposal) ProposalType() string {
	return ProposalTypeRegisterCollateralProposal
}

// ValidateBasic runs basic stateless validity checks.
func (m *RegisterCollateralProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}
	if err := sdk.VerifyAddressFormat(sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", err)
	}

	return nil
}
