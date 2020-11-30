//barchart ama çoklu

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

	stacked := widgets.NewStackedBarChart()
	stacked.Title = "NewStackedBarChart"
	stacked.TitleStyle.Fg = ui.ColorBlue

	stacked.Labels = []string{"Bir", "İki", "Üç"}
	stacked.BarWidth = 6 // bar 'ların en lerini ayarlıyor

	stacked.Data = make([][]float64, 3)
	stacked.Data[0] = []float64{1, 2, 3, 4, 5, 6, 7} // 1 ve 2 bar'da gözükmüyor :D
	stacked.Data[1] = []float64{1, 2, 3, 4, 5, 6}
	stacked.Data[2] = []float64{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	stacked.Border = false // borde [] olmayacak
	//stacked.BarColors[0] = ui.ColorClear // neye göre ayarlıyor bilmiyorum

	en, boy := ui.TerminalDimensions()
	stacked.SetRect(0, 0, en, boy)

	ui.Render(stacked)

	fmt.Scanln()

}
