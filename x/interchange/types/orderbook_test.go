package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/interchange/types"
	"math/rand"
	"sort"
	"testing"
)

// ------------------------------ Utils ------------------------------

func GenAmount() uint64 {
	return uint64(rand.Intn(int(types.MaxAmount))+1)
}

func GenPrice() uint64 {
	return uint64(rand.Intn(int(types.MaxPrice))+1)
}

func GenPair() (string, string) {
	return GenString(10), GenString(10)
}

func GenOrder() (types.Account, uint64, uint64)  {
	return GenLocalAccount(), GenAmount(), GenPrice()
}

// ------------------------------ Sell Order ------------------------------

func TestNewSellOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	book := types.NewSellOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint64(0), book.OrderIDTrack)
	require.Equal(t, amountDenom, book.AmountDenom)
	require.Equal(t, priceDenom, book.PriceDenom)
	require.Empty(t, book.Orders)
}

func TestAppendSellOrder(t *testing.T) {
	book := types.NewSellOrderBook(GenPair())

	for i := 0; i < 20; i++ {
		// Append a new sell order
		seller, amount, price := GenOrder()
		newOrder := types.SellOrder{
			ID: book.OrderIDTrack,
			Seller: seller,
			Amount: amount,
			Price: price,
		}
		newBook, orderID, err := types.AppendSellOrder(book, seller, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, newBook.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)

		book = newBook
	}
	require.Len(t, book.Orders, 20)
	require.True(t, sort.IsSorted(book))

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, _, err := types.AppendSellOrder(book, seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, _, err = types.AppendSellOrder(book, seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, _, err = types.AppendSellOrder(book, seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, _, err = types.AppendSellOrder(book, seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)
}

// ------------------------------ Buy Order ------------------------------

func TestNewBuyOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	book := types.NewBuyOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint64(1), book.OrderIDTrack)
	require.Equal(t, amountDenom, book.AmountDenom)
	require.Equal(t, priceDenom, book.PriceDenom)
	require.Empty(t, book.Orders)
}

func TestAppendBuyOrder(t *testing.T) {
	book := types.NewBuyOrderBook(GenPair())

	for i := 0; i < 20; i++ {
		// Append a new sell order
		buyer, amount, price := GenOrder()
		newOrder := types.BuyOrder{
			ID: book.OrderIDTrack,
			Buyer: buyer,
			Amount: amount,
			Price: price,
		}
		newBook, orderID, err := types.AppendBuyOrder(book, buyer, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, newBook.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)

		book = newBook
	}
	require.Len(t, book.Orders, 20)
	require.True(t, sort.IsSorted(book))

	// Prevent zero amount
	buyer, amount, price := GenOrder()
	_, _, err := types.AppendBuyOrder(book, buyer, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, _, err = types.AppendBuyOrder(book, buyer, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, _, err = types.AppendBuyOrder(book, buyer, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, _, err = types.AppendBuyOrder(book, buyer, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)
}