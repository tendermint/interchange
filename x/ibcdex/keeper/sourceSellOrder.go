package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

// TransmitSourceSellOrderPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitSourceSellOrderPacket(
	ctx sdk.Context,
	packetData types.SourceSellOrderPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvSourceSellOrderPacket processes packet reception
func (k Keeper) OnRecvSourceSellOrderPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SourceSellOrderPacketData) (packetAck types.SourceSellOrderPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// Check if the buy order book exists
	pairIndex := types.OrderBookIndex(packet.SourcePort, packet.SourceChannel, data.AmountDenom, data.PriceDenom)
	book, found := k.GetBuyOrderBook(ctx, pairIndex)
	if !found {
		return packetAck, errors.New("the pair doesn't exist")
	}

	// Fill sell order
	book, remaining, liquidated, gain, _ := types.FillSellOrder(book, types.Order{
		Amount: data.Amount,
		Price:  data.Price,
	})

	// Return remaining amount and gains
	packetAck.RemainingAmount = remaining.Amount
	packetAck.Gain = gain

	// Dispatch liquidated buy order
	for _, liquidation := range liquidated {
		liquidation := liquidation

		if liquidation.Creator.Remote {
			// Send token back for remote account
			packetAck.Sales = append(packetAck.Sales, &liquidation)
		} else {
			// Mint tokens for local account
			voucherDenom := types.VoucherDenom(packet.DestinationPort, packet.DestinationChannel, data.AmountDenom)

			addr, err := sdk.AccAddressFromBech32(liquidation.Creator.Address)
			if err != nil {
				return packetAck, err
			}

			if err := k.MintTokens(ctx, addr, sdk.NewCoin(voucherDenom, sdk.NewInt(int64(liquidation.Amount)))); err != nil {
				return packetAck, err
			}
		}
	}

	// Save the new order book
	k.SetBuyOrderBook(ctx, book)

	return packetAck, nil
}

// OnAcknowledgementSourceSellOrderPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementSourceSellOrderPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SourceSellOrderPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		// In case of error we mint back the native token
		receiver, err := sdk.AccAddressFromBech32(data.Seller)
		if err != nil {
			return err
		}

		if err := k.UnlockTokens(
			ctx,
			packet.SourcePort,
			packet.SourceChannel,
			receiver,
			sdk.NewCoin(data.AmountDenom, sdk.NewInt(int64(data.Amount))),
		); err != nil {
			return err
		}

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.SourceSellOrderPacketAck
		err := packetAck.Unmarshal(dispatchedAck.Result)
		if err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// Get the sell order book
		pairIndex := types.OrderBookIndex(packet.SourcePort, packet.SourceChannel, data.AmountDenom, data.PriceDenom)
		book, found := k.GetSellOrderBook(ctx, pairIndex)
		if !found {
			panic("sell order book must exist")
		}

		// Append the remaining amount of the order
		newBook, _, err := types.AppendOrder(
			book,
			types.Account{
				Address: data.Seller,
				Remote:  false,
			},
			packetAck.RemainingAmount,
			data.Price,
		)
		if err != nil {
			return err
		}
		book = newBook.(types.SellOrderBook)

		// Distribute the sales
		for _, sale := range packetAck.Sales {
			receiver, err := sdk.AccAddressFromBech32(sale.Creator.Address)
			if err != nil {
				return err
			}

			if err := k.UnlockTokens(
				ctx,
				packet.SourcePort,
				packet.SourceChannel,
				receiver,
				sdk.NewCoin(data.AmountDenom, sdk.NewInt(int64(sale.Amount))),
			); err != nil {
				return err
			}
		}

		// Mint the gains
		voucherDenom := types.VoucherDenom(packet.SourcePort, packet.SourceChannel, data.PriceDenom)
		receiver, err := sdk.AccAddressFromBech32(data.Seller)
		if err != nil {
			return err
		}
		if err := k.MintTokens(ctx, receiver, sdk.NewCoin(voucherDenom, sdk.NewInt(int64(packetAck.Gain)))); err != nil {
			return err
		}

		// Save the new order book
		k.SetSellOrderBook(ctx, book)

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutSourceSellOrderPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutSourceSellOrderPacket(ctx sdk.Context, packet channeltypes.Packet, data types.SourceSellOrderPacketData) error {
	// In case of error we mint back the native token
	receiver, err := sdk.AccAddressFromBech32(data.Seller)
	if err != nil {
		return err
	}

	if err := k.UnlockTokens(
		ctx,
		packet.SourcePort,
		packet.SourceChannel,
		receiver,
		sdk.NewCoin(data.AmountDenom, sdk.NewInt(int64(data.Amount))),
	); err != nil {
		return err
	}

	return nil
}
