package types

import "sort"

type BuyOrderBook struct {
	OrderIDTrack uint64
	AmountDenom  string
	PriceDenom   string
	Orders       []Order
}

// sort.Interface
func (book BuyOrderBook) Len() int {
	return len(book.Orders)
}
func (book BuyOrderBook) Less(i, j int) bool {
	// Buy orders are decreasingly sorted
	return book.Orders[i].Price < book.Orders[j].Price
}
func (book BuyOrderBook) Swap(i, j int) {
	book.Orders[i], book.Orders[j] = book.Orders[j], book.Orders[i]
}

// InsertOrder inserts the order in the increasing order in the book
// it doesn't set the ID or check if the ID already exist
func (book BuyOrderBook) InsertOrder(order Order) OrderBook {
	// Insert the order in the increasing order
	if len(book.Orders) > 0 {
		i := sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price > order.Price })
		orders := append(book.Orders, order)
		copy(orders[i+1:], orders[i:])
		orders[i] = order
		book.Orders = orders
	} else {
		book.Orders = append(book.Orders, order)
	}

	return book
}

// GetOrder gets the order from an index
func (book BuyOrderBook) GetOrder(index int) (order Order, err error) {
	if index >= len(book.Orders) {
		return order, ErrOrderNotFound
	}

	return book.Orders[index], nil
}

// SetOrder gets the order from an index
func (book BuyOrderBook) SetOrder(index int, order Order) (OrderBook, error) {
	if index >= len(book.Orders) {
		return book, ErrOrderNotFound
	}

	book.Orders[index] = order

	return book, nil
}

// GetNextOrderID gets the ID of the next order to append
func (book BuyOrderBook) GetNextOrderID() uint64 {
	return book.OrderIDTrack
}

// IncrementNextOrderID updates the ID tracker for buy orders
func (book BuyOrderBook) IncrementNextOrderID() OrderBook {
	// Even numbers to have different ID than buy orders
	book.OrderIDTrack += 2

	return book
}

func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrderBook {
	return BuyOrderBook{
		OrderIDTrack: 1,
		AmountDenom:  AmountDenom,
		PriceDenom:   PriceDenom,
	}
}

// RemoveOrder removes an order from the book and keep it ordered
func (book BuyOrderBook) RemoveOrder(index int) (OrderBook, error) {
	if index >= len(book.Orders) {
		return book, ErrOrderNotFound
	}

	book.Orders = append(book.Orders[:index], book.Orders[index+1:]...)
	return book, nil
}