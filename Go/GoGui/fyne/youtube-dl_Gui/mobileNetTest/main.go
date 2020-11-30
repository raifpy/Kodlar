package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	uyg.Settings().SetTheme(theme.DarkTheme())
	uyg.SendNotification(fyne.NewNotification("şef", "merhaba codeksiyon"))

	win := uyg.NewWindow("codeksiyon net test")
	label := widget.NewLabel("şef , yükleniyor :)")

	go func() {
		time.Sleep(time.Second * 5)
		istek, err := http.Get("http://ip-api.com/json")
		if err != nil {
			label.SetText("Hüüüü " + err.Error())
		}
		html, _ := ioutil.ReadAll(istek.Body)
		label.SetText(string(html))
	}()

	scrollBox := widget.NewScrollContainer(label)
	scrollBox.SetMinSize(fyne.NewSize(400, 400))

	win.SetContent(scrollBox)

	win.ShowAndRun()
}
