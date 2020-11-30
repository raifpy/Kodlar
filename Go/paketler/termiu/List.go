package main

import (
	"fmt"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/raifpy/Go/errHandler"
)

func makeBeauty(list []string) []string {
	newList := make([]string, 0)
	for _, eleman := range list {
		if eqineleman := strings.Index(eleman, "="); eqineleman >= 1 {
			tmp := eleman[:eqineleman]
			edited := strings.Replace(eleman, tmp, fmt.Sprintf("[%s](fg:green)", tmp), 1)
			newList = append(newList, edited)

		} else {
			newList = append(newList, eleman)
		}

	}
	return newList
}

func main() {

	errHandler.HandlerExit(ui.Init())
	defer func() {
		ui.Clear()
		ui.Close()
	}()

	liste := widgets.NewList()
	liste.Title = "NewList"
	//liste.Rows = os.Environ() // []string{"veri","veri"}

	liste.Rows = makeBeauty(os.Environ()) // []string

	liste.WrapText = false                      // ne işe yaradığı hakkında hiç bir fikrim yok
	liste.TextStyle = ui.NewStyle(ui.ColorCyan) // tüm elemanlar cyan olmuş oluyor
	liste.TitleStyle = ui.NewStyle(ui.ColorRed) // title (NewList) kırmızı oldu

	en, boy := ui.TerminalDimensions()
	liste.SetRect(0, 0, en, boy)
	//liste.SetRect(0, 0, 20, 5)

	ui.Render(liste)

	for key := range ui.PollEvents() {
		switch key.ID {
		case "q", "<C-c>":
			return

		case "<Down>":
			//fmt.Println("Down")
			liste.ScrollDown()
			ui.Render(liste) // Bu eleman çok önemli

		case "<Up>":
			liste.ScrollUp()
			ui.Render(liste) // Bu eleman çok önemli

		case "<Home>":
			liste.ScrollTop()
			ui.Render(liste)
		case "<End>":
			liste.ScrollBottom()
			ui.Render(liste)

		case "<C-d>": // bunlar gereksiz gibi ama kalsınlar , belki lazım olur
			liste.ScrollHalfPageDown()
		case "<C-u>":
			liste.ScrollHalfPageUp()
		case "<C-f>":
			liste.ScrollPageDown()
		case "<C-b>":
			liste.ScrollPageUp()

		}

	}

}
