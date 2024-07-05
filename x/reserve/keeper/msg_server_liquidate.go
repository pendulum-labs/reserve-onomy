package keeper

import (
	"context"
	"strconv"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Liquidate(goCtx context.Context, msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	liquidator, _ := sdk.AccAddressFromBech32(msg.Creator)

	uid, err := strconv.ParseUint(msg.Uid, 10, 64)
	if err != nil {
		return nil, err
	}

	vault, found := k.GetVault(ctx, uid)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "vault not found")
	}

	collateral, found := k.GetCollateral(ctx, vault.Collateral.Denom)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral not found")
	}

	denom, found := k.GetDenom(ctx, vault.DebtDenom)
	if !found {
		sdkerrors.Wrapf(types.ErrDenomNotFound, "Denom with name %s not found", vault.DebtDenom)
	}

	numerator, denominator, err := k.GetRate(ctx, vault.DebtDenom, vault.Collateral.Denom)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "rate not found")
	}

	// Collateralization Ratio = (vault_collateral * numerator * 100) / (denominator * vault_denoms)
	collateralization_ratio := (vault.Collateral.Amount.Mul(numerator).Mul(sdk.NewIntFromUint64(100))).Quo(denominator.Mul(DebtAmount(denom, vault)))
	if collateralization_ratio.GT(sdk.NewIntFromUint64((collateral.LiquidationRatio))) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateralization ratio not less than liquidation ratio")
	}

	liquidatorReward, ok := sdk.NewIntFromString(k.getParams(ctx).LiquidatorReward)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "liquidator reward not integer string")
	}

	liquidatorCoinAmount := (vault.Collateral.Amount.Mul(liquidatorReward)).Quo(sdk.NewIntFromUint64(10000))
	liquidatorCoin := sdk.NewCoin(vault.Collateral.Denom, liquidatorCoinAmount)

	burnRate, ok := sdk.NewIntFromString(k.getParams(ctx).BurnRate)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "burn rate not integer string")
	}

	burnCoinAmount := ((vault.Collateral.Amount.Sub(liquidatorCoinAmount)).Mul(burnRate)).Quo(sdk.NewIntFromUint64(10000))
	burnCoin := sdk.NewCoin(vault.Collateral.Denom, burnCoinAmount)

	reserveCoin := vault.Collateral.Sub(burnCoin.Add(liquidatorCoin))
	reserve, found := k.GetReserve(ctx)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "resserve not found")
	}

	reserve.Reserves = reserve.Reserves.Add(reserveCoin)
	k.SetReserve(ctx, reserve)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidator, sdk.NewCoins(liquidatorCoin))
	if err != nil {
		return nil, err
	}

	return &types.MsgLiquidateResponse{}, nil
}
