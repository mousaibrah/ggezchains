package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mousaibrah/ggezchain/x/trade/types"
)

func (k msgServer) ProcessTrade(goCtx context.Context, msg *types.MsgProcessTrade) (*types.MsgProcessTradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgProcessTradeResponse{}, nil
}
