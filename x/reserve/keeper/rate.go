package keeper

import (
	"sort"
	"strings"

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
	pegs []string,
) (rate sdk.Dec, err error) {
	var rateSum sdk.Dec

	for _, peg := range pegs {

		pairArray := []string{collateral, peg}

		sort.Strings(pairArray)

		pair := strings.Join(pairArray, ",")

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

		rateSum = (memberPeg.Balance.ToDec().Quo(memberCollateral.Balance.ToDec())).Add(rateSum)
	}

	rate = rateSum.Quo(sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(len(pegs)))))

	return rate, err
}
