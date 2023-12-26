package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/mousaibrah/ggezchain/testutil/keeper"
	"github.com/mousaibrah/ggezchain/testutil/nullify"
	"github.com/mousaibrah/ggezchain/x/trade/keeper"
	"github.com/mousaibrah/ggezchain/x/trade/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredTrade(keeper keeper.Keeper, ctx context.Context, n int) []types.StoredTrade {
	items := make([]types.StoredTrade, n)
	for i := range items {
		items[i].TradeIndex = strconv.Itoa(i)

		keeper.SetStoredTrade(ctx, items[i])
	}
	return items
}

func TestStoredTradeGet(t *testing.T) {
	keeper, ctx := keepertest.TradeKeeper(t)
	items := createNStoredTrade(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredTrade(ctx,
			item.TradeIndex,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredTradeRemove(t *testing.T) {
	keeper, ctx := keepertest.TradeKeeper(t)
	items := createNStoredTrade(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredTrade(ctx,
			item.TradeIndex,
		)
		_, found := keeper.GetStoredTrade(ctx,
			item.TradeIndex,
		)
		require.False(t, found)
	}
}

func TestStoredTradeGetAll(t *testing.T) {
	keeper, ctx := keepertest.TradeKeeper(t)
	items := createNStoredTrade(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredTrade(ctx)),
	)
}
