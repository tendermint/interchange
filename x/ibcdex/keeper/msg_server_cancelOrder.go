package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) SendCancelOrder(goCtx context.Context, msg *types.MsgSendCancelOrder) (*types.MsgSendCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.CancelOrderPacketData

	packet.AmountDenom = msg.AmountDenom
	packet.PriceDenom = msg.PriceDenom
	packet.OrderID = msg.OrderID

	// Transmit the packet
	err := k.TransmitCancelOrderPacket(
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

	return &types.MsgSendCancelOrderResponse{}, nil
}
