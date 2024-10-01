package main

import (
	"fmt"

	"github.com/sergioazevedo/argentum_go/internal/repositories"
)

// func main() {
// 	fmt.Println("running...:")
// 	repository := repositories.TradeRepository{}
// 	binanceTrades, _ := repository.FetchRecentTradesFromBinance(
// 		"BTCEUR",
// 		10,
// 	)
// 	fmt.Println("Trades::")
// 	fmt.Println(binanceTrades[0].Date().String())

// 	c := models.CadlesticksFrom(binanceTrades, "1s")
// 	fmt.Println("Candlesticks:")
// 	fmt.Println(c)
// }

func main() {
	fmt.Println("running...:")
	repository := repositories.TradeRepository{}
	trades, _ := repository.FetchRecentTradesFromKraken(
		"XXBTZEUR",
		10,
	)
	fmt.Println("Trades::")
	fmt.Println(trades[0].Date().String())

	// c := models.CadlesticksFrom(trades, "1s")
	// fmt.Println("Candlesticks:")
	// fmt.Println(c)
}
