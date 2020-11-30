package main

import (
	"net/http"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("netTest")

	label := widget.NewLabel("Testing ...")
	go func() {
		_, err := http.Get("https://google.com")
		if err != nil {
			label.SetText(err.Error())
		} else {
			label.SetText("Connected :)")
		}
	}()

	win.SetContent(label)
	win.ShowAndRun()
}
