package trade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mousaibrah/ggezchain/x/trade/keeper"
	"github.com/mousaibrah/ggezchain/x/trade/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.TradeIndex != nil {
		k.SetTradeIndex(ctx, *genState.TradeIndex)
	}
	// Set all the storedTrade
	for _, elem := range genState.StoredTradeList {
		k.SetStoredTrade(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all tradeIndex
	tradeIndex, found := k.GetTradeIndex(ctx)
	if found {
		genesis.TradeIndex = &tradeIndex
	}
	genesis.StoredTradeList = k.GetAllStoredTrade(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}