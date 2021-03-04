package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendCancelOrder{}

func NewMsgSendCancelOrder(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	amountDenom string,
	priceDenom string,
	orderID int32,
) *MsgSendCancelOrder {
	return &MsgSendCancelOrder{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		AmountDenom:      amountDenom,
		PriceDenom:       priceDenom,
		OrderID:          orderID,
	}
}

func (msg *MsgSendCancelOrder) Route() string {
	return RouterKey
}

func (msg *MsgSendCancelOrder) Type() string {
	return "SendCancelOrder"
}

func (msg *MsgSendCancelOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendCancelOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
