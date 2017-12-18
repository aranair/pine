package main

import (
	"./config"

	"encoding/json"
	"fmt"
	ui "github.com/gizak/termui"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	"Gain",
	"% Gain",
	"Overall",
}

func findHolding(holdings []config.Coin, coin string) (config.Coin, bool) {
	for _, h := range holdings {
		if h.Ticker == coin {
			return h, true
		}
	}
	return config.Coin{}, false
}

func fts(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func getCoins(holdings []config.Coin) ([]ui.Attribute, [][]string) {
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
			holding, included := findHolding(holdings, coin.Id)
			var gain float64
			var gs, gsp, gst string = "-", "-", "-"

			if included {
				priceUsdFloat, _ := strconv.ParseFloat(coin.PriceUsd, 64)
				if holding.Units > 0 {
					gain = priceUsdFloat - holding.Cost
					gs = fts(gain)
					gsp = fts(gain/holding.Cost*100) + "%"
					gst = fts(gain * holding.Units)

					if gain >= 0 {
						colors = append(colors, ui.ColorGreen)
					} else {
						colors = append(colors, ui.ColorRed)
					}
				} else {
					colors = append(colors, ui.ColorWhite)
				}

				rows = append(rows, []string{
					coin.Name,
					coin.Symbol,
					coin.PriceUsd,
					coin.PriceBtc,
					coin.PercentChange1h + "%",
					coin.PercentChange24h + "%",
					gs,
					gsp,
					gst,
				})

			}
		}
		return colors, rows
	}
}

func setTableDefaults(tb *ui.Table) {
	tb.BgColor = ui.ColorDefault
	tb.Separator = false
	tb.Analysis()
	tb.SetSize()
	tb.BorderFg = ui.ColorCyan
	tb.Y = 0
	tb.X = 0
	tb.Height = 30
}

func refresh(tb *ui.Table) {
	colors, rows := getCoins(conf.Coins)
	tb.Rows = rows
	tb.FgColors = colors

	ui.Clear()
	ui.Body.Align()
	ui.Render(ui.Body)
}

var conf config.Config

func main() {
	// Grab list of coins to display
	conf = config.LoadConfiguration("./configs.yaml")
	fmt.Println(conf)

	// Render stuff
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	tb := ui.NewTable()
	setTableDefaults(tb)
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, tb)),
	)
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
