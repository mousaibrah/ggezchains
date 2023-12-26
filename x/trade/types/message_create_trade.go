package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTrade{}

func NewMsgCreateTrade(creator string, tradeType string, coin string, price string, quantity string, receiverAddress string, tradeData string) *MsgCreateTrade {
	return &MsgCreateTrade{
		Creator:         creator,
		TradeType:       tradeType,
		Coin:            coin,
		Price:           price,
		Quantity:        quantity,
		ReceiverAddress: receiverAddress,
		TradeData:       tradeData,
	}
}

func (msg *MsgCreateTrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
