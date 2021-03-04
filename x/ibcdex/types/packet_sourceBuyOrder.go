package types

// ValidateBasic is used for validating the packet
func (p SourceBuyOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p SourceBuyOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcdexPacketData

	modulePacket.Packet = &IbcdexPacketData_SourceBuyOrderPacket{&p}

	return modulePacket.Marshal()
}
