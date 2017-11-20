package main

import "github.com/gizak/termui"

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	rows := [][]string{
		[]string{
      "Name",
      "Symbol",
      "Price (USD)",
      "Price (BTC)",
      "% Change (1h)",
      "% Change (24h)"
    },
	}

	table1 := termui.NewTable()
	table1.Rows = rows
	table1.FgColor = termui.ColorWhite
	table1.BgColor = termui.ColorDefault
	table1.Y = 0
	table1.X = 0
	table1.Width = 62
	table1.Height = 7

	termui.Render(table1)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/timer/1m", func(e ui.Event) {
    // Refresh drawing
	})

	termui.Loop()
}
