package keeper_test

import (
	"strings"
	"testing"

	keepertest "reserve/testutil/keeper"
	"reserve/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	marketkeeper "github.com/pendulum-labs/market/x/market/keeper"
	markettypes "github.com/pendulum-labs/market/x/market/types"

	"reserve/x/reserve/keeper"
	"reserve/x/reserve/types"
)

func commonWithdraw(t *testing.T) (vault types.Vault) {
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
	return vault
}

func TestWithdraw(t *testing.T) {
	commonWithdraw(t)
}

func TestWithdraw_Insufficient_Funds(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA"}
	coinPair, _ := sample.SampleCoins("10CoinA", "10CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))

	// Create Vault
	var p = types.MsgCreate{Creator: addr, Collateral: testdata.coinAStr}
	response, err := keeper.NewMsgServerImpl(*testInput.ReserveKeeper).Create(sdk.WrapSDKContext(testInput.Context), &p)

	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient funds")
	require.NotContains(t, p.GetCreator(), response.String())

}
func TestWithdraw_With_New_Creator(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB"}
	coinPair, _ := sample.SampleCoins("10CoinA", "10CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))
	var p = markettypes.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: sample.AccAddress()}
	response, err := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient funds")
	require.NotContains(t, p.GetCreator(), response.String())

}

func TestWithdraw_With_Swap_Coins(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinB", coinBStr: "15CoinA"}
	coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))
	//validate SetUidCount function.
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	var p = markettypes.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.NoError(t, err)
	require.Contains(t, response.String(), p.GetCreator())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

}

func TestWithdraw_Invalid_Coin(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "20Coin", coinBStr: "20CoinB"},
		{coinAStr: "20CoinA", coinBStr: "20Coin"},
		//{coinAStr: "20CoinA", coinBStr: "20",},
	}
	for _, s := range scenarios {
		coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")
		require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
		requestAddress, _ := sdk.AccAddressFromBech32(addr)
		require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))
		var p = markettypes.MsgCreatePool{CoinA: s.coinAStr, CoinB: s.coinBStr, Creator: addr}
		response, err := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
		require.Error(t, err)
		require.NotContains(t, p.GetCreator(), response.String())

	}

}
