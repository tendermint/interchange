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
			ID:     book.OrderIDTrack,
			Buyer:  buyer,
			Amount: amount,
			Price:  price,
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
