// x/ibcdex/types/order_book.go
package types

import (
	"errors"
	"sort"
)

const (
	MaxAmount = int32(100000)
	MaxPrice  = int32(100000)
)

type Ordering int

const (
	Increasing Ordering = iota
	Decreasing
)

var (
	ErrMaxAmount     = errors.New("max amount reached")
	ErrMaxPrice      = errors.New("max price reached")
	ErrZeroAmount    = errors.New("amount is zero")
	ErrZeroPrice     = errors.New("price is zero")
	ErrOrderNotFound = errors.New("order not found")
)

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

func NewOrderBook() OrderBook {
	return OrderBook{
		IdCount: 0,
	}
}

// GetOrder gets the order from an index
func (book OrderBook) GetOrder(index int) (order Order, err error) {
	if index >= len(book.Orders) {
		return order, ErrOrderNotFound
	}

	return *book.Orders[index], nil
}

// GetNextOrderID gets the ID of the next order to append
func (book OrderBook) GetNextOrderID() int32 {
	return book.IdCount
}

// GetOrderFromID gets an order from the book from its id
func (book OrderBook) GetOrderFromID(id int32) (Order, error) {
	for _, order := range book.Orders {
		if order.Id == id {
			return *order, nil
		}
	}
	return Order{}, ErrOrderNotFound
}

// SetOrder gets the order from an index
func (book *OrderBook) SetOrder(index int, order Order) error {
	if index >= len(book.Orders) {
		return ErrOrderNotFound
	}

	book.Orders[index] = &order

	return nil
}

// IncrementNextOrderID updates the ID count for orders
func (book *OrderBook) IncrementNextOrderID() {
	// Even numbers to have different ID than buy orders
	book.IdCount++
}

// RemoveOrderFromID removes an order from the book and keep it ordered
func (book* OrderBook) RemoveOrderFromID(id int32) error {
	for i, order := range book.Orders {
		if order.Id == id {
			book.Orders = append(book.Orders[:i], book.Orders[i+1:]...)
			return nil
		}
	}
	return ErrOrderNotFound
}

// AppendOrder initializes and appends a new order in a book from order information
func (book *OrderBook) appendOrder(creator string, amount int32, price int32, ordering Ordering) (int32, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return 0, err
	}

	// Initialize the order
	var order Order
	order.Id = book.GetNextOrderID()
	order.Creator = creator
	order.Amount = amount
	order.Price = price

	// Increment ID tracker
	book.IncrementNextOrderID()

	// Insert the order
	book.insertOrder(order, ordering)

	return order.Id, nil
}

// insertOrder inserts the order in the book with the provided order
func (book *OrderBook) insertOrder(order Order, ordering Ordering) {
	if len(book.Orders) > 0 {
		var i int

		// get the index of the new order depending on the provided ordering
		if ordering == Increasing {
			i = sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price > order.Price })
		} else {
			i = sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price < order.Price })
		}

		// insert order
		orders := append(book.Orders, &order)
		copy(orders[i+1:], orders[i:])
		orders[i] = &order
		book.Orders = orders
	} else {
		book.Orders = append(book.Orders, &order)
	}
}
