package types

import (
	"fmt"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		DenomTraceList:    []*DenomTrace{},
		BuyOrderBookList:  []*BuyOrderBook{},
		SellOrderBookList: []*SellOrderBook{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in denomTrace
	denomTraceIndexMap := make(map[string]bool)

	for _, elem := range gs.DenomTraceList {
		if _, ok := denomTraceIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for denomTrace")
		}
		denomTraceIndexMap[elem.Index] = true
	}
	// Check for duplicated index in buyOrderBook
	buyOrderBookIndexMap := make(map[string]bool)

	for _, elem := range gs.BuyOrderBookList {
		if _, ok := buyOrderBookIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for buyOrderBook")
		}
		buyOrderBookIndexMap[elem.Index] = true
	}
	// Check for duplicated index in sellOrderBook
	sellOrderBookIndexMap := make(map[string]bool)

	for _, elem := range gs.SellOrderBookList {
		if _, ok := sellOrderBookIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for sellOrderBook")
		}
		sellOrderBookIndexMap[elem.Index] = true
	}

	return nil
}
