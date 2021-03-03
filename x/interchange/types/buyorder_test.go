package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/interchange/types"
	"testing"
)

func TestNewBuyOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	book := types.NewBuyOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint64(1), book.OrderIDTrack)
	require.Equal(t, amountDenom, book.AmountDenom)
	require.Equal(t, priceDenom, book.PriceDenom)
	require.Empty(t, book.Orders)
}
