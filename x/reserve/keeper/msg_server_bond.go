package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Bond(goCtx context.Context, msg *types.MsgBond) (*types.MsgBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, err := sdk.ParseCoinNormalized(msg.Denom)
	if err != nil {
		return nil, err
	}

	_, found := k.GetDenom(ctx, denom.Denom)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom not found")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, sdk.NewCoins(denom))
	if err != nil {
		return nil, err
	}

	totalSupply := k.bankKeeper.GetSupply(ctx, denom.Denom)
	if totalSupply.Equal(sdk.ZeroInt()) {
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins())
	}

	return &types.MsgBondResponse{}, nil
}
