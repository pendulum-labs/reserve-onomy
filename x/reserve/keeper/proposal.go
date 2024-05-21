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
		return sdkerrors.Wrapf(types.ErrMetadataExists, "Metadata already exists for %s", request.Metadata.Base)
	}

	// Need to store initial amount of NOM or other collateral

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	k.SetDenom(ctx, types.Denom{
		Base:     request.Metadata.Base,
		Display:  request.Metadata.Display,
		InitTime: ctx.BlockHeader().Time.Unix(),
	})

	return nil
}

// RegisterCollaterProposal executes the register-collateral proposal.
func (k Keeper) RegisterCollateralProposal(ctx sdk.Context, request *types.RegisterCollateralProposal) error {

	_, found := k.bankKeeper.GetDenomMetaData(ctx, request.Metadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrMetadataExists, "Metadata already exists for %s", request.Metadata.Base)
	}

	// Need to store initial amount of NOM or other collateral

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	k.SetCollateral(ctx, types.Collateral{
		Base:           request.Metadata.Base,
		Display:        request.Metadata.Display,
		MinimumDeposit: request.MinimumDeposit,
		InitTime:       ctx.BlockHeader().Time.Unix(),
	})

	return nil
}
