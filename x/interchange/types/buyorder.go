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

// RemoveOrder removes an order from the book and keep it ordered
func (book BuyOrderBook) RemoveOrder(index int) (OrderBook, error) {
	if index >= len(book.Orders) {
		return book, ErrOrderNotFound
	}

	book.Orders = append(book.Orders[:index], book.Orders[index+1:]...)
	return book, nil
}


func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrderBook {
	return BuyOrderBook{
		OrderIDTrack: 1,
		AmountDenom:  AmountDenom,
		PriceDenom:   PriceDenom,
	}
}

// LiquidateFromBuyOrder liquidates the first sell order of the book from the buy order
// if no match is found, return false for match
func LiquidateFromBuyOrder(book SellOrderBook, order Order) (
	newBook SellOrderBook,
	remainingBuyOrder Order,
	liquidatedSellOrder Order,
	purchase uint64,
	match bool,
	filled bool,
) {
	// No match if no order
	if book.Len() == 0 {
		return newBook, order, liquidatedSellOrder, purchase, false, false
	}

	// Check if match
	lowestAsk := book.Orders[book.Len()-1]
	if order.Price < lowestAsk.Price {
		return newBook, order, liquidatedSellOrder, purchase, false, false
	}

	liquidatedSellOrder = lowestAsk

	// Check if buy order can be entirely filled
	if lowestAsk.Amount >= order.Amount {
		purchase = order.Amount
		liquidatedSellOrder.Amount = order.Amount

		// Remove lowest ask if it has been entirely liquidated
		lowestAsk.Amount -= order.Amount
		if lowestAsk.Amount == 0 {
			book.Orders = book.Orders[:book.Len()-1]
		} else {
			book.Orders[book.Len()-1] = lowestAsk
		}

		return book, remainingBuyOrder, liquidatedSellOrder, purchase, true, true
	}

	// Not entirely filled
	purchase = lowestAsk.Amount
	book.Orders = book.Orders[:book.Len()-1]
	remainingBuyOrder = order
	remainingBuyOrder.Amount -= lowestAsk.Amount

	return book, remainingBuyOrder, liquidatedSellOrder, purchase, true, false
}

// FillBuyOrder try to fill the buy order with the order book and returns all the side effects
func FillBuyOrder(book SellOrderBook, order Order) (
	newBook SellOrderBook,
	remainingBuyOrder Order,
	liquidated []Order,
	purchase uint64,
	filled bool,
) {
	totalPurchase := uint64(0)
	remainingBuyOrder = order

	// Liquidate as long as there is match
	for {
		var match bool
		var liquidation Order
		book, remainingBuyOrder, liquidation, purchase, match, filled = LiquidateFromBuyOrder(
			book,
			remainingBuyOrder,
		)
		if !match {
			break
		}

		// Update gains
		totalPurchase += purchase

		// Update liquidated
		liquidated = append(liquidated, liquidation)

		if filled {
			break
		}
	}

	return book, remainingBuyOrder, liquidated, totalPurchase, filled
}
