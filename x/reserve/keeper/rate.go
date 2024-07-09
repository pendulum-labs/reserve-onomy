package keeper

import (
	"sort"
	"strings"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	markettypes "github.com/pendulum-labs/market/x/market/types"
)

// GetRate takes base (numerator) and quote (denominator) as input
// GetRate gives exchange rate in integer array format [numerator, denominator]
// Rate is simply base balance divided by quote balance
func (k Keeper) GetRate(
	ctx sdk.Context,
	collateral string,
	pegPairs []string,
) (rate sdk.Dec, err error) {
	rateSum := sdk.NewDec(0)
	var i uint64

	for _, pegPair := range pegPairs {
		pegPairArray := strings.Split(pegPair, ",")
		peg := pegPairArray[0]
		if pegPairArray[1] == collateral {
			sort.Strings(pegPairArray)
			pair := strings.Join(pegPairArray, ",")

			// Check to see if Pool exists
			_, found := k.marketKeeper.GetPool(ctx, pair)
			if !found {
				return rate, sdkerrors.Wrapf(markettypes.ErrPoolNotFound, "%s", pair)
			}

			memberPeg, found := k.marketKeeper.GetMember(ctx, collateral, peg)
			if !found {
				return rate, sdkerrors.Wrapf(markettypes.ErrMemberNotFound, "%s", strings.Join([]string{collateral, peg}, ", "))
			}

			memberCollateral, found := k.marketKeeper.GetMember(ctx, peg, collateral)
			if !found {
				return rate, sdkerrors.Wrapf(markettypes.ErrMemberNotFound, "%s", strings.Join([]string{peg, collateral}, ", "))
			}

			if memberCollateral.Balance.LTE(sdk.ZeroInt()) {
				return rate, sdkerrors.Wrapf(types.ErrZero, "collateral %s balance is zero in pool %s", collateral, pair)
			}
			rateSum = (memberPeg.Balance.ToDec().Quo(memberCollateral.Balance.ToDec())).Add(rateSum)
			i++
		}
	}

	rate = rateSum.Quo(sdk.NewDecFromInt(sdk.NewIntFromUint64(i)))

	return rate, err
}
