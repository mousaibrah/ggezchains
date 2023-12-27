package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/mousaibrah/ggezchains/testutil/keeper"
	"github.com/mousaibrah/ggezchains/x/trade/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TradeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
