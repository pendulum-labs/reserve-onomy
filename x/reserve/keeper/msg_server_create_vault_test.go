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

type testData struct {
	coinAStr      string
	coinBStr      string
	RateAstrArray []string
	RateBstrArray []string
	name          string
}

var addr string = sample.AccAddress()

func TestCreateVault(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	// TestData
	testdata := testData{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}, name: "testvault"}
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
	var p = types.MsgCreateVault{Creator: addr, Collateral: testdata.coinAStr, Name: testdata.name}
	response, err := keeper.NewMsgServerImpl(*testInput.ReserveKeeper).CreateVault(sdk.WrapSDKContext(testInput.Context), &p)

	// Validate CreateVault
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

func TestCreatePool_PoolAlreadyExist(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	count := 0
	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		// check reversed
		{coinAStr: "20CoinB", coinBStr: "20CoinA", RateAstrArray: []string{"20", "30"}, RateBstrArray: []string{"10", "20"}},
	}
	for _, s := range scenarios {
		coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")

		require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
		requestAddress, _ := sdk.AccAddressFromBech32(addr)
		require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))
		var p = markettypes.MsgCreatePool{CoinA: s.coinAStr, CoinB: s.coinBStr, Creator: addr}
		response, err := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
		if count == 0 {
			require.NoError(t, err)
			require.Contains(t, p.GetCreator(), response.String())

		} else {
			require.Error(t, err) //Pool Already exists
			require.ErrorContains(t, err, "pool already exists")
			require.NotContains(t, p.GetCreator(), response.String())
		}

		count++

	}

}

func TestCreatePool_Insufficient_Funds(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}}
	coinPair, _ := sample.SampleCoins("10CoinA", "10CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))
	var p = markettypes.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient funds")
	require.NotContains(t, p.GetCreator(), response.String())

}

func TestCreatePool_PoolAlready_Exists_ReSubmit(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))
	var p = markettypes.MsgCreatePool{CoinA: "15CoinA", CoinB: "15CoinB", Creator: addr}
	var p1 = markettypes.MsgCreatePool{CoinA: "30CoinA", CoinB: "30CoinB", Creator: addr}
	response, err := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	response1, err1 := marketkeeper.NewMsgServerImpl(testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p1)
	require.NoError(t, err)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "pool already exists")
	require.Contains(t, p.GetCreator(), response.String())
	require.NotContains(t, p.GetCreator(), response1.String())

}

func TestCreatePool_With_New_Creator(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}}
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

func TestCreatePool_With_Empty_Rates(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB", RateAstrArray: []string{"0", "0"}, RateBstrArray: []string{"0", "0"}}
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
	require.Contains(t, p.GetCreator(), response.String())
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

func TestCreatePool_With_Swap_Coins(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinB", coinBStr: "15CoinA", RateAstrArray: []string{"0", "0"}, RateBstrArray: []string{"0", "0"}}
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
	require.Contains(t, p.GetCreator(), response.String())
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

func TestCreatePool_Invalid_Coins(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "20Coin", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "20Coin", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		//{coinAStr: "20CoinA", coinBStr: "20", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
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
