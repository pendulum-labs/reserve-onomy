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

	if request.BondInterestRate > request.DebtInterestRate {
		return sdkerrors.Wrapf(types.ErrBondGtDebt, "Bond Interest Rate %s greater than Debt Interest Rate %s", request.BondInterestRate, request.DebtInterestRate)
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(request.Sender), types.ModuleName, sdk.NewCoins(request.CollateralDeposit))
	if err != nil {
		return err
	}

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	k.SetDenom(ctx, types.Denom{
		Base:             request.Metadata.Base,
		Display:          request.Metadata.Display,
		InitTime:         ctx.BlockHeader().Time.Unix(),
		DebtInterestRate: request.DebtInterestRate,
		BondInterestRate: request.BondInterestRate,
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
