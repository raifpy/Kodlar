package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/raifpy/Go/errHandler"
)

func main() {

	errHandler.HandlerExit(ui.Init())
	defer ui.Close()

	eleman := widgets.NewBarChart()
	eleman.Data = []float64{5, 2, 3, 4, 7, 8, 4, 10}
	eleman.Labels = []string{"Bir", "iki", "üç", "dört", "beş", "altı", "yedi", "sekiz"}
	eleman.Title = "NewBarChart"
	//eleman.BarColors = []ui.Color{ui.ColorYellow, ui.ColorCyan} // renk girilmez ise varsayılan renkler kullanılacaktır
	//eleman.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)} // aynı iş , hiç gerek yok
	//eleman.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)} // hiç gerek yok :D

	en, boy := ui.TerminalDimensions()
	eleman.SetRect(0, 0, en, boy)

	ui.Render(eleman)
	var author = true
	for keys := range ui.PollEvents() { // aktif edildiği zaman ctl-c , ctrl-z gibi elemanlar çalışmayacak
		if keys.Type == ui.MouseEvent {
			if keys.ID == "<MouseWheelDown>" && author {
				fmt.Print("Poweredby \033[4;31mraifpy\033[0m") // Bozuyor :D
				author = false
			}
		} else if keys.Type == ui.KeyboardEvent {
			break
		}
	}
	//fmt.Scanln()
}
