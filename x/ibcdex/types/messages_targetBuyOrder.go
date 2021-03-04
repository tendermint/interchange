package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendTargetBuyOrder{}

func NewMsgSendTargetBuyOrder(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	buyer string,
	amountDenom string,
	amount int32,
	priceDenom string,
	price int32,
) *MsgSendTargetBuyOrder {
	return &MsgSendTargetBuyOrder{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Buyer:            buyer,
		AmountDenom:      amountDenom,
		Amount:           amount,
		PriceDenom:       priceDenom,
		Price:            price,
	}
}

func (msg *MsgSendTargetBuyOrder) Route() string {
	return RouterKey
}

func (msg *MsgSendTargetBuyOrder) Type() string {
	return "SendTargetBuyOrder"
}

func (msg *MsgSendTargetBuyOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendTargetBuyOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendTargetBuyOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
