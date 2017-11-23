package main

import (
	"./config"

	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
}

func includes(coins []string, coin string) bool {
	for _, n := range coins {
		if n == coin {
			return true
		}
	}
	return false
}

func getCoins(holdings []string) ([]ui.Attribute, [][]string) {
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

		colors := []ui.Attribute{ui.ColorWhite}

		for _, coin := range ar {
			if includes(holdings, coin.Id) {
				rows = append(rows, []string{
					coin.Name,
					coin.Symbol,
					coin.PriceUsd,
					coin.PriceBtc,
					coin.PercentChange1h + "%",
				})

				f, _ := strconv.ParseFloat(coin.PercentChange1h, 64)
				if f >= 0 {
					colors = append(colors, ui.ColorGreen)
				} else {
					colors = append(colors, ui.ColorRed)
				}
			}
		}
		return colors, rows
	}
}

func setTableDefaults(tb *ui.Table) {
	tb.BgColor = ui.ColorDefault
	tb.Separator = true
	tb.Analysis()
	tb.SetSize()
	tb.BorderFg = ui.ColorCyan
	tb.Y = 50
	tb.X = 0
	tb.Height = 100
}

func refresh(tb *ui.Table) {
	colors, rows := getCoins(tickers)
	tb.Rows = rows
	tb.FgColors = colors

	ui.Clear()
	ui.Body.Align()
	ui.Render(ui.Body)
}

var tickers []string

func main() {
	// Grab list of coins to display
	conf := config.LoadConfiguration("./configs.yaml")
	fmt.Println(conf)
	for _, coin := range conf.Coins {
		tickers = append(tickers, coin.Ticker)
	}

	// Render stuff
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	tb := ui.NewTable()

	setTableDefaults(tb)
	ui.Body.AddRows(ui.NewRow(ui.NewCol(12, 0, tb)))
	refresh(tb)

	// Press q to quit
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	// Press r to refresh
	ui.Handle("/sys/kbd/r", func(ui.Event) {
		refresh(tb)
	})

	// Endpoints only update every 5mins
	ui.Merge("/timer/1m", ui.NewTimerCh(time.Second*60))
	ui.Handle("/timer/1m", func(e ui.Event) {
		refresh(tb)
	})

	ui.Loop()
}
