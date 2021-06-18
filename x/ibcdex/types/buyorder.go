// x/ibcdex/types/buyorder.go
package types

import (
	"sort"
)

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
		orders := append(book.Orders, &order)
		copy(orders[i+1:], orders[i:])
		orders[i] = &order
		book.Orders = orders
	} else {
		book.Orders = append(book.Orders, &order)
	}

	return book
}

// GetOrder gets the order from an index
func (book BuyOrderBook) GetOrder(index int) (order Order, err error) {
	if index >= len(book.Orders) {
		return order, ErrOrderNotFound
	}

	return *book.Orders[index], nil
}

// SetOrder gets the order from an index
func (book BuyOrderBook) SetOrder(index int, order Order) (OrderBook, error) {
	if index >= len(book.Orders) {
		return book, ErrOrderNotFound
	}

	book.Orders[index] = &order

	return book, nil
}

// GetNextOrderID gets the ID of the next order to append
func (book BuyOrderBook) GetNextOrderID() int32 {
	return book.OrderIDTrack
}

// IncrementNextOrderID updates the ID tracker for buy orders
func (book BuyOrderBook) IncrementNextOrderID() OrderBook {
	// Even numbers to have different ID than buy orders
	book.OrderIDTrack += 2

	return book
}

// GetOrderFromID gets an order from the book from its id
func (book BuyOrderBook) GetOrderFromID(id int32) (Order, error) {
	for _, order := range book.Orders {
		if order.Id == id {
			return *order, nil
		}
	}
	return Order{}, ErrOrderNotFound
}

// RemoveOrderFromID removes an order from the book and keep it ordered
func (book BuyOrderBook) RemoveOrderFromID(id int32) (OrderBook, error) {
	for i, order := range book.Orders {
		if order.Id == id {
			book.Orders = append(book.Orders[:i], book.Orders[i+1:]...)
			return book, nil
		}
	}
	return book, ErrOrderNotFound
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
	purchase int32,
	match bool,
	filled bool,
) {
	remainingBuyOrder = order

	// No match if no order
	if book.Len() == 0 {
		return book, order, liquidatedSellOrder, purchase, false, false
	}

	// Check if match
	lowestAsk := book.Orders[book.Len()-1]
	if order.Price < lowestAsk.Price {
		return book, order, liquidatedSellOrder, purchase, false, false
	}

	liquidatedSellOrder = *lowestAsk

	// Check if buy order can be entirely filled
	if lowestAsk.Amount >= order.Amount {
		remainingBuyOrder.Amount = 0
		liquidatedSellOrder.Amount = order.Amount
		purchase = order.Amount

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
	remainingBuyOrder.Amount -= lowestAsk.Amount

	return book, remainingBuyOrder, liquidatedSellOrder, purchase, true, false
}

// FillBuyOrder try to fill the buy order with the order book and returns all the side effects
func FillBuyOrder(book SellOrderBook, order Order) (
	newBook SellOrderBook,
	remainingBuyOrder Order,
	liquidated []Order,
	purchase int32,
	filled bool,
) {
	var liquidatedList []Order
	totalPurchase := int32(0)
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
		liquidatedList = append(liquidatedList, liquidation)

		if filled {
			break
		}
	}

	return book, remainingBuyOrder, liquidatedList, totalPurchase, filled
}
