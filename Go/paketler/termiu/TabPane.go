package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	ui.Init()
	defer ui.Close()
	defer ui.Clear()

	tab := widgets.NewTabPane("Bir", "İki", "Üç", "Dört", "Beş")
	tab.SetRect(0, 0, 50, 5)

	loading := widgets.NewGauge()
	loading.Percent = 40
	loading.SetRect(0, 30, 20, 20)
	loading.Border = false

	shark := widgets.NewSparkline()
	shark.Data = []float64{1, 2, 3, 5, 6, 7, 2}
	shark.LineColor = ui.ColorYellow

	sharkGroup := widgets.NewSparklineGroup(shark)
	sharkGroup.SetRect(0, 30, 20, 10)

	renderTab := func() {
		switch tab.ActiveTabIndex {
		case 0:
			ui.Render(loading)
		case 1:
			ui.Render(sharkGroup)
			//default:
			//ui.Clear()

		}
	}

	ui.Render(tab)

	for ch := range ui.PollEvents() {
		switch ch.ID {
		case "q", "<C-c>":
			return
		case "<Left>":
			tab.FocusLeft()
			ui.Clear()
			ui.Render(tab)
			renderTab()
		case "<Right>":
			tab.FocusRight()
			ui.Render(tab)
			renderTab()
		}
	}
}
