package main

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/sergioazevedo/argentum_go/internal/api/binance"
	"github.com/sergioazevedo/argentum_go/internal/api/kraken"
	candlestick "github.com/sergioazevedo/argentum_go/internal/candlelstick"
	"github.com/sergioazevedo/argentum_go/internal/trade"
)

type xAxisData []string
type seriesData []opts.KlineData

func main() {
	var (
		binanceClient = binance.NewClient("")
		krakenClient  = kraken.NewClient("")
		repository    = trade.NewRepository(binanceClient, krakenClient)
	)

	fmt.Println("running...:")
	trades, err := repository.FetchRecentTradesFromKraken(
		"XXBTZEUR",
		1000,
	)

	if err != nil {
		fmt.Println("Error fetching trades:", err)
		return
	}

	binanceTrades, err := repository.FetchRecentTradesFromBinance(
		"BTCEUR",
		1000,
	)
	if err != nil {
		fmt.Println("Error fetching trades:", err)
		return
	}

	chartOptions := []charts.GlobalOpts{
		charts.WithTitleOpts(opts.Title{
			Title: "Latest trades per minute - BTC/EUR",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: opts.Bool(true),
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	}
	candlesticks := candlestick.FromTrades(trades, "1m")
	krakenChart := buildChart(candlesticks, chartOptions, "Kraken")

	binanceCandlesticks := candlestick.FromTrades(binanceTrades, "1m")
	binanceChart := buildChart(binanceCandlesticks, chartOptions, "Binance")

	page := components.NewPage()
	page.AddCharts(krakenChart, binanceChart)

	f, err := os.Create("chart.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func buildChart(candlesticks []candlestick.Candlestick, chartOptions []charts.GlobalOpts, name string) *charts.Kline {
	chart := charts.NewKLine()
	chart.SetGlobalOptions(chartOptions...)
	xAxisData, seriesData := mapCandlestickToChartData(candlesticks)
	chart.SetXAxis(xAxisData).AddSeries(name, seriesData,
		// fix candlestick colors
		charts.WithItemStyleOpts(opts.ItemStyle{
			Color:        "#47b262",
			Color0:       "#eb5454",
			BorderColor:  "#47b262",
			BorderColor0: "#eb5454",
		}))

	return chart
}

func mapCandlestickToChartData(candlesticks []candlestick.Candlestick) (xAxisData, seriesData) {
	xAxisData := []string{}
	seriesData := []opts.KlineData{}
	for _, c := range candlesticks {
		arr := [4]float32{}
		arr[0] = float32(c.Open.InexactFloat64())
		arr[1] = float32(c.Close.InexactFloat64())
		arr[2] = float32(c.Low.InexactFloat64())
		arr[3] = float32(c.High.InexactFloat64())

		xAxisData = append(xAxisData, c.Date.Format("2006-01-02 15:04:05"))
		seriesData = append(seriesData, opts.KlineData{Value: arr})
	}

	return xAxisData, seriesData
}
