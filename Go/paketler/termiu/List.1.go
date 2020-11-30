package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/raifpy/Go/raiFile"

	"github.com/raifpy/Go/errHandler"

	"github.com/gizak/termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
)

var sayi = 30

func main() {
	errHandler.HandlerExit(ui.Init())
	defer func() {
		ui.Clear()
		ui.Close()
	}()

	list := widgets.NewList()
	list.Title = "NewList_4.1"
	list.TitleStyle = ui.NewStyle(ui.ColorMagenta)
	list.TextStyle = ui.NewStyle(ui.ColorMagenta)
	list.BorderStyle.Fg = ui.ColorBlue // Yukarıdaki elemanı da bu biçimde kullanabiliriz
	//list.Rows = os.Environ()
	file, err := raiFile.ReadFile("./termiu4.1.go")
	if errHandler.HandlerBool(err) {
		fmt.Println("./termiu4.1.go not exists !")
		os.Exit(1)
	}
	list.Rows = strings.Split(file, "\n")

	en, boy := ui.TerminalDimensions()
	list.SetRect(0, 0, en, boy)

	ui.Render(list)

	for ch := range ui.PollEvents() {
		switch ch.ID {
		case "q", "<C-c>":
			return
		case "<Up>":
			list.ScrollUp()

		case "<Down>":
			list.ScrollDown()

		case "<Left>":
			list.ScrollTop()

		case "<Right>":
			list.ScrollBottom()
		case "<Space>":
			newList(list, os.Environ())
		}
		if ch.Type == ui.KeyboardEvent {
			ui.Render(list)
			sayi--
		}

	}

}

func newList(list *widgets.List, nwlist []string) {
	list.Rows = nwlist[:sayi]
	ui.Render(list)
}
