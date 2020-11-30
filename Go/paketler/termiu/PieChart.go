// piechart
package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	ui.Init()

	defer ui.Close()
	defer ui.Clear()

	turta := widgets.NewPieChart()
	turta.Title = "NewPieChart"

	turta.Data = []float64{1, 2, 3, 4, 5}
	//turta.LabelFormatter(1, 2)

	en, boy := ui.TerminalDimensions()
	turta.SetRect(0, 0, en, boy)

	ui.Render(turta)

	fmt.Scanln()

	// Bu pie işinden pek birşey anlayamadım

}
