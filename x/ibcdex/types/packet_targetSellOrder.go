package types

// ValidateBasic is used for validating the packet
func (p TargetSellOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p TargetSellOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcdexPacketData

	modulePacket.Packet = &IbcdexPacketData_TargetSellOrderPacket{&p}

	return modulePacket.Marshal()
}
