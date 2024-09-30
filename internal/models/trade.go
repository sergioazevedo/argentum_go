package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade interface {
	Date() time.Time
	Quantity() decimal.Decimal
	Price() decimal.Decimal
	Volume() decimal.Decimal
}
