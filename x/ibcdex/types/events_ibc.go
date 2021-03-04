package types

// IBC events
const (
	EventTypeTimeout = "timeout"
	// this line is used by starport scaffolding # ibc/packet/event
	EventTypeCancelOrderPacket = "cancelOrder_packet"

	EventTypeTargetBuyOrderPacket = "targetBuyOrder_packet"

	EventTypeSourceBuyOrderPacket = "sourceBuyOrder_packet"

	EventTypeTargetSellOrderPacket = "targetSellOrder_packet"

	EventTypeSourceSellOrderPacket = "sourceSellOrder_packet"

	EventTypeCreatePairPacket = "createPair_packet"

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
