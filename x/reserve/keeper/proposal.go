package keeper

import (
	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateDenomProposal executes the create-denom proposal.
func (k Keeper) CreateDenomProposal(ctx sdk.Context, request *types.CreateDenomProposal) error {

	_, found := k.bankKeeper.GetDenomMetaData(ctx, request.Metadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrDenomExists, "Denom: %s already exists", request.Metadata.Base)
	}

	// Need to store initial amount of NOM or other collateral

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	return nil
}

// RegisterCollaterProposal executes the register-collateral proposal.
func (k Keeper) RegisterCollateralProposal(ctx sdk.Context, request *types.RegisterCollateralProposal) error {

	_, found := k.bankKeeper.GetDenomMetaData(ctx, request.Metadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrDenomExists, "Denom: %s already exists", request.Metadata.Base)
	}

	// Need to store initial amount of NOM or other collateral

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	return nil
}
