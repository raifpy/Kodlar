// Bar (Loading)

package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/raifpy/Go/errHandler"
)

func main() {
	errHandler.HandlerExit(ui.Init())
	defer ui.Close()
	defer ui.Clear()

	bar := widgets.NewGauge()
	bar.Title = "NewGauge"
	bar.BorderStyle.Fg = ui.ColorRed
	bar.LabelStyle.Fg = ui.ColorBlue

	bar.Percent = 1 // yüzde 1 seviyesinde bekleyen bir bar oluştu :D
	bar.SetRect(0, 0, 140, 5)
	bar.BarColor = ui.ColorBlue

	ui.Render(bar)

	for ch := range ui.PollEvents() {
		if ch.Type == ui.KeyboardEvent {
			if ch.ID == "q" {
				return
			}
			bar.Percent++
			ui.Render(bar)

		}
	}
}
