package models

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	date     time.Time
	quantity decimal.Decimal
	price    decimal.Decimal
}

type BinanceTradeJsonData struct {
	Price string
	Qty   string
	Time  int64
	// IsBuyerMaker bool
}

type TradesJSONData []BinanceTradeJsonData

func NewTrade(date time.Time, quantity decimal.Decimal, price decimal.Decimal) (Trade, error) {

	if date.IsZero() {
		return Trade{}, errors.New("date is missing")
	}
	if quantity.LessThan(decimal.Zero) {
		return Trade{}, errors.New("quantitiy must be > 0")
	}
	if price.LessThan(decimal.Zero) {
		return Trade{}, errors.New("price must be >= 0.0")
	}

	return Trade{
		date:     date,
		quantity: quantity,
		price:    price,
	}, nil
}

func newTradeFromString(date int64, quantity string, price string) (Trade, error) {
	_time := time.UnixMilli(date)
	_qty, _ := decimal.NewFromString(quantity)
	_price, _ := decimal.NewFromString(price)

	return NewTrade(_time, _qty, _price)
}

func TradesFromJSON(tradesJson TradesJSONData) []Trade {
	if len(tradesJson) == 0 {
		return nil
	}

	list := make([]Trade, 0, 30)
	for _, json := range tradesJson {
		// skips if the json data is empty
		if (BinanceTradeJsonData{}) != json {
			trade, err := newTradeFromString(json.Time, json.Qty, json.Price)
			if err == nil {
				list = append(list, trade)
			}
		}
	}
	return list
}

func (trade Trade) Date() time.Time {
	return trade.date
}

func (trade Trade) Quantity() decimal.Decimal {
	return trade.quantity
}

func (trade Trade) Price() decimal.Decimal {
	return trade.price
}

func (trade Trade) Volume() decimal.Decimal {
	return trade.price.Mul(trade.quantity)
}
