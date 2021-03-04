package types

import "sort"

type SellOrderBook struct {
	OrderIDTrack uint32
	AmountDenom  string
	PriceDenom   string
	Orders       []Order
}

// sort.Interface
func (book SellOrderBook) Len() int {
	return len(book.Orders)
}
func (book SellOrderBook) Less(i, j int) bool {
	return book.Orders[i].Price > book.Orders[j].Price
}
func (book SellOrderBook) Swap(i, j int) {
	book.Orders[i], book.Orders[j] = book.Orders[j], book.Orders[i]
}

// InsertOrder inserts the order in the decreasing order in the book
// it doesn't set the ID or check if the ID already exist
func (book SellOrderBook) InsertOrder(order Order) OrderBook {
	// Insert the order in the increasing order
	if len(book.Orders) > 0 {
		i := sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price < order.Price })
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
func (book SellOrderBook) GetOrder(index int) (order Order, err error) {
	if index >= len(book.Orders) {
		return order, ErrOrderNotFound
	}

	return book.Orders[index], nil
}

// SetOrder gets the order from an index
func (book SellOrderBook) SetOrder(index int, order Order) (OrderBook, error) {
	if index >= len(book.Orders) {
		return book, ErrOrderNotFound
	}

	book.Orders[index] = order

	return book, nil
}

// GetNextOrderID gets the ID of the next order to append
func (book SellOrderBook) GetNextOrderID() uint32 {
	return book.OrderIDTrack
}

// IncrementNextOrderID updates the ID tracker for sell orders
func (book SellOrderBook) IncrementNextOrderID() OrderBook {
	// Even numbers to have different ID than buy orders
	book.OrderIDTrack += 2

	return book
}

// RemoveOrder removes an order from the book and keep it ordered
func (book SellOrderBook) RemoveOrder(index int) (OrderBook, error) {
	if index >= len(book.Orders) {
		return book, ErrOrderNotFound
	}

	book.Orders = append(book.Orders[:index], book.Orders[index+1:]...)
	return book, nil
}

func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrderBook {
	return SellOrderBook{
		OrderIDTrack: 0,
		AmountDenom:  AmountDenom,
		PriceDenom:   PriceDenom,
	}
}

// LiquidateFromSellOrder liquidates the first buy order of the book from the sell order
// if no match is found, return false for match
func LiquidateFromSellOrder(book BuyOrderBook, order Order) (
	newBook BuyOrderBook,
	remainingSellOrder Order,
	liquidatedBuyOrder Order,
	gain uint32,
	match bool,
	filled bool,
) {
	remainingSellOrder = order

	// No match if no order
	if book.Len() == 0 {
		return book, order, liquidatedBuyOrder, gain, false, false
	}

	// Check if match
	highestBid := book.Orders[book.Len()-1]
	if order.Price > highestBid.Price {
		return book, order, liquidatedBuyOrder, gain, false, false
	}

	liquidatedBuyOrder = highestBid

	// Check if sell order can be entirely filled
	if highestBid.Amount >= order.Amount {
		remainingSellOrder.Amount = 0
		liquidatedBuyOrder.Amount = order.Amount
		gain = order.Amount * highestBid.Price

		// Remove highest bid if it has been entirely liquidated
		highestBid.Amount -= order.Amount
		if highestBid.Amount == 0 {
			book.Orders = book.Orders[:book.Len()-1]
		} else {
			book.Orders[book.Len()-1] = highestBid
		}
		return book, remainingSellOrder, liquidatedBuyOrder, gain, true, true
	}

	// Not entirely filled
	gain = highestBid.Amount * highestBid.Price
	book.Orders = book.Orders[:book.Len()-1]
	remainingSellOrder.Amount -= highestBid.Amount

	return book, remainingSellOrder, liquidatedBuyOrder, gain, true, false
}

// FillSellOrder try to fill the sell order with the order book and returns all the side effects
func FillSellOrder(book BuyOrderBook, order Order) (
	newBook BuyOrderBook,
	remainingSellOrder Order,
	liquidated []Order,
	gain uint32,
	filled bool,
) {
	var liquidatedList []Order
	totalGain := uint32(0)
	remainingSellOrder = order

	// Liquidate as long as there is match
	for {
		var match bool
		var liquidation Order
		book, remainingSellOrder, liquidation, gain, match, filled = LiquidateFromSellOrder(
			book,
			remainingSellOrder,
		)
		if !match {
			break
		}

		// Update gains
		totalGain += gain

		// Update liquidated
		liquidatedList = append(liquidatedList, liquidation)

		if filled {
			break
		}
	}

	return book, remainingSellOrder, liquidatedList, totalGain, filled
}
