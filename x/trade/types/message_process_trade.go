package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProcessTrade{}

func NewMsgProcessTrade(creator string, processType string, tradeIndex uint64) *MsgProcessTrade {
	return &MsgProcessTrade{
		Creator:     creator,
		ProcessType: processType,
		TradeIndex:  tradeIndex,
	}
}

func (msg *MsgProcessTrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
