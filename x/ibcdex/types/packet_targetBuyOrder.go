package types

// ValidateBasic is used for validating the packet
func (p TargetBuyOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p TargetBuyOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcdexPacketData

	modulePacket.Packet = &IbcdexPacketData_TargetBuyOrderPacket{&p}

	return modulePacket.Marshal()
}
