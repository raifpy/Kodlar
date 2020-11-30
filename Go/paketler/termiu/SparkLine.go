package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/raifpy/Go/errHandler"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func getAvarenge(list []float64) int {
	lenlist := len(list)
	var toplam float64
	for _, eleman := range list {
		toplam += eleman
	}
	return int(toplam) / lenlist
}

func main() {
	errHandler.HandlerExit(ui.Init())
	defer ui.Close()
	defer ui.Clear()
	sp1 := widgets.NewSparkline()

	sp1.LineColor = ui.ColorBlue
	sp1.TitleStyle.Fg = ui.ColorBlue //

	sp1.Data = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sp1.Title = fmt.Sprint("sp1 ", getAvarenge(sp1.Data))

	sp2 := widgets.NewSparkline()

	sp2.LineColor = ui.ColorCyan
	sp2.TitleStyle.Fg = ui.ColorCyan //

	sp2.Data = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sp2.Title = fmt.Sprint("sp2 ", getAvarenge(sp2.Data))

	spGroup := widgets.NewSparklineGroup(sp1, sp2)
	spGroup.Title = "NewSparkLineGroup"
	spGroup.TitleStyle.Fg = ui.ColorGreen
	spGroup.BorderStyle.Fg = ui.ColorYellow
	en, boy := ui.TerminalDimensions()
	spGroup.SetRect(0, 0, en, boy)

	ui.Render(spGroup)

	for ch := range ui.PollEvents() {
		switch ch.ID {
		case "q", "<C-c>":
			return
		default:
			for index, eleman := range sp1.Data {
				sp1.Data[index] = eleman + float64(randomInt(-100, 100))
			}
			sp1.Title = fmt.Sprint("sp1 ", getAvarenge(sp1.Data))
			if getAvarenge(sp1.Data) < 0 {
				sp1.LineColor = ui.ColorRed
			} else if sp1.LineColor == ui.ColorRed {
				sp1.LineColor = ui.ColorBlue
			}

			//

			for index, eleman := range sp2.Data {
				sp2.Data[index] = eleman + float64(randomInt(-100, 100))
			}
			sp2.Title = fmt.Sprint("sp2 ", getAvarenge(sp2.Data))
			if getAvarenge(sp2.Data) < 0 {
				sp2.LineColor = ui.ColorRed
			} else if sp2.LineColor == ui.ColorRed {
				sp2.LineColor = ui.ColorCyan
			}
			ui.Render(spGroup)
		}
	}

}
