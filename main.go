package main

import (
  "./config"

  ui "github.com/gizak/termui"
  "fmt"

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

func main() {
  c := config.LoadConfiguration("./configs.yaml")

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
