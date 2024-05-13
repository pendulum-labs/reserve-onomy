package keeper_test

import (
	"context"
	"testing"

	keepertest "denom/testutil/keeper"
	"denom/x/reserve/keeper"
	"denom/x/reserve/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ReserveKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
