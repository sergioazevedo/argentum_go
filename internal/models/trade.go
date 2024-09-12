package models

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	date     time.Time
	quantity int
	price    decimal.Decimal
}

func NewTrade(date time.Time, quantity int, price decimal.Decimal) (Trade, error) {
	if price.LessThan(decimal.New(0, 0)) {
		return Trade{}, errors.New("price value must be >= 0")
	}

	return Trade{
		date:     date,
		quantity: quantity,
		price:    price,
	}, nil
}
