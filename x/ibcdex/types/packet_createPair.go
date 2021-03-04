package types

// ValidateBasic is used for validating the packet
func (p CreatePairPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p CreatePairPacketData) GetBytes() ([]byte, error) {
	var modulePacket IbcdexPacketData

	modulePacket.Packet = &IbcdexPacketData_CreatePairPacket{&p}

	return modulePacket.Marshal()
}
