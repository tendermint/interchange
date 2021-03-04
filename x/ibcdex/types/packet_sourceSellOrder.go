package types

// ValidateBasic is used for validating the packet
func (p SourceSellOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p SourceSellOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcdexPacketData

	modulePacket.Packet = &IbcdexPacketData_SourceSellOrderPacket{&p}

	return modulePacket.Marshal()
}
