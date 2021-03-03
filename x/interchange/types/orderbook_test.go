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
	book := types.NewSellOrderBook(GenPair())

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, _, err := types.AppendOrder(book, seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, _, err = types.AppendOrder(book, seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, _, err = types.AppendOrder(book, seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, _, err = types.AppendOrder(book, seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)

	// Can append sell orders
	for i := 0; i < 20; i++ {
		// Append a new order
		creator, amount, price := GenOrder()
		newOrder := types.Order{
			ID:      book.OrderIDTrack,
			Creator: creator,
			Amount:  amount,
			Price:   price,
		}
		newBook, orderID, err := types.AppendOrder(book, creator, amount, price)
		book, ok = newBook.(types.SellOrderBook)

		// Checks
		require.True(t, ok)
		require.NoError(t, err)
		require.Contains(t, book.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)
	}
	require.Len(t, book.Orders, 20)
	require.True(t, sort.IsSorted(book))

	// TODO: Can append buy orders
	//buyBook := types.NewBuyOrderBook(GenPair())
	//for i := 0; i < 20; i++ {
	//	// Append a new order
	//	creator, amount, price := GenOrder()
	//	newOrder := types.Order{
	//		ID:     book.OrderIDTrack,
	//		Creator: creator,
	//		Amount: amount,
	//		Price:  price,
	//	}
	//	newBook, orderID, err := types.AppendOrder(buyBook, creator, amount, price)
	//	buyBook, ok := newBook.(types.BuyOrderBook)
	//
	//	// Checks
	//	require.True(t, ok)
	//	require.NoError(t, err)
	//	require.Contains(t, book.Orders, newOrder)
	//	require.Equal(t, newOrder.ID, orderID)
	//}
	//require.Len(t, buyBook.Orders, 20)
	//require.True(t, sort.IsSorted(buyBook))
}

func TestUpdateOrderBook(t *testing.T) {
	inputBook := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
	}

	inputOrder := types.Order{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20}
	outputBook := []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 200, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
	}
	simulateUpdateSellOrderBook(t, inputBook, inputOrder, outputBook)

	inputOrder = types.Order{ID: 4, Creator: MockAccount("1"), Amount: 100, Price: 17}
	outputBook = []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 4, Creator: MockAccount("1"), Amount: 100, Price: 17},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
	}
	simulateUpdateSellOrderBook(t, inputBook, inputOrder, outputBook)

	inputOrder = types.Order{ID: 5, Creator: MockAccount("1"), Amount: 500, Price: 1}
	outputBook = []types.Order{
		{ID: 0, Creator: MockAccount("0"), Amount: 100, Price: 25},
		{ID: 1, Creator: MockAccount("1"), Amount: 100, Price: 20},
		{ID: 2, Creator: MockAccount("2"), Amount: 100, Price: 15},
		{ID: 3, Creator: MockAccount("3"), Amount: 100, Price: 10},
		{ID: 5, Creator: MockAccount("1"), Amount: 500, Price: 1},
	}
	simulateUpdateSellOrderBook(t, inputBook, inputOrder, outputBook)
}

func simulateUpdateSellOrderBook(t *testing.T, inputList []types.Order, inputOrder types.Order, outputList []types.Order) {
	inBook := OrderListToSellOrderBook(inputList)
	outBook := OrderListToSellOrderBook(outputList)

	require.True(t, sort.IsSorted(inBook))
	require.True(t, sort.IsSorted(outBook))

	simulatedBook := types.UpdateOrderBook(inBook, inputOrder)

	require.Equal(t, outBook, simulatedBook)
}
