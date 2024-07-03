package keeper

import (
	"context"
	"strconv"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
		case vault.Status == "ready" && vault.Collateral.Denom == coin.Denom:
			vault.Collateral = vault.Collateral.Add(coin)
		case vault.Status == "active" && vault.Collateral.Denom == coin.Denom:
			vault.Collateral = vault.Collateral.Add(coin)
		case vault.Status == "active" && vault.DebtDenom == coin.Denom:
			denom, found := k.GetDenom(ctx, coin.Denom)
			if !found {
				sdkerrors.Wrapf(types.ErrDenomNotFound, "Denom with name %s not found", coin.Denom)
			}
			debt_owed := (vault.DebtShares.Mul(denom.DebtDenoms)).Quo(denom.DebtShares)
			deposit := coin.Amount
			// Cannot deposit more denoms than owed
			if coin.Amount.GT(sdk.Int(debt_owed)) {
				deposit = sdk.Int(debt_owed)
			}

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit")
		}
	} else {
		return nil, sdkerrors.Wrapf(types.ErrVaultNotFound, "Vault with uid %s not found", msg.Uid)
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}

	k.SetVault(ctx, vault)
	return &types.MsgDepositResponse{
		Uid:  msg.Uid,
		Coin: msg.Coin,
	}, nil
}
