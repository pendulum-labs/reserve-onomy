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
	base string,
	quote string,
) (numerator sdk.Int, denominator sdk.Int, err error) {
	pairArray := []string{base, quote}

	sort.Strings(pairArray)

	pair := strings.Join(pairArray, ",")

	// Check to see if Pool exists
	_, found := k.marketKeeper.GetPool(ctx, pair)
	if !found {
		return numerator, denominator, sdkerrors.Wrapf(markettypes.ErrPoolNotFound, "%s", pair)
	}

	memberBase, found := k.marketKeeper.GetMember(ctx, quote, base)
	if !found {
		return numerator, denominator, sdkerrors.Wrapf(markettypes.ErrMemberNotFound, "%s", strings.Join([]string{quote, base}, ", "))
	}

	memberQuote, found := k.marketKeeper.GetMember(ctx, base, quote)
	if !found {
		return numerator, denominator, sdkerrors.Wrapf(markettypes.ErrMemberNotFound, "%s", strings.Join([]string{base, quote}, ", "))
	}

	return memberBase.Balance, memberQuote.Balance, err
}
