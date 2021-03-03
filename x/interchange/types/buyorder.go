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
	return book.Orders[i].Price > book.Orders[j].Price
}
func (book BuyOrderBook) Swap(i, j int) {
	book.Orders[i], book.Orders[j] = book.Orders[j], book.Orders[i]
}

func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrderBook {
	return BuyOrderBook{
		OrderIDTrack: 1,
		AmountDenom:  AmountDenom,
		PriceDenom:   PriceDenom,
	}
}

func AppendBuyOrder(book BuyOrderBook, buyer Account, amount uint64, price uint64) (BuyOrderBook, uint64, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return book, 0, err
	}

	var order Order
	order.ID = book.OrderIDTrack
	order.Creator = buyer
	order.Amount = amount
	order.Price = price

	// Odd numbers to have different ID than buy orders
	book.OrderIDTrack += 2

	// Insert the order in the decreasing order
	if len(book.Orders) > 0 {
		i := sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price < order.Price })
		orders := append(book.Orders, order)
		copy(orders[i+1:], orders[i:])
		orders[i] = order
		book.Orders = orders
	} else {
		book.Orders = append(book.Orders, order)
	}

	return book, order.ID, nil
}
