package keeper_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	keepertest "reserve/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	marketkeeper "github.com/pendulum-labs/market/x/market/keeper"
	markettypes "github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPoolWithMembers(keeper marketkeeper.Keeper, ctx sdk.Context, n int) []markettypes.Pool {
	poolItems := make([]markettypes.Pool, n)
	memberItems := make([]markettypes.Member, n)

	for i := range poolItems {
		poolItems[i].Denom1 = strconv.Itoa(i)
		poolItems[i].Denom2 = strconv.Itoa(i)
		poolItems[i].Leaders = []*markettypes.Leader{
			{
				Address: strconv.Itoa(i),
				Drops:   sdk.NewInt(int64(i)),
			},
		}
		poolItems[i].Drops = sdk.NewIntFromUint64(uint64(0))

		pairArray := []string{poolItems[i].Denom1, poolItems[i].Denom2}

		sort.Strings(pairArray)

		pair := strings.Join(pairArray, ",")

		poolItems[i].Pair = pair
		poolItems[i].Denom1 = pairArray[0]
		poolItems[i].Denom2 = pairArray[1]

		keeper.SetPool(ctx, poolItems[i])

		memberItems[i].Pair = strconv.Itoa(i)
		memberItems[i].DenomA = poolItems[i].Denom1
		memberItems[i].DenomB = poolItems[i].Denom2
		memberItems[i].Balance = sdk.NewIntFromUint64(uint64(0))
		memberItems[i].Previous = sdk.NewIntFromUint64(uint64(0))

		keeper.SetMember(ctx, memberItems[i])

		memberItems[i].DenomA = poolItems[i].Denom2
		memberItems[i].DenomB = poolItems[i].Denom1

		keeper.SetMember(ctx, memberItems[i])

	}
	return poolItems
}

func TestGetRate(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)

	items := createNPoolWithMembers(keeper.MarketKeeper, keeper.Context, 10)

	for _, item := range items {

		_, _, error := keeper.ReserveKeeper.GetRate(
			keeper.Context,
			item.Denom1,
			item.Denom2,
		)

		require.Nil(t, error)

		_, _, error = keeper.ReserveKeeper.GetRate(
			keeper.Context,
			item.Denom2,
			item.Denom1,
		)

		require.Nil(t, error)

	}
}
