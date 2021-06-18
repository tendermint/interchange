package types

// IBC events
const (
	EventTypeTimeout = "timeout"
	// this line is used by starport scaffolding # ibc/packet/event
	EventTypeBuyOrderPacket = "buyOrder_packet"

	EventTypeSellOrderPacket = "sellOrder_packet"

	EventTypeCreatePairPacket = "createPair_packet"

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
