package models_test

import (
	"testing"
	"time"

	"github.com/sergioazevedo/argentum_go/internal/models"
	"github.com/shopspring/decimal"
)

func TestTradeCantHaveNegativePrices(t *testing.T) {
	_, err := models.NewTrade(time.Now(), 10, decimal.New(-10, 0))
	if err == nil {
		t.Fatalf("A trade can't have negative price")
	}
}
