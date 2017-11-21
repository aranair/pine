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

func getApiCoins(tickers []string) []ApiCoin {
	res, err := http.Get("https://api.coinmarketcap.com/v1/ticker/bitcoin")
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

		return ar
	}
}

func startUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	rows := [][]string{
		headers,
	}

	table1 := ui.NewTable()
	table1.Rows = rows
	table1.FgColor = ui.ColorWhite
	table1.BgColor = ui.ColorDefault
	table1.Separator = false
	table1.Analysis()
	table1.SetSize()
	table1.Y = 0
	table1.X = 0
	table1.Height = 200

	// build
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, table1)),
	)

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1m", func(e ui.Event) {
		// Refresh drawing
	})

	ui.Loop()
}

func main() {
	c := config.LoadConfiguration("./configs.yaml")
	fmt.Println(c)

	coins := []string{"bitcoin"}
	getApiCoins(coins)

	// startUI()
}
