package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) SendSourceSellOrder(goCtx context.Context, msg *types.MsgSendSourceSellOrder) (*types.MsgSendSourceSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Cannot send a order if the pair doesn't exist
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
	_, found := k.GetSellOrderBook(ctx, pairIndex)
	if !found {
		return &types.MsgSendSourceSellOrderResponse{}, errors.New("the pair doesn't exist")
	}

	// Lock the token to send
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return &types.MsgSendSourceSellOrderResponse{}, err
	}
	if err := k.LockTokens(
		ctx,
		msg.Port,
		msg.ChannelID,
		sender,
		sdk.NewCoin(msg.AmountDenom, sdk.NewInt(int64(msg.Amount))),
	); err != nil {
		return &types.MsgSendSourceSellOrderResponse{}, err
	}

	// Construct the packet
	var packet types.SourceSellOrderPacketData

	packet.Seller = msg.Seller
	packet.AmountDenom = msg.AmountDenom
	packet.Amount = msg.Amount
	packet.PriceDenom = msg.PriceDenom
	packet.Price = msg.Price

	// Transmit the packet
	err = k.TransmitSourceSellOrderPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendSourceSellOrderResponse{}, nil
}
