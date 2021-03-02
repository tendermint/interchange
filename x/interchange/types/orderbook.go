package types

import (
	"errors"
	"sort"
)

const (
	MaxAmount = uint64(100000)
	MaxPrice = uint64(100000)
)

var (
	ErrMaxAmount = errors.New("max amount reached")
	ErrMaxPrice = errors.New("max price reached")
	ErrZeroAmount = errors.New("amount is zero")
	ErrZeroPrice = errors.New("price is zero")
)

type CoinAllocation struct {
	amount uint64
	account Account
}

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

// ------------------------------ Sell Order ------------------------------

type SellOrder struct {
	ID uint64
	Seller Account
	Amount uint64
	Price uint64
}

type SellOrderBook struct {
	OrderIDTrack uint64
	AmountDenom string
	PriceDenom string
	Orders []SellOrder
}

// sort.Interface
func (book SellOrderBook) Len() int {
	return len(book.Orders)
}
func (book SellOrderBook) Less(i, j int) bool {
	return book.Orders[i].Price < book.Orders[j].Price
}
func (book SellOrderBook) Swap(i, j int) {
	book.Orders[i], book.Orders[j] = book.Orders[j], book.Orders[i]
}


func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrderBook {
	return SellOrderBook{
		OrderIDTrack: 0,
		AmountDenom: AmountDenom,
		PriceDenom: PriceDenom,
	}
}

func AppendSellOrder(book SellOrderBook, seller Account, amount uint64, price uint64) (SellOrderBook, uint64, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return book, 0, err
	}

	var order SellOrder
	order.ID = book.OrderIDTrack
	order.Seller = seller
	order.Amount = amount
	order.Price = price

	// Even numbers to have different ID than buy orders
	book.OrderIDTrack += 2

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

	return book, order.ID, nil
}

func FillSellOrder(book BuyOrderBook, order SellOrder) {

}

// ------------------------------ Buy Order ------------------------------

type BuyOrder struct {
	ID uint64
	Buyer Account
	Amount uint64
	Price uint64
}

type BuyOrderBook struct {
	OrderIDTrack uint64
	AmountDenom string
	PriceDenom string
	Orders []BuyOrder
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
		AmountDenom: AmountDenom,
		PriceDenom: PriceDenom,
	}
}

func AppendBuyOrder(book BuyOrderBook, buyer Account, amount uint64, price uint64) (BuyOrderBook, uint64, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return book, 0, err
	}

	var order BuyOrder
	order.ID = book.OrderIDTrack
	order.Buyer = buyer
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
