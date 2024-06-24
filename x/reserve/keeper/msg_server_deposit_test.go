package keeper_test

import (
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

var addr string = sample.AccAddress()

func commonDeposit(t *testing.T) (vault types.Vault) {
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
	var p = types.MsgCreateVault{Creator: addr, Collateral: testdata.coinAStr}
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
	return vault
}

func TestDeposit(t *testing.T) {
	commonDeposit(t)
}

func TestDeposit_Insufficient_Funds(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA"}
	coinPair, _ := sample.SampleCoins("10CoinA", "10CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, markettypes.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, markettypes.ModuleName, requestAddress, coinPair))

	// Create Vault
	var p = types.MsgCreateVault{Creator: addr, Collateral: testdata.coinAStr}
	response, err := keeper.NewMsgServerImpl(*testInput.ReserveKeeper).CreateVault(sdk.WrapSDKContext(testInput.Context), &p)

	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient funds")
	require.NotContains(t, p.GetCreator(), response.String())

}
func TestDeposit_With_New_Creator(t *testing.T) {
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
