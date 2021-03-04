package ibcdex

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

// OnChanOpenInit implements the IBCModule interface
func (am AppModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) error {
	if order != channeltypes.UNORDERED {
		return sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.UNORDERED, order)
	}

	// Require portID is the portID module is bound to
	boundPort := am.keeper.GetPort(ctx)
	if boundPort != portID {
		return sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if version != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.Version)
	}

	// Claim channel capability passed back by IBC module
	if err := am.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return err
	}

	return nil
}

// OnChanOpenTry implements the IBCModule interface
func (am AppModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version,
	counterpartyVersion string,
) error {
	if order != channeltypes.UNORDERED {
		return sdkerrors.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.UNORDERED, order)
	}

	// Require portID is the portID module is bound to
	boundPort := am.keeper.GetPort(ctx)
	if boundPort != portID {
		return sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if version != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "got: %s, expected %s", version, types.Version)
	}

	if counterpartyVersion != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", counterpartyVersion, types.Version)
	}

	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
	// If module can already authenticate the capability then module already owns it so we don't need to claim
	// Otherwise, module does not have channel capability and we must claim it from IBC
	if !am.keeper.AuthenticateCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)) {
		// Only claim channel capability passed back by IBC module if we do not already own it
		if err := am.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
			return err
		}
	}

	return nil
}

// OnChanOpenAck implements the IBCModule interface
func (am AppModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyVersion string,
) error {
	if counterpartyVersion != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.Version)
	}
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (am AppModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (am AppModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (am AppModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (am AppModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (*sdk.Result, []byte, error) {
	var modulePacketData types.IbcdexPacketData
	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
		return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	var ack channeltypes.Acknowledgement

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	// this line is used by starport scaffolding # ibc/packet/module/recv
	case *types.IbcdexPacketData_CancelOrderPacket:
		packetAck, err := am.keeper.OnRecvCancelOrderPacket(ctx, modulePacket, *packet.CancelOrderPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCancelOrderPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	case *types.IbcdexPacketData_TargetBuyOrderPacket:
		packetAck, err := am.keeper.OnRecvTargetBuyOrderPacket(ctx, modulePacket, *packet.TargetBuyOrderPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeTargetBuyOrderPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	case *types.IbcdexPacketData_SourceBuyOrderPacket:
		packetAck, err := am.keeper.OnRecvSourceBuyOrderPacket(ctx, modulePacket, *packet.SourceBuyOrderPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSourceBuyOrderPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	case *types.IbcdexPacketData_TargetSellOrderPacket:
		packetAck, err := am.keeper.OnRecvTargetSellOrderPacket(ctx, modulePacket, *packet.TargetSellOrderPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeTargetSellOrderPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	case *types.IbcdexPacketData_SourceSellOrderPacket:
		packetAck, err := am.keeper.OnRecvSourceSellOrderPacket(ctx, modulePacket, *packet.SourceSellOrderPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSourceSellOrderPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	case *types.IbcdexPacketData_CreatePairPacket:
		packetAck, err := am.keeper.OnRecvCreatePairPacket(ctx, modulePacket, *packet.CreatePairPacket)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err.Error())
		} else {
			// Encode packet acknowledgment
			packetAckBytes, err := packetAck.Marshal()
			if err != nil {
				return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
			}
			ack = channeltypes.NewResultAcknowledgement(packetAckBytes)
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreatePairPacket,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
			),
		)
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	// Encode acknowledgement
	ackBytes, err := ack.Marshal()
	if err != nil {
		return nil, []byte{}, sdkerrors.Wrap(sdkerrors.ErrInvalidType, err.Error())
	}

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, ackBytes, nil
}

// OnAcknowledgementPacket implements the IBCModule interface
func (am AppModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
) (*sdk.Result, error) {
	var ack channeltypes.Acknowledgement
	if err := ack.Unmarshal(acknowledgement); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
	}
	var modulePacketData types.IbcdexPacketData
	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	var eventType string

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	// this line is used by starport scaffolding # ibc/packet/module/ack
	case *types.IbcdexPacketData_CancelOrderPacket:
		err := am.keeper.OnAcknowledgementCancelOrderPacket(ctx, modulePacket, *packet.CancelOrderPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeCancelOrderPacket
	case *types.IbcdexPacketData_TargetBuyOrderPacket:
		err := am.keeper.OnAcknowledgementTargetBuyOrderPacket(ctx, modulePacket, *packet.TargetBuyOrderPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeTargetBuyOrderPacket
	case *types.IbcdexPacketData_SourceBuyOrderPacket:
		err := am.keeper.OnAcknowledgementSourceBuyOrderPacket(ctx, modulePacket, *packet.SourceBuyOrderPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeSourceBuyOrderPacket
	case *types.IbcdexPacketData_TargetSellOrderPacket:
		err := am.keeper.OnAcknowledgementTargetSellOrderPacket(ctx, modulePacket, *packet.TargetSellOrderPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeTargetSellOrderPacket
	case *types.IbcdexPacketData_SourceSellOrderPacket:
		err := am.keeper.OnAcknowledgementSourceSellOrderPacket(ctx, modulePacket, *packet.SourceSellOrderPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeSourceSellOrderPacket
	case *types.IbcdexPacketData_CreatePairPacket:
		err := am.keeper.OnAcknowledgementCreatePairPacket(ctx, modulePacket, *packet.CreatePairPacket, ack)
		if err != nil {
			return nil, err
		}
		eventType = types.EventTypeCreatePairPacket
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			eventType,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
		),
	)

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
			),
		)
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// OnTimeoutPacket implements the IBCModule interface
func (am AppModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	var modulePacketData types.IbcdexPacketData
	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	// this line is used by starport scaffolding # ibc/packet/module/timeout
	case *types.IbcdexPacketData_CancelOrderPacket:
		err := am.keeper.OnTimeoutCancelOrderPacket(ctx, modulePacket, *packet.CancelOrderPacket)
		if err != nil {
			return nil, err
		}
	case *types.IbcdexPacketData_TargetBuyOrderPacket:
		err := am.keeper.OnTimeoutTargetBuyOrderPacket(ctx, modulePacket, *packet.TargetBuyOrderPacket)
		if err != nil {
			return nil, err
		}
	case *types.IbcdexPacketData_SourceBuyOrderPacket:
		err := am.keeper.OnTimeoutSourceBuyOrderPacket(ctx, modulePacket, *packet.SourceBuyOrderPacket)
		if err != nil {
			return nil, err
		}
	case *types.IbcdexPacketData_TargetSellOrderPacket:
		err := am.keeper.OnTimeoutTargetSellOrderPacket(ctx, modulePacket, *packet.TargetSellOrderPacket)
		if err != nil {
			return nil, err
		}
	case *types.IbcdexPacketData_SourceSellOrderPacket:
		err := am.keeper.OnTimeoutSourceSellOrderPacket(ctx, modulePacket, *packet.SourceSellOrderPacket)
		if err != nil {
			return nil, err
		}
	case *types.IbcdexPacketData_CreatePairPacket:
		err := am.keeper.OnTimeoutCreatePairPacket(ctx, modulePacket, *packet.CreatePairPacket)
		if err != nil {
			return nil, err
		}
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
