package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllVaultsByOwner(goCtx context.Context, req *types.QueryGetAllVaultsByOwnerRequest) (*types.QueryGetAllVaultsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vaults := k.GetVaultsByOwner(ctx, req.Address)

	return &types.QueryGetAllVaultsByOwnerResponse{
		Vaults: vaults,
	}, nil
}
