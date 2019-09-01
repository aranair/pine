package main

import (
	"github.com/aranair/pine/config"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

func getCoins(holdings []config.Coin) (map[int]ui.Style, [][]string) {
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

		colors := map[int]ui.Style{
			0: ui.NewStyle(ui.ColorWhite),
		}

		for i, coin := range ar {
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
						colors[i+1] = ui.NewStyle(ui.ColorGreen)
					} else {
						colors[i+1] = ui.NewStyle(ui.ColorRed)
					}
				} else {
					colors[i+1] = ui.NewStyle(ui.ColorWhite)
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

func setTableDefaults(tb *widgets.Table) {
	tb.RowSeparator = true
}

func refresh(tb *widgets.Table) {
	colors, rows := getCoins(conf.Coins)
	tb.Rows = rows

	tb.TextStyle = ui.NewStyle(ui.ColorWhite)
	tb.RowSeparator = true
	tb.BorderStyle = ui.NewStyle(ui.ColorWhite)
	tb.SetRect(0, 0, 150, 50)
	tb.FillRow = true
	tb.RowStyles = colors

	ui.Clear()
	ui.Render(tb)
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

	tb := widgets.NewTable()
	setTableDefaults(tb)
	refresh(tb)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second * 5).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "r":
				refresh(tb)
			}
		case <-ticker:
			refresh(tb)
		}
	}
}
