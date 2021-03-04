package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/interchange/types"
	"sort"
	"testing"
)

func TestNewSellOrderBook(t *testing.T) {
	amountDenom, priceDenom := types2.GenPair()
	book := types2.NewSellOrderBook(amountDenom, priceDenom)
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

func TestSellOrderBook_RemoveOrderFromID(t *testing.T) {
	inputList := []types.Order{
		{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
		{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
		{ID: 3, Creator: types2.MockAccount("3"), Amount: 2, Price: 10},
	}

	inputBook := types2.OrderListToSellOrderBook(inputList)
	expectedList := []types.Order{
		{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
		{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		{ID: 3, Creator: types2.MockAccount("3"), Amount: 2, Price: 10},
	}
	expectedBook := types2.OrderListToSellOrderBook(expectedList)
	outputBook, err := inputBook.RemoveOrderFromID(2)
	require.NoError(t, err)
	require.Equal(t, expectedBook, outputBook)

	inputBook = types2.OrderListToSellOrderBook(inputList)
	expectedList = []types.Order{
		{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
		{ID: 3, Creator: types2.MockAccount("3"), Amount: 2, Price: 10},
	}
	expectedBook = types2.OrderListToSellOrderBook(expectedList)
	outputBook, err = inputBook.RemoveOrderFromID(0)
	require.NoError(t, err)
	require.Equal(t, expectedBook, outputBook)

	inputBook = types2.OrderListToSellOrderBook(inputList)
	expectedList = []types.Order{
		{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
		{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
	}
	expectedBook = types2.OrderListToSellOrderBook(expectedList)
	outputBook, err = inputBook.RemoveOrderFromID(3)
	require.NoError(t, err)
	require.Equal(t, expectedBook, outputBook)

	inputBook = types2.OrderListToSellOrderBook(inputList)
	_, err = inputBook.RemoveOrderFromID(4)
	require.ErrorIs(t, err, types2.ErrOrderNotFound)
}

func simulateLiquidateFromSellOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected liquidateSellRes,
) {
	inputBook := types2.OrderListToBuyOrderBook(inputList)
	expectedBook := types2.OrderListToBuyOrderBook(expected.Book)
	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	outputBook, remaining, liquidated, gain, match, filled := types2.LiquidateFromSellOrder(inputBook, inputOrder)

	require.Equal(t, expectedBook, outputBook)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquidated)
	require.Equal(t, expected.Gain, gain)
	require.Equal(t, expected.Match, match)
	require.Equal(t, expected.Filled, filled)
}

func TestLiquidateFromSellOrder(t *testing.T) {
	// No match for empty book
	inputOrder := types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 100, Price: 30}
	_, _, _, _, match, _ := types2.LiquidateFromSellOrder(types.BuyOrderBook{
		Orders: []types.Order{},
	}, inputOrder)
	require.False(t, match)

	// Buy book
	inputBook := []types.Order{
		{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
	}

	// Test no match if highest bid too low (25 < 30)
	_, _, _, _, match, _ = types2.LiquidateFromSellOrder(types.BuyOrderBook{
		Orders: inputBook,
	}, inputOrder)
	require.False(t, match)

	// Entirely filled (30 < 50)
	inputOrder = types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 30, Price: 22}
	expected := liquidateSellRes{
		Book: []types.Order{
			{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
			{ID: 0, Creator: types2.MockAccount("0"), Amount: 20, Price: 25},
		},
		Remaining:  types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 0, Price: 22},
		Liquidated: types.Order{ID: 0, Creator: types2.MockAccount("0"), Amount: 30, Price: 25},
		Gain:       uint32(30 * 25),
		Match:      true,
		Filled:     true,
	}
	simulateLiquidateFromSellOrder(t, inputBook, inputOrder, expected)

	// Entirely filled and liquidated ( 50 = 50)
	inputOrder = types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 50, Price: 15}
	expected = liquidateSellRes{
		Book: []types.Order{
			{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining:  types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 0, Price: 15},
		Liquidated: types.Order{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
		Gain:       uint32(50 * 25),
		Match:      true,
		Filled:     true,
	}
	simulateLiquidateFromSellOrder(t, inputBook, inputOrder, expected)

	// Not filled and entirely liquidated (60 > 50)
	inputOrder = types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 60, Price: 10}
	expected = liquidateSellRes{
		Book: []types.Order{
			{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining:  types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 10, Price: 10},
		Liquidated: types.Order{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
		Gain:       uint32(50 * 25),
		Match:      true,
		Filled:     false,
	}
	simulateLiquidateFromSellOrder(t, inputBook, inputOrder, expected)
}

type fillSellRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated []types.Order
	Gain       uint32
	Filled     bool
}

func simulateFillSellOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected fillSellRes,
) {
	inputBook := types2.OrderListToBuyOrderBook(inputList)
	expectedBook := types2.OrderListToBuyOrderBook(expected.Book)
	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	outputBook, remaining, liquidated, gain, filled := types2.FillSellOrder(inputBook, inputOrder)

	require.Equal(t, expectedBook, outputBook)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquidated)
	require.Equal(t, expected.Gain, gain)
	require.Equal(t, expected.Filled, filled)
}

func TestFillSellOrder(t *testing.T) {
	var inputBook []types.Order

	// Empty book
	inputOrder := types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 30, Price: 30}
	expected := fillSellRes{
		Book:       []types.Order{},
		Remaining:  inputOrder,
		Liquidated: []types.Order(nil),
		Gain:       uint32(0),
		Filled:     false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// No match
	inputBook = []types.Order{
		{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
	}
	expected = fillSellRes{
		Book:       inputBook,
		Remaining:  inputOrder,
		Liquidated: []types.Order(nil),
		Gain:       uint32(0),
		Filled:     false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// First order liquidated, not filled
	inputOrder = types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 60, Price: 22}
	expected = fillSellRes{
		Book: []types.Order{
			{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining: types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 10, Price: 22},
		Liquidated: []types.Order{
			{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
		},
		Gain:   uint32(50 * 25),
		Filled: false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// Filled with two order
	inputOrder = types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 60, Price: 18}
	expected = fillSellRes{
		Book: []types.Order{
			{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 190, Price: 20},
		},
		Remaining:  types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 0, Price: 18},
		Liquidated: []types.Order{
			{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 10, Price: 20},
		},
		Gain:       uint32(50*25+10*20),
		Filled:     true,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)

	// Not filled, buy order book liquidated
	inputOrder = types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 300, Price: 10}
	expected = fillSellRes{
		Book: []types.Order{},
		Remaining:  types.Order{ID: 10, Creator: types2.MockAccount("1"), Amount: 20, Price: 10},
		Liquidated: []types.Order{
			{ID: 0, Creator: types2.MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: types2.MockAccount("1"), Amount: 200, Price: 20},
			{ID: 2, Creator: types2.MockAccount("2"), Amount: 30, Price: 15},
		},
		Gain:       uint32(50*25+200*20+30*15),
		Filled:     false,
	}
	simulateFillSellOrder(t, inputBook, inputOrder, expected)
}
