package types_test

import (
	"testing"

	"github.com/mousaibrah/ggezchain/x/trade/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				TradeIndex: types.TradeIndex{
					NextId: 80,
				},
				StoredTradeList: []types.StoredTrade{
					{
						TradeIndex: 0,
					},
					{
						TradeIndex: 1,
					},
				},
				StoredTempTradeList: []types.StoredTempTrade{
					{
						TradeIndex: "0",
					},
					{
						TradeIndex: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedTrade",
			genState: &types.GenesisState{
				StoredTradeList: []types.StoredTrade{
					{
						TradeIndex: 0,
					},
					{
						TradeIndex: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedTempTrade",
			genState: &types.GenesisState{
				StoredTempTradeList: []types.StoredTempTrade{
					{
						TradeIndex: "0",
					},
					{
						TradeIndex: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
