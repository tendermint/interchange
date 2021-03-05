package types

import (
	"errors"
	"sort"
)

const (
	MaxAmount = int32(100000)
	MaxPrice  = int32(100000)
)

var (
	ErrMaxAmount     = errors.New("max amount reached")
	ErrMaxPrice      = errors.New("max price reached")
	ErrZeroAmount    = errors.New("amount is zero")
	ErrZeroPrice     = errors.New("price is zero")
	ErrOrderNotFound = errors.New("order not found")
)

//type Order struct {
//	ID      int32
//	Creator Account
//	Amount  int32
//	Price   int32
//}

type OrderBook interface {
	sort.Interface
	InsertOrder(Order) OrderBook
	GetOrder(int) (Order, error)
	SetOrder(int, Order) (OrderBook, error)
	GetNextOrderID() int32
	IncrementNextOrderID() OrderBook
	RemoveOrderFromID(int32) (OrderBook, error)
}

// UpdateOrderBook updates an order book with an order
// if the ID already exist, it append the amount to the existing order (without checking price)
// if it doesn't exist, the order is inserted
func UpdateOrderBook(book OrderBook, order Order) OrderBook {
	// Search of the order of the same ID
	var found bool
	var orderToUpdate Order
	var i int
	for i = book.Len() - 1; i >= 0; i-- {
		orderToUpdate, _ = book.GetOrder(i)
		if orderToUpdate.Id == order.Id {
			found = true
			break
		}
	}

	// If order found
	if found {
		orderToUpdate.Amount += order.Amount
		book, _ = book.SetOrder(i, orderToUpdate)
	} else {
		book = book.InsertOrder(order)
	}

	return book
}

// RestoreOrderBook restores the order book from a order book transition
func RestoreOrderBook(book OrderBook, liquidated []Order) OrderBook {
	// Restore all liquidation inside the order book
	for _, liquidation := range liquidated {
		book = UpdateOrderBook(book, liquidation)
	}

	return book
}

// AppendOrder initializes and appends a new order in a book from order information
func AppendOrder(book OrderBook, creator Account, amount int32, price int32) (OrderBook, int32, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return book, 0, err
	}

	// Initialize the order
	var order Order
	order.Id = book.GetNextOrderID()
	order.Creator = &creator
	order.Amount = amount
	order.Price = price

	// Increment ID tracker
	book = book.IncrementNextOrderID()

	// Insert the order
	book = book.InsertOrder(order)

	return book, order.Id, nil
}

// checkAmountAndPrice checks correct amount or price
func checkAmountAndPrice(amount int32, price int32) error {
	if amount == int32(0) {
		return ErrZeroAmount
	}
	if amount > MaxAmount {
		return ErrMaxAmount
	}
	if price == int32(0) {
		return ErrZeroPrice
	}
	if price > MaxPrice {
		return ErrMaxPrice
	}

	return nil
}
