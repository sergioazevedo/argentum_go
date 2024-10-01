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
	volume   decimal.Decimal
}

func NewTrade(
	date time.Time,
	quantity decimal.Decimal,
	price decimal.Decimal,
	volume decimal.Decimal) (Trade, error) {

	if date.IsZero() {
		return Trade{}, errors.New("date is missing")
	}
	if quantity.LessThan(decimal.Zero) {
		return Trade{}, errors.New("quantitiy must be > 0.0")
	}
	if price.LessThan(decimal.Zero) {
		return Trade{}, errors.New("price must be >= 0.0")
	}

	if volume.LessThan(decimal.Zero) {
		return Trade{}, errors.New("volume must be >= 0.0")
	}

	return Trade{
		date:     date,
		quantity: quantity,
		price:    price,
		volume:   volume,
	}, nil
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
	return trade.volume
}
