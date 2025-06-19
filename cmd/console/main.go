package main

import (
	"fmt"
	"time"

	"github.com/sergioazevedo/argentum_go/internal/api/binance"
	"github.com/sergioazevedo/argentum_go/internal/api/kraken"
	"github.com/sergioazevedo/argentum_go/internal/trade"
)

func main() {
	var (
		binanceClient = binance.NewClient("")
		krakenClient  = kraken.NewClient("")
		repository    = trade.NewRepository(binanceClient, krakenClient)
	)

	fmt.Println("running...:")
	trades, err := repository.FetchRecentTradesFromKraken(
		"XXBTZEUR",
		10,
	)

	if err != nil {
		fmt.Println("Error fetching trades:", err)
		return
	}

	fmt.Println("Kraken Trades::")
	fmt.Println(trades[0].Date.Format(time.RFC3339))

	binanceTrades, err := repository.FetchRecentTradesFromBinance(
		"BTCEUR",
		10,
	)
	if err != nil {
		fmt.Println("Error fetching trades:", err)
		return
	}

	fmt.Println("Binance Trades::")
	fmt.Println(binanceTrades[0].Date.Format(time.RFC3339))

	// c := models.CadlesticksFrom(trades, "1s")
	// fmt.Println("Candlesticks:")
	// fmt.Println(c)
}
