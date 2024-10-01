package repositories

import (
	"time"

	"github.com/sergioazevedo/argentum_go/internal/models"
	services "github.com/sergioazevedo/argentum_go/internal/services/binance"
	"github.com/sergioazevedo/argentum_go/internal/services/kraken"
	"github.com/shopspring/decimal"
)

type TradeRepository struct{}

func (r TradeRepository) FetchRecentTradesFromBinance(pair string, limit int16) ([]models.Trade, error) {
	binance := services.BinanceAPIClient{}
	binanceTrades, _ := binance.FetchRecentTrades(pair, uint16(limit))

	return buildFromBinanceTrades(binanceTrades), nil
}

func (r TradeRepository) FetchRecentTradesFromKraken(pair string, limit int16) ([]models.Trade, error) {
	kraken := kraken.APIClient{}
	krakenTrades, _ := kraken.FetchRecentTrades(pair, limit)

	list := buildFromKrakenTrades(krakenTrades)

	return list, nil
}

func buildFromKrakenTrades(krakenTrades kraken.KrakenTrades) []models.Trade {
	list := make([]models.Trade, 0, len(krakenTrades))
	for _, v := range krakenTrades {
		_time := int64(v[2].(float64))
		_price, _ := decimal.NewFromString(v[0].(string))
		_volume, _ := decimal.NewFromString(v[1].(string))

		trade, err := models.NewTrade(
			time.Unix(_time, 0),
			_price,
			_price.Div(_volume),
			_volume,
		)

		if err == nil {
			list = append(list, trade)
		}
	}

	return list
}

func buildFromBinanceTrades(binanceTrades []services.BinanceTrade) []models.Trade {
	list := make([]models.Trade, 0, len(binanceTrades))
	for _, v := range binanceTrades {
		_price, _ := decimal.NewFromString(v.Price)
		_qty, _ := decimal.NewFromString(v.Qty)

		trade, err := models.NewTrade(
			time.UnixMilli(int64(v.Time)),
			_price,
			_qty,
			_price.Mul(_qty),
		)

		if err == nil {
			list = append(list, trade)
		}
	}

	return list
}
