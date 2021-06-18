package types

import "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "ibcdex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ibcdex"

	// Version defines the current version the IBC module supports
	Version = "ibcdex-1"

	// PortID is the default port id that module binds to
	PortID = "ibcdex"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("ibcdex-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SellOrderBookKey = "SellOrderBook-value-"
)

const (
	BuyOrderBookKey = "BuyOrderBook-value-"
)

const (
	DenomTraceKey = "DenomTrace-value-"
)

func OrderBookIndex(
	portID string,
	channelID string,
	sourceDenom string,
	targetDenom string,
) string {
	return fmt.Sprintf("%s-%s-%s-%s",
		portID,
		channelID,
		sourceDenom,
		targetDenom,
	)
}
