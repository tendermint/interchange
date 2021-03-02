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
func (sob SellOrderBook) Len() int {
	return len(sob.Orders)
}
func (sob SellOrderBook) Less(i, j int) bool {
	return sob.Orders[i].Price < sob.Orders[j].Price
}
func (sob SellOrderBook) Swap(i, j int) {
	sob.Orders[i], sob.Orders[j] = sob.Orders[j], sob.Orders[i]
}


func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrderBook {
	return SellOrderBook{
		OrderIDTrack: 0,
		AmountDenom: AmountDenom,
		PriceDenom: PriceDenom,
	}
}

func AppendSellOrder(sob SellOrderBook, seller Account, amount uint64, price uint64) (SellOrderBook, uint64, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return sob, 0, err
	}

	var so SellOrder
	so.ID = sob.OrderIDTrack
	so.Seller = seller
	so.Amount = amount
	so.Price = price

	// Even numbers to have different ID than buy orders
	sob.OrderIDTrack += 2

	// Insert the order in the increasing order
	if len(sob.Orders) > 0 {
		i := sort.Search(len(sob.Orders), func(i int) bool { return sob.Orders[i].Price > so.Price })
		orders := append(sob.Orders, so)
		copy(orders[i+1:], orders[i:])
		orders[i] = so
		sob.Orders = orders
	} else {
		sob.Orders = append(sob.Orders, so)
	}

	return sob, so.ID, nil
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
func (bob BuyOrderBook) Len() int {
	return len(bob.Orders)
}
func (bob BuyOrderBook) Less(i, j int) bool {
	// Buy orders are decreasingly sorted
	return bob.Orders[i].Price > bob.Orders[j].Price
}
func (bob BuyOrderBook) Swap(i, j int) {
	bob.Orders[i], bob.Orders[j] = bob.Orders[j], bob.Orders[i]
}

func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrderBook {
	return BuyOrderBook{
		OrderIDTrack: 1,
		AmountDenom: AmountDenom,
		PriceDenom: PriceDenom,
	}
}

func AppendBuyOrder(bob BuyOrderBook, buyer Account, amount uint64, price uint64) (BuyOrderBook, uint64, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return bob, 0, err
	}

	var bo BuyOrder
	bo.ID = bob.OrderIDTrack
	bo.Buyer = buyer
	bo.Amount = amount
	bo.Price = price

	// Odd numbers to have different ID than buy orders
	bob.OrderIDTrack += 2

	// Insert the order in the decreasing order
	if len(bob.Orders) > 0 {
		i := sort.Search(len(bob.Orders), func(i int) bool { return bob.Orders[i].Price < bo.Price })
		orders := append(bob.Orders, bo)
		copy(orders[i+1:], orders[i:])
		orders[i] = bo
		bob.Orders = orders
	} else {
		bob.Orders = append(bob.Orders, bo)
	}

	return bob, bo.ID, nil
}
