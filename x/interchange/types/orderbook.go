package types

import (
	"errors"
	"sort"
)

const (
	MaxAmount = uint64(100000)
	MaxPrice  = uint64(100000)
)

var (
	ErrMaxAmount     = errors.New("max amount reached")
	ErrMaxPrice      = errors.New("max price reached")
	ErrZeroAmount    = errors.New("amount is zero")
	ErrZeroPrice     = errors.New("price is zero")
	ErrOrderNotFound = errors.New("order not found")
)

type Order struct {
	ID      uint64
	Creator Account
	Amount  uint64
	Price   uint64
}

type OrderBook interface {
	sort.Interface
	InsertOrder(Order)
	GetOrder(int) (Order, error)
	SetOrder(int, Order) error
	GetNextOrderID() uint64
	IncrementNextOrderID()
	RemoveOrder() error
}

// UpdateOrderBook updates an order book with an order
// if the ID already exist, it append the amount to the existing order (without checking price)
// if it doesn't exist, the order is inserted
func UpdateOrderBook(book OrderBook, order Order) OrderBook {
	// Search of the order of the same ID
	i := sort.Search(book.Len(), func(i int) bool {
		tmp, _ := book.GetOrder(i)
		return tmp.ID == order.ID
	})

	// If order found
	if i < book.Len() {
		orderToUpdate, _ := book.GetOrder(i)
		orderToUpdate.Amount += order.Amount
		book.SetOrder(i, orderToUpdate)
	} else {
		book.InsertOrder(order)
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
func AppendOrder(book OrderBook, creator Account, amount uint64, price uint64) (OrderBook, uint64, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return book, 0, err
	}

	// Initialize the order
	var order Order
	order.ID = book.GetNextOrderID()
	order.Creator = creator
	order.Amount = amount
	order.Price = price

	// Increment ID tracker
	book.IncrementNextOrderID()

	// Insert the order
	book.InsertOrder(order)

	return book, order.ID, nil
}

// checkAmountAndPrice checks correct amount or price
func checkAmountAndPrice(amount uint64, price uint64) error {
	if amount == uint64(0) {
		return ErrZeroAmount
	}
	if amount > MaxAmount {
		return ErrMaxAmount
	}
	if price == uint64(0) {
		return ErrZeroPrice
	}
	if price > MaxPrice {
		return ErrMaxPrice
	}

	return nil
}
