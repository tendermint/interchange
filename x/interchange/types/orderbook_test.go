package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/interchange/types"
	"math/rand"
	"sort"
	"testing"
)

func GenAmount() uint64 {
	return uint64(rand.Intn(int(types.MaxAmount)) + 1)
}

func GenPrice() uint64 {
	return uint64(rand.Intn(int(types.MaxPrice)) + 1)
}

func GenPair() (string, string) {
	return GenString(10), GenString(10)
}

func GenOrder() (types.Account, uint64, uint64) {
	return GenLocalAccount(), GenAmount(), GenPrice()
}

func OrderListToSellOrderBook(list []types.Order) types.SellOrderBook {
	listCopy := make([]types.Order, len(list))
	copy(listCopy, list)
	book := types.SellOrderBook{
		OrderIDTrack: 0,
		AmountDenom:  "foo",
		PriceDenom:   "bar",
		Orders: listCopy,
	}
	return book
}

func OrderListToBuyOrderBook(list []types.Order) types.SellOrderBook {
	book := types.SellOrderBook{
		OrderIDTrack: 0,
		AmountDenom:  "foo",
		PriceDenom:   "bar",
	}
	copy(book.Orders, list)
	return book
}

func TestAppendOrder(t *testing.T) {
	var ok bool
	sellBook := types.NewSellOrderBook(GenPair())

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, _, err := types.AppendOrder(sellBook, seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, _, err = types.AppendOrder(sellBook, seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, _, err = types.AppendOrder(sellBook, seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, _, err = types.AppendOrder(sellBook, seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)

	// Can append sell orders
	for i := 0; i < 20; i++ {
		// Append a new order
		creator, amount, price := GenOrder()
		newOrder := types.Order{
			ID:      sellBook.OrderIDTrack,
			Creator: creator,
			Amount:  amount,
			Price:   price,
		}
		newBook, orderID, err := types.AppendOrder(sellBook, creator, amount, price)
		sellBook, ok = newBook.(types.SellOrderBook)

		// Checks
		require.True(t, ok)
		require.NoError(t, err)
		require.Contains(t, sellBook.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)
	}
	require.Len(t, sellBook.Orders, 20)
	require.True(t, sort.IsSorted(sellBook))

	// Can append buy orders
	buyBook := types.NewBuyOrderBook(GenPair())
	for i := 0; i < 20; i++ {
		// Append a new order
		creator, amount, price := GenOrder()
		newOrder := types.Order{
			ID:     buyBook.OrderIDTrack,
			Creator: creator,
			Amount: amount,
			Price:  price,
		}
		newBook, orderID, err := types.AppendOrder(buyBook, creator, amount, price)
		buyBook, ok = newBook.(types.BuyOrderBook)

		// Checks
		require.True(t, ok)
		require.NoError(t, err)
		require.Contains(t, buyBook.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)
	}
	require.Len(t, buyBook.Orders, 20)
	require.True(t, sort.IsSorted(buyBook))
}

func simulateUpdateOrderBook(t *testing.T, sell bool, inputList []types.Order, inputOrder types.Order, expectedList []types.Order) {
	var inputBook types.OrderBook
	var expectedBook types.OrderBook
	if sell {
		inputBook = OrderListToSellOrderBook(inputList)
		expectedBook = OrderListToSellOrderBook(expectedList)
	} else {
		inputBook = OrderListToBuyOrderBook(inputList)
		expectedBook = OrderListToBuyOrderBook(expectedList)
	}

	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	simulatedBook := types.UpdateOrderBook(inputBook, inputOrder)

	require.Equal(t, expectedBook, simulatedBook)
}

func TestUpdateOrderBook(t *testing.T) {
	// Sell order book
	inputBook := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
	}

	// Sell 1
	inputOrder := types.Order{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20}
	expectedBook := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
	}
	simulateUpdateOrderBook(t, true, inputBook, inputOrder, expectedBook)

	// Sell 2
	inputOrder = types.Order{ID: 4, Creator: MockAccount("1"), Amount: 100, Price: 17}
	expectedBook = []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 4, Creator: MockAccount("1"), Amount: 100, Price: 17},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
	}
	simulateUpdateOrderBook(t,true, inputBook, inputOrder, expectedBook)

	// Sell 3
	inputOrder = types.Order{ID: 5, Creator: MockAccount("1"), Amount: 500, Price: 1}
	expectedBook = []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
		{ID: 5, Creator: MockAccount("1"), Amount: 500, Price: 1},
	}
	simulateUpdateOrderBook(t,true, inputBook, inputOrder, expectedBook)

	// Buy order book
	inputBook = []types.Order{
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
	}

	// Buy 1
	inputOrder = types.Order{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20}
	expectedBook = []types.Order{
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
	}
	simulateUpdateOrderBook(t, false, inputBook, inputOrder, expectedBook)

	// Buy 2
	inputOrder = types.Order{ID: 4, Creator: MockAccount("1"), Amount: 100, Price: 17}
	expectedBook = []types.Order{
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 4, Creator: MockAccount("1"), Amount: 100, Price: 17},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
	}
	simulateUpdateOrderBook(t,false, inputBook, inputOrder, expectedBook)

	// Buy 3
	inputOrder = types.Order{ID: 5, Creator: MockAccount("1"), Amount: 500, Price: 1}
	expectedBook = []types.Order{
		{ID: 5, Creator: MockAccount("1"), Amount: 500, Price: 1},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
	}
	simulateUpdateOrderBook(t,false, inputBook, inputOrder, expectedBook)
}

func simulateRestoreSellOrderBook(t *testing.T, sell bool, inputList []types.Order, liquidated []types.Order, expectedList []types.Order) {
	var inputBook types.OrderBook
	var expectedBook types.OrderBook
	if sell {
		inputBook = OrderListToSellOrderBook(inputList)
		expectedBook = OrderListToSellOrderBook(expectedList)
	} else {
		inputBook = OrderListToBuyOrderBook(inputList)
		expectedBook = OrderListToBuyOrderBook(expectedList)
	}

	require.True(t, sort.IsSorted(inputBook))
	require.True(t, sort.IsSorted(expectedBook))

	simulatedBook := types.RestoreOrderBook(inputBook, liquidated)

	require.Equal(t, expectedBook, simulatedBook)
}

func TestRestoreOrderBook(t *testing.T) {
	// Sell
	inputBook := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 4, Creator: MockAccount("3"), Amount: 45, Price: 10},
		{ID: 5, Creator: MockAccount("3"), Amount: 12, Price: 5},
	}
	liquidated := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 5, Creator: MockAccount("3"), Amount: 200, Price: 5},
		{ID: 6, Creator: MockAccount("4"), Amount: 40, Price: 30},
		{ID: 7, Creator: MockAccount("5"), Amount: 42, Price: 1},
	}
	expectedBook := []types.Order{
		{ID: 6, Creator: MockAccount("4"), Amount: 40, Price: 30},
		{ID: 0, Creator: MockAccount("0"), Amount: 150, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 4, Creator: MockAccount("3"), Amount: 45, Price: 10},
		{ID: 5, Creator: MockAccount("3"), Amount: 212, Price: 5},
		{ID: 7, Creator: MockAccount("5"), Amount: 42, Price: 1},
	}
	simulateRestoreSellOrderBook(t, true, inputBook, liquidated, expectedBook)

	// Buy
	inputBook = []types.Order{
		{ID: 5, Creator: MockAccount("3"), Amount: 12, Price: 5},
		{ID: 4, Creator: MockAccount("3"), Amount: 45, Price: 10},
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 50, Price: 25},
	}
	liquidated = []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 5, Creator: MockAccount("3"), Amount: 200, Price: 5},
		{ID: 6, Creator: MockAccount("4"), Amount: 40, Price: 30},
		{ID: 7, Creator: MockAccount("5"), Amount: 42, Price: 1},
	}
	expectedBook = []types.Order{
		{ID: 7, Creator: MockAccount("5"), Amount: 42, Price: 1},
		{ID: 5, Creator: MockAccount("3"), Amount: 212, Price: 5},
		{ID: 4, Creator: MockAccount("3"), Amount: 45, Price: 10},
		{ID: 3, Creator: MockAccount("3"), Amount: 2, Price: 10},
		{ID: 2, Creator: MockAccount("2"), Amount: 30, Price: 15},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 0, Creator: MockAccount("0"), Amount: 150, Price: 25},
		{ID: 6, Creator: MockAccount("4"), Amount: 40, Price: 30},
	}
	simulateRestoreSellOrderBook(t, false, inputBook, liquidated, expectedBook)
}