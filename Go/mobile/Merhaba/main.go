package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("codeksiyon fyne mobile")

	win.SetContent(widget.NewLabel("Hello World\n\n@codeksiyon"))
	win.ShowAndRun()
}
