package main

import (
	"./config"

	"fmt"
	"io/ioutil"
	"net/http"
	// "os"

	"encoding/json"
	ui "github.com/gizak/termui"
	// "fsnotify"
)

type ApiCoin struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	VolumeUsd        string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
}

type ApiResponse []ApiCoin

var headers = []string{
	"Name",
	"Symbol",
	"Price in USD",
	"Price in BTC",
	"% Change (1h)",
	"% Change (24h)",
}

func includes(coins []string, coin string) bool {
	for _, n := range coins {
		if n == coin {
			return true
		}
	}
	return false
}

func getCoins(holdings []string) [][]string {
	rows := [][]string{headers}
	res, err := http.Get("https://api.coinmarketcap.com/v1/ticker")
	if err != nil {
		panic(err)
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var ar ApiResponse
		err = json.Unmarshal(body, &ar)
		if err != nil {
			panic(err)
		}

		for _, coin := range ar {
			if includes(holdings, coin.Id) {
				rows = append(rows, []string{
					coin.Name,
					coin.Symbol,
					coin.PriceUsd,
					coin.PriceBtc,
					coin.PercentChange1h + "%",
					coin.PercentChange24h + "%",
				})
			}
		}
		return rows
	}
}

func buildTable(rows [][]string) {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	tb := ui.NewTable()
	tb.Rows = rows
	tb.FgColor = ui.ColorWhite
	tb.BgColor = ui.ColorDefault
	tb.Separator = true
	tb.Analysis()
	tb.SetSize()
	tb.BorderFg = ui.ColorCyan
	tb.Y = 50
	tb.X = 0
	tb.Height = 100

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0)),
		ui.NewRow(ui.NewCol(12, 0, tb)),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/5s", func(e ui.Event) {
		ui.Body.Align()
		ui.Render(ui.Body)
	})

	ui.Loop()
}

func main() {
	conf := config.LoadConfiguration("./configs.yaml")
	fmt.Println(conf)

	var tickers []string
	for _, coin := range conf.Coins {
		tickers = append(tickers, coin.Ticker)
	}
	fmt.Println(tickers)

	rows := getCoins(tickers)
	buildTable(rows)
}
