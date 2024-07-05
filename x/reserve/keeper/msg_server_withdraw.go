package keeper

import (
	"context"
	"strconv"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	uid, err := strconv.ParseUint(msg.Uid, 10, 64)
	if err != nil {
		return nil, err
	}

	coin, err := sdk.ParseCoinNormalized(msg.Coin)
	if err != nil {
		return nil, err
	}

	vault, found := k.GetVault(ctx, uid)
	if found {
		switch {
		case vault.Status == "ready" && vault.Collateral.IsGTE(coin):
			vault.Collateral = vault.Collateral.Sub(coin)
			if vault.Collateral.Amount == sdk.ZeroInt() {
				vault.Status = "inactive"
			}
		case vault.Status == "active" && vault.Collateral.Denom == coin.Denom:
			// Need Minting Ratio from Collateral Record
			collateral, found := k.GetCollateral(ctx, vault.Collateral.Denom)
			if !found {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral not found")
			}

			denom, found := k.GetDenom(ctx, vault.DebtDenom)
			if !found {
				sdkerrors.Wrapf(types.ErrDenomNotFound, "Denom with name %s not found", vault.DebtDenom)
			}

			// Remainder is the amount of Collateral left in vault after withdrawal
			remainder := vault.Collateral.Sub(coin)

			numerator, denominator, err := k.GetRate(ctx, vault.DebtDenom, vault.Collateral.Denom)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "rate not found")
			}

			// Collateralization Ratio = (remainder_vault_collateral * numerator * 100) / (denominator * vault_denoms)
			remainder_collateralization_ratio := (remainder.Amount.Mul(numerator).Mul(sdk.NewIntFromUint64(100))).Quo(denominator.Mul(DebtAmount(denom, vault)))
			if remainder_collateralization_ratio.LT(sdk.NewIntFromUint64(collateral.MintingRatio)) {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "remainder collateralization ratio less than minting ratio")
			}

			vault.Collateral = vault.Collateral.Sub(coin)
		case vault.Status == "active" && vault.DebtDenom == coin.Denom:
			// Need Minting Ratio from Collateral Record
			collateral, found := k.GetCollateral(ctx, vault.Collateral.Denom)
			if !found {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral not found")
			}

			denom, found := k.GetDenom(ctx, vault.DebtDenom)
			if !found {
				sdkerrors.Wrapf(types.ErrDenomNotFound, "Denom with name %s not found", vault.DebtDenom)
			}

			// Lien is amount of Denoms owed to vault after minting
			lien := DebtAmount(denom, vault).Add(coin.Amount)

			numerator, denominator, err := k.GetRate(ctx, vault.DebtDenom, vault.Collateral.Denom)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "rate not found")
			}

			// Collateralization Ratio = (vault_collateral * numerator * 100) / (denominator * lien)
			lien_collateralization_ratio := (vault.Collateral.Amount.Mul(numerator).Mul(sdk.NewIntFromUint64(100))).Quo(denominator.Mul(lien.Amount))
			if lien_collateralization_ratio.LT(sdk.Int(collateral.MintingRatio)) {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "lien collateralization ratio less than minting ratio")
			}

			vault.Denom = vault.Denom.Add(coin)
			k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid withdraw")
		}
	} else {
		return nil, sdkerrors.Wrapf(types.ErrVaultNotFound, "Vault with uid %s not found", msg.Uid)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
