package main

import (
	"net/url"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("HyperLink (link yauv)")

	url, err := url.Parse("https://t.me/raifpy")
	var hyperlinkStr string
	if err != nil {
		hyperlinkStr = "Olamaz bir hata aldım !"
	} else {
		hyperlinkStr = "Buradan bana ulaş 0_0"
	}
	hyperlink := widget.NewHyperlink(hyperlinkStr, url)

	win.SetContent(hyperlink)

	win.CenterOnScreen()

	win.ShowAndRun()

}
