package types

import ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"

// VoucherDenom returns the voucher of the denom from the port ID and channel ID
func VoucherDenom(port string, channel string, denom string) string {
	// since SendPacket did not prefix the denomination, we must prefix denomination here
	sourcePrefix := ibctransfertypes.GetDenomPrefix(port, channel)

	// NOTE: sourcePrefix contains the trailing "/"
	prefixedDenom := sourcePrefix + denom

	// construct the denomination trace from the full raw denomination
	denomTrace := ibctransfertypes.ParseDenomTrace(prefixedDenom)

	return denomTrace.IBCDenom()
}

// OriginalDenom returns back the original denom of the voucher
// False is returned if the port ID and channel ID provided are not the origins of the voucher
func OriginalDenom(port string, channel string, voucher string) (string, bool) {
	// Check if original chain
	if ibctransfertypes.ReceiverChainIsSource(port, channel, voucher) {
		// remove prefix added by sender chain
		voucherPrefix := ibctransfertypes.GetDenomPrefix(port, channel)
		unprefixedDenom := voucher[len(voucherPrefix):]

		// coin denomination used in sending from the escrow address
		denom := unprefixedDenom

		// The denomination used to send the coins is either the native denom or the hash of the path
		// if the denomination is not native.
		denomTrace := ibctransfertypes.ParseDenomTrace(unprefixedDenom)
		if denomTrace.Path != "" {
			denom = denomTrace.IBCDenom()
		}

		return denom, true
	}

	// Not the original chain
	return "", false
}