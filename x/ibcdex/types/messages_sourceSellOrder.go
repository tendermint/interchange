package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendSourceSellOrder{}

func NewMsgSendSourceSellOrder(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	seller string,
	amountDenom string,
	amount int32,
	priceDenom string,
	price int32,
) *MsgSendSourceSellOrder {
	return &MsgSendSourceSellOrder{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Seller:           seller,
		AmountDenom:      amountDenom,
		Amount:           amount,
		PriceDenom:       priceDenom,
		Price:            price,
	}
}

func (msg *MsgSendSourceSellOrder) Route() string {
	return RouterKey
}

func (msg *MsgSendSourceSellOrder) Type() string {
	return "SendSourceSellOrder"
}

func (msg *MsgSendSourceSellOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendSourceSellOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendSourceSellOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
