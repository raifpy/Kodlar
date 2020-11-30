package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	ui.Init()
	defer ui.Close() // Bu olmazsa bozuluyor :D
	//eleman := widgets.NewParagraph()
	eleman := widgets.NewParagraph()
	eleman.Title = "Merhaba Dünya !"
	eleman.Text = "\nMakale ve devamı"

	eleman.SetRect(0, 0, 30, 10)

	ui.Render(eleman)

	//time.Sleep(time.Second * 5)
	fmt.Scanln()

}
