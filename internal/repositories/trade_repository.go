package repositories

import (
	"github.com/sergioazevedo/argentum_go/internal/models"
	services "github.com/sergioazevedo/argentum_go/internal/services/binance"
)

type TradeRepository struct{}

func (r *TradeRepository) FetchRecentTradesFromBinance(pair string, limit int16) ([]models.Trade, error) {
	binance := services.BinanceAPIClient{}
	binanceTrades, _ := binance.FetchRecentTrades(pair, uint16(limit))

	return castToTrade[models.BinanceTrade](binanceTrades), nil
}

func castToTrade[T models.Trade](tradesList []T) []models.Trade {
	result := make([]models.Trade, 0, len(tradesList))
	for _, v := range tradesList {
		result = append(result, v)
	}

	return result
}
