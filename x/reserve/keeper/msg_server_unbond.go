package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Unbond(goCtx context.Context, msg *types.MsgUnbond) (*types.MsgUnbondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bonds, err := sdk.ParseCoinNormalized(msg.Bonds)
	if err != nil {
		return nil, err
	}

	bonded, found := k.GetBonded(ctx, bonds.Denom)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bond not found")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, sdk.NewCoins(bonds))
	if err != nil {
		return nil, err
	}

	denom, found := k.GetDenom(ctx, bonded.DenomBase)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom not found")
	}

	denoms := (denom.BondDenoms.Mul(bonds.Amount)).Quo(denom.BondShares)

	denom.BondDenoms = denom.BondDenoms.Sub(denoms)
	denom.BondShares = denom.BondShares.Sub(bonds.Amount)

	k.SetDenom(ctx, denom)

	return &types.MsgUnbondResponse{}, nil
}
