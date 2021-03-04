package types

// ValidateBasic is used for validating the packet
func (p CancelOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p CancelOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcdexPacketData

	modulePacket.Packet = &IbcdexPacketData_CancelOrderPacket{&p}

	return modulePacket.Marshal()
}
