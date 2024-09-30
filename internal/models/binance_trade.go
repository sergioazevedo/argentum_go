package models

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type BinanceTrade struct {
	date     time.Time
	quantity decimal.Decimal
	price    decimal.Decimal
}

func (b BinanceTrade) Date() time.Time {
	return b.date
}

func (b BinanceTrade) Quantity() decimal.Decimal {
	return b.quantity
}

func (b BinanceTrade) Price() decimal.Decimal {
	return b.price
}

func (b BinanceTrade) Volume() decimal.Decimal {
	return b.Price().Mul(b.Quantity())
}

func (b *BinanceTrade) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	_time := int64(v["time"].(float64))
	_date := time.UnixMilli(_time)
	_qty, err := decimal.NewFromString(v["qty"].(string))
	if err != nil {
		return err
	}
	_price, err := decimal.NewFromString(v["price"].(string))
	if err != nil {
		return err
	}

	b.date = _date
	b.price = _price
	b.quantity = _qty

	return nil
}
