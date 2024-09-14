package models_test

import (
	"testing"
	"time"

	"github.com/sergioazevedo/argentum_go/internal/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestMustCreateATradeWithNoError(t *testing.T) {
	now := time.Now()
	trade, err := models.NewTrade(now, 10, decimal.New(10, 22))
	assert.Nil(t, err)
	assert.Equal(t, trade.Quantity(), int32(10))
	assert.True(t, trade.Date().Equal(now))
	assert.True(t, trade.Price().Equal(decimal.New(10, 22)))
}

func TestTradeVolume(t *testing.T) {
	trade, _ := models.NewTrade(time.Now(), 10, decimal.New(20, 0))
	assert.Equal(t, trade.Volume(), decimal.NewFromInt32(200))
}

func TestTradeCantHaveNegativeQuantity(t *testing.T) {
	trade, err := models.NewTrade(time.Now(), -10, decimal.New(-10, 0))

	assert.Empty(t, trade)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "quantitiy must be > 0")
}

func TestTradeCantHaveNegativePrices(t *testing.T) {
	trade, err := models.NewTrade(time.Now(), 10, decimal.New(-10, 0))

	assert.Empty(t, trade)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "price must be >= 0.0")
}

func TestTradeMustHaveADate(t *testing.T) {
	trade, err := models.NewTrade(time.Time{}, 10, decimal.New(10, 22))

	assert.Empty(t, trade)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "date is missing")
}
