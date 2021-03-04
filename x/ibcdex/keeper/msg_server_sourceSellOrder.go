package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) SendSourceSellOrder(goCtx context.Context, msg *types.MsgSendSourceSellOrder) (*types.MsgSendSourceSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.SourceSellOrderPacketData

	packet.Seller = msg.Seller
	packet.AmountDenom = msg.AmountDenom
	packet.Amount = msg.Amount
	packet.PriceDenom = msg.PriceDenom
	packet.Price = msg.Price

	// Transmit the packet
	err := k.TransmitSourceSellOrderPacket(
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
