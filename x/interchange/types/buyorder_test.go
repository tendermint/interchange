package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/interchange/types"
	"sort"
	"testing"
)

func TestNewBuyOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	book := types.NewBuyOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint32(1), book.OrderIDTrack)
	require.Equal(t, amountDenom, book.AmountDenom)
	require.Equal(t, priceDenom, book.PriceDenom)
	require.Empty(t, book.Orders)
}

type liquidateBuyRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated types.Order
	Purchase       uint32
	Match      bool
	Filled     bool
}

func TestBuyOrderBook_RemoveOrderFromID(t *testing.T) {
	inputList := []types.Order{
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}

	inputBook := OrderListToBuyOrderBook(inputList)
	expectedList := []types.Order{
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}
	expectedBook := OrderListToBuyOrderBook(expectedList)
	outputBook, err := inputBook.RemoveOrderFromID(2)
	require.NoError(t, err)
	require.Equal(t, expectedBook, outputBook)

	inputBook = OrderListToBuyOrderBook(inputList)
	expectedList = []types.Order{
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
	}
	expectedBook = OrderListToBuyOrderBook(expectedList)
	outputBook, err = inputBook.RemoveOrderFromID(0)
	require.NoError(t, err)
	require.Equal(t, expectedBook, outputBook)

	inputBook = OrderListToBuyOrderBook(inputList)
	expectedList = []types.Order{
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}
	expectedBook = OrderListToBuyOrderBook(expectedList)
	outputBook, err = inputBook.RemoveOrderFromID(3)
	require.NoError(t, err)
	require.Equal(t, expectedBook, outputBook)

	inputBook = OrderListToBuyOrderBook(inputList)
	_, err = inputBook.RemoveOrderFromID(4)
	require.ErrorIs(t, err, types.ErrOrderNotFound)
}

func simulateLiquidateFromBuyOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected liquidateBuyRes,
) {
	inputBook := OrderListToSellOrderBook(inputList)
	expectedBook := OrderListToSellOrderBook(expected.Book)
	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	outputBook, remaining, liquidated, purchase, match, filled := types.LiquidateFromBuyOrder(inputBook, inputOrder)

	require.Equal(t, expectedBook, outputBook)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquidated)
	require.Equal(t, expected.Purchase, purchase)
	require.Equal(t, expected.Match, match)
	require.Equal(t, expected.Filled, filled)
}

func TestLiquidateFromBuyOrder(t *testing.T) {
	// No match for empty book
	inputOrder := types.Order{ID: 10, Creator: MockAccount("1"), Amount: 100, Price: 10}
	_, _, _, _, match, _ := types.LiquidateFromBuyOrder(types.SellOrderBook{
		Orders: []types.Order{},
	}, inputOrder)
	require.False(t, match)

	// Sell book
	inputBook := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
	}

	// Test no match if lowest ask too high (25 < 30)
	_, _, _, _, match, _ = types.LiquidateFromBuyOrder(types.SellOrderBook{
		Orders: inputBook,
	}, inputOrder)
	require.False(t, match)

	// Entirely filled (30 > 15)
	inputOrder = types.Order{ID: 10, Creator: MockAccount("1"), Amount: 20, Price: 30}
	expected := liquidateBuyRes{
		Book: []types.Order{
			{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
			{ID: 2, Creator: MockAccount("2"), Amount: 10, Price: 15},
		},
		Remaining:  types.Order{ID: 10, Creator: MockAccount("1"), Amount: 0, Price: 30},
		Liquidated: types.Order{ID: 2, Creator: MockAccount("2"), Amount: 20, Price: 15},
		Purchase:   uint32(20),
		Match:      true,
		Filled:     true,
	}
	simulateLiquidateFromBuyOrder(t, inputBook, inputOrder, expected)

	// Entirely filled (30 = 30)
	inputOrder = types.Order{ID: 10, Creator: MockAccount("1"), Amount: 30, Price: 30}
	expected = liquidateBuyRes{
		Book: []types.Order{
			{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining:  types.Order{ID: 10, Creator: MockAccount("1"), Amount: 0, Price: 30},
		Liquidated: types.Order{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		Purchase:   uint32(30),
		Match:      true,
		Filled:     true,
	}
	simulateLiquidateFromBuyOrder(t, inputBook, inputOrder, expected)

	// Not filled and entirely liquidated (60 > 30)
	inputOrder = types.Order{ID: 10, Creator: MockAccount("1"), Amount: 60, Price: 30}
	expected = liquidateBuyRes{
		Book: []types.Order{
			{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining:  types.Order{ID: 10, Creator: MockAccount("1"), Amount: 30, Price: 30},
		Liquidated: types.Order{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		Purchase:   uint32(30),
		Match:      true,
		Filled:     false,
	}
	simulateLiquidateFromBuyOrder(t, inputBook, inputOrder, expected)
}

type fillBuyRes struct {
	Book       []types.Order
	Remaining  types.Order
	Liquidated []types.Order
	Purchase       uint32
	Filled     bool
}

func simulateFillBuyOrder(
	t *testing.T,
	inputList []types.Order,
	inputOrder types.Order,
	expected fillBuyRes,
) {
	inputBook := OrderListToSellOrderBook(inputList)
	expectedBook := OrderListToSellOrderBook(expected.Book)
	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	outputBook, remaining, liquidated, purchase, filled := types.FillBuyOrder(inputBook, inputOrder)

	require.Equal(t, expectedBook, outputBook)
	require.Equal(t, expected.Remaining, remaining)
	require.Equal(t, expected.Liquidated, liquidated)
	require.Equal(t, expected.Purchase, purchase)
	require.Equal(t, expected.Filled, filled)
}

func TestFillBuyOrder(t *testing.T) {
	var inputBook []types.Order

	// Empty book
	inputOrder := types.Order{ID: 10, Creator: MockAccount("1"), Amount: 30, Price: 10}
	expected := fillBuyRes{
		Book:       []types.Order{},
		Remaining:  inputOrder,
		Liquidated: []types.Order(nil),
		Purchase:       uint32(0),
		Filled:     false,
	}
	simulateFillBuyOrder(t, inputBook, inputOrder, expected)

	// No match
	inputBook = []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
	}
	expected = fillBuyRes{
		Book:       inputBook,
		Remaining:  inputOrder,
		Liquidated: []types.Order(nil),
		Purchase:       uint32(0),
		Filled:     false,
	}
	simulateFillBuyOrder(t, inputBook, inputOrder, expected)

	// First order liquidated, not filled
	inputOrder = types.Order{ID: 10, Creator: MockAccount("1"), Amount: 60, Price: 18}
	expected = fillBuyRes{
		Book: []types.Order{
			{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		},
		Remaining: types.Order{ID: 10, Creator: MockAccount("1"), Amount: 30, Price: 18},
		Liquidated: []types.Order{
			{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		},
		Purchase:   uint32(30),
		Filled: false,
	}
	simulateFillBuyOrder(t, inputBook, inputOrder, expected)

	// Filled with two order
	inputOrder = types.Order{ID: 10, Creator: MockAccount("1"), Amount: 60, Price: 22}
	expected = fillBuyRes{
		Book: []types.Order{
			{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
			{ID: 1, Creator: MockAccount("1"), Amount: 170, Price: 20},
		},
		Remaining: types.Order{ID: 10, Creator: MockAccount("1"), Amount: 0, Price: 22},
		Liquidated: []types.Order{
			{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: MockAccount("1"), Amount: 30, Price: 20},
		},
		Purchase:   uint32(30+30),
		Filled: true,
	}
	simulateFillBuyOrder(t, inputBook, inputOrder, expected)

	// Not filled, sell order book liquidated
	inputOrder = types.Order{ID: 10, Creator: MockAccount("1"), Amount: 300, Price: 30}
	expected = fillBuyRes{
		Book: []types.Order{},
		Remaining: types.Order{ID: 10, Creator: MockAccount("1"), Amount: 20, Price: 30},
		Liquidated: []types.Order{
			{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
			{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
			{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		},
		Purchase:   uint32(30+200+50),
		Filled: false,
	}
	simulateFillBuyOrder(t, inputBook, inputOrder, expected)
}