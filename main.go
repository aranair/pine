package main

import (
  "./config"

  ui "github.com/gizak/termui"
  "fmt"
  "net/http"
  "os"
  "io/ioutil"

  // "encoding/json"
  // "fsnotify"
)

var headers = []string{
  "Name",
  "Symbol",
  "Price in USD",
  "Price in BTC",
  "% Change (1h)",
  "% Change (24h)",
}

func makeApiCall(tickers []string) {
  res, err := http.Get("https://api.coinmarketcap.com/v1/ticker/")
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  } else {
    defer res.Body.Close()
    contents, err := ioutil.ReadAll(res.Body)
    if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
    }
    for i, v := range contents {
      fmt.Printf("%s, %s", i, v)
    }
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
  // table1.Width = 100
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
  makeApiCall(coins)

  // startUI()
}
