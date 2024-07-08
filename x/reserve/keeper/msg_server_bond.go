package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Bond(goCtx context.Context, msg *types.MsgBond) (*types.MsgBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denoms, err := sdk.ParseCoinNormalized(msg.Denom)
	if err != nil {
		return nil, err
	}

	denom, found := k.GetDenom(ctx, denoms.Denom)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom not found")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, sdk.NewCoins(denom))
	if err != nil {
		return nil, err
	}

	bondShares := (denom.BondShares.Mul(denoms.Amount)).Quo(denom.BondDenoms)

	bondName := denoms.Denom + "_bond"

	bondCoin := sdk.NewCoin(bondName, bondShares)

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(bondCoin))
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, sdk.NewCoins(bondCoin))
	if err != nil {
		return nil, err
	}

	return &types.MsgBondResponse{}, nil
}
