package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Println("PANIC !  : ", rec)
		}
	}()
	uyg := app.NewWithID("123")
	uyg.Settings().SetTheme(theme.DarkTheme())
	//uyg.SetIcon(theme.InfoIcon())

	win := uyg.NewWindow("gurl Gui")
	win.Resize(fyne.Size{800, 400})
	win.SetFixedSize(true)

	urlBox := widget.NewEntry()
	urlBox.SetPlaceHolder("https://ip-api.com/json")
	urlBox.SetText("https://ip-api.com/json")

	label := widget.NewTextGridFromString("a")
	label.SetText("Merhaba Dünya\nasdadasda\nasdasdasd")

	get := widget.NewButton("get", func() {
		if strings.Trim(urlBox.Text, " ") == "" {
			label.SetText("no url!")
		} else {
			req, err := http.Get(urlBox.Text)
			if err != nil {
				label.SetText(err.Error())
			} else {
				html, _ := ioutil.ReadAll(req.Body)
				label.SetText(string(html))
			}
		}
	})

	/*
		post := widget.NewButton("post", func() {
			if strings.Trim(urlBox.Text, " ") == "" {
				label.SetText("no url!")
			} else {
				req, err := http.PostForm(urlBox.Text, nil)
				if err != nil {
					label.SetText(err.Error())
				} else {
					html, _ := ioutil.ReadAll(req.Body)
					label.SetText(string(html))
				}
			}
		})
	*/
	//butonlar := widget.NewHBox(get, post)
	//butonlar.Resize(fyne.NewSize(800, 20))
	//grup := widget.NewVBox(urlBox, butonlar, label)
	grup := widget.NewVBox(urlBox, get, label)

	win.SetContent(grup)
	win.ShowAndRun()

}
