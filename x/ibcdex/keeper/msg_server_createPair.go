package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) SendCreatePair(goCtx context.Context, msg *types.MsgSendCreatePair) (*types.MsgSendCreatePairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.CreatePairPacketData

	packet.SourceDenom = msg.SourceDenom
	packet.TargetDenom = msg.TargetDenom

	// Transmit the packet
	err := k.TransmitCreatePairPacket(
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

	return &types.MsgSendCreatePairResponse{}, nil
}
