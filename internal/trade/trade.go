package trade

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	Date     time.Time
	Quantity decimal.Decimal
	Price    decimal.Decimal
	Volume   decimal.Decimal
}
