package models

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	date     time.Time
	quantity int32
	price    decimal.Decimal
}

func NewTrade(date time.Time, quantity int32, price decimal.Decimal) (Trade, error) {
	if date.IsZero() {
		return Trade{}, errors.New("date is missing")
	}
	if quantity < 0 {
		return Trade{}, errors.New("quantitiy must be > 0")
	}
	if price.LessThan(decimal.New(0, 0)) {
		return Trade{}, errors.New("price must be >= 0.0")
	}

	return Trade{
		date:     date,
		quantity: quantity,
		price:    price,
	}, nil
}

func (trade Trade) Date() time.Time {
	return trade.date
}

func (trade Trade) Quantity() int32 {
	return trade.quantity
}

func (trade Trade) Price() decimal.Decimal {
	return trade.price
}

func (trade Trade) Volume() decimal.Decimal {
	return trade.price.Mul(decimal.NewFromInt32(trade.quantity))
}
