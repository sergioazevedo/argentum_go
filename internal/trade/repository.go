package trade

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/sergioazevedo/argentum_go/internal/api/binance"
	"github.com/sergioazevedo/argentum_go/internal/api/kraken"
)

type Repository struct {
	binance *binance.APIClient
	kraken  *kraken.APIClient
}

func NewRepository(binanceClient *binance.APIClient, krakenClient *kraken.APIClient) *Repository {
	return &Repository{
		binance: binanceClient,
		kraken:  krakenClient,
	}
}

func (r Repository) FetchRecentTradesFromBinance(pair string, limit int16) ([]Trade, error) {
	binanceTrades, _ := r.binance.FetchRecentTrades(pair, uint16(limit))

	return buildFromBinanceTrades(binanceTrades), nil
}

func (r Repository) FetchRecentTradesFromKraken(pair string, limit int16) ([]Trade, error) {
	krakenTrades, _ := r.kraken.FetchRecentTrades(pair, limit)

	list := buildFromKrakenTrades(krakenTrades)

	return list, nil
}

func buildFromKrakenTrades(krakenTrades []kraken.Trade) []Trade {
	list := make([]Trade, 0, len(krakenTrades))
	for _, v := range krakenTrades {
		list = append(list, Trade{
			Date:     v.Date,
			Price:    v.Price,
			Volume:   v.Volume,
			Quantity: v.Quantity,
		})
	}

	return list
}

func buildFromBinanceTrades(binanceTrades []binance.Trade) []Trade {
	list := make([]Trade, 0, len(binanceTrades))
	for _, v := range binanceTrades {
		_price, _ := decimal.NewFromString(v.Price)
		_qty, _ := decimal.NewFromString(v.Qty)

		list = append(list, Trade{
			Date:   time.UnixMilli(int64(v.Time)),
			Price:  _price,
			Volume: _qty,
		})
	}

	return list
}
