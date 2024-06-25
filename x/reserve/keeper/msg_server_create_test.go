package keeper_test

import (
	"testing"

	keepertest "reserve/testutil/keeper"
	"reserve/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	markettypes "github.com/pendulum-labs/market/x/market/types"

	"reserve/x/reserve/keeper"
	"reserve/x/reserve/types"
)

type testData struct {
	coinAStr      string
	coinBStr      string
	RateAstrArray []string
	RateBstrArray []string
	name          string
}

func TestCreate(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	// TestData
	testdata := testData{coinAStr: "20CoinA", coinBStr: "20CoinB", name: "testvault"}
	coinPair, _ := sample.SampleCoins(testdata.coinAStr, testdata.coinBStr)

	// MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))

	// SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))

	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	// Create Vault
	var p = types.MsgCreate{Creator: addr, Collateral: testdata.coinAStr}
	response, err := keeper.NewMsgServerImpl(*testInput.ReserveKeeper).Create(sdk.WrapSDKContext(testInput.Context), &p)

	// Validate Create
	require.NoError(t, err)
	require.Equal(t, beforecount, response.Uid)

	// Validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate GetVault
	rst, found := testInput.ReserveKeeper.GetVault(testInput.Context, beforecount)
	require.True(t, found)
	require.Equal(t, rst.Uid, beforecount)
	require.Equal(t, rst.Name, testdata.name)
	require.Equal(t, rst.Collateral.String(), testdata.coinAStr)
}
