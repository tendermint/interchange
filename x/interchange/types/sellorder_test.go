package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/interchange/types"
	"sort"
	"testing"
)

func TestNewSellOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	book := types.NewSellOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint32(0), book.OrderIDTrack)
	require.Equal(t, amountDenom, book.AmountDenom)
	require.Equal(t, priceDenom, book.PriceDenom)
	require.Empty(t, book.Orders)
}

type liquidateSellRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated types.Order
	Gain       uint32
	Match      bool
	Filled     bool
}

func simulateLiquidateFromSellOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected liquidateSellRes,
) {
	inputBook := OrderListToBuyOrderBook(inputList)
	expectedBook := OrderListToBuyOrderBook(expected.Book)
	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	outputBook, remaining, liquated, gain, match, filled := types.LiquidateFromSellOrder(inputBook, inputOrder)

	require.Equal(t, expectedBook, outputBook)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquated)
	require.Equal(t, expected.Gain, gain)
	require.Equal(t, expected.Match, match)
	require.Equal(t, expected.Filled, filled)
}

func TestLiquidateFromSellOrder(t *testing.T) {

}

type fillSellRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated types.Order
	Gain       uint32
	Filled     bool
}

func simulateFillSellOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected fillSellRes,
) {
	inputBook := OrderListToBuyOrderBook(inputList)
	expectedBook := OrderListToBuyOrderBook(expected.Book)
	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	outputBook, remaining, liquated, gain, filled := types.FillSellOrder(inputBook, inputOrder)

	require.Equal(t, expectedBook, outputBook)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquated)
	require.Equal(t, expected.Gain, gain)
	require.Equal(t, expected.Filled, filled)
}

func TestFillSellOrder(t *testing.T) {

}
