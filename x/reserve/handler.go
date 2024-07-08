package reserve

import (
	"fmt"

	"reserve/x/reserve/keeper"
	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateVault:
			res, err := msgServer.CreateVault(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDeposit:
			res, err := msgServer.Deposit(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdraw:
			res, err := msgServer.Withdraw(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgLiquidate:
			res, err := msgServer.Liquidate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBond:
			res, err := msgServer.Bond(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUnbond:
			res, err := msgServer.Unbond(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// NewSoftwareUpgradeProposalHandler creates a governance handler to manage new proposal types.
// It enables SoftwareUpgradeProposal to propose an Upgrade, and CancelSoftwareUpgradeProposal
// to abort a previously voted upgrade.
func NewReserveProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterCollateralProposal:
			return handleRegisterCollateralProposal(ctx, k, c)

		case *types.CreateDenomProposal:
			return handleCreateDenomProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized reserve proposal content type: %T", c)
		}
	}
}

func handleRegisterCollateralProposal(ctx sdk.Context, k keeper.Keeper, c *types.RegisterCollateralProposal) error {
	return k.RegisterCollateralProposal(ctx, c)
}

func handleCreateDenomProposal(ctx sdk.Context, k keeper.Keeper, c *types.CreateDenomProposal) error {
	return k.CreateDenomProposal(ctx, c)
}
