package keeper

import (
	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateDenomProposal executes the create-denom proposal.
func (k Keeper) CreateDenomProposal(ctx sdk.Context, request *types.CreateDenomProposal) error {

	_, found := k.bankKeeper.GetDenomMetaData(ctx, request.DenomMetadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrMetadataExists, "Denom Metadata already exists for %s", request.DenomMetadata.Base)
	}

	_, found = k.bankKeeper.GetDenomMetaData(ctx, request.BondMetadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrMetadataExists, "Bond Metadata already exists for %s", request.BondMetadata.Base)
	}

	if request.BondInterestRate > request.DebtInterestRate {
		return sdkerrors.Wrapf(types.ErrBondGtDebt, "Bond Interest Rate greater than Debt Interest Rate")
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(request.Sender), types.ModuleName, sdk.NewCoins(request.CollateralDeposit))
	if err != nil {
		return err
	}

	k.bankKeeper.SetDenomMetaData(ctx, *request.DenomMetadata)
	k.bankKeeper.SetDenomMetaData(ctx, *request.BondMetadata)

	k.SetDenom(ctx, types.Denom{
		DenomBase:        request.DenomMetadata.Base,
		DenomDisplay:     request.DenomMetadata.Display,
		BondBase:         request.BondMetadata.Base,
		BondDisplay:      request.BondMetadata.Display,
		InitTime:         ctx.BlockHeader().Time.Unix(),
		DebtInterestRate: request.DebtInterestRate,
		BondInterestRate: request.BondInterestRate,
	})

	k.SetBonded(ctx, types.Bonded{
		DenomBase: request.DenomMetadata.Base,
		BondBase:  request.BondMetadata.Base,
	})

	return nil
}

// RegisterCollaterProposal executes the register-collateral proposal.
func (k Keeper) RegisterCollateralProposal(ctx sdk.Context, request *types.RegisterCollateralProposal) error {

	_, found := k.bankKeeper.GetDenomMetaData(ctx, request.Metadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrMetadataExists, "Metadata already exists for %s", request.Metadata.Base)
	}

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	k.SetCollateral(ctx, types.Collateral{
		Base:             request.Metadata.Base,
		Display:          request.Metadata.Display,
		MinimumDeposit:   request.MinimumDeposit,
		LendingRatio:     request.LendingRatio,
		LiquidationRatio: request.LiquidationRatio,
		InitTime:         ctx.BlockHeader().Time.Unix(),
	})

	return nil
}
