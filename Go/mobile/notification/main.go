package main

import (
	"strings"

	"fyne.io/fyne/theme"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	uyg.Settings().SetTheme(theme.DarkTheme())
	win := uyg.NewWindow("codeksiyon - send notification")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Codeksiyon Mobile")
	entry.SetText("Hi Codeksiyon!")

	//buton := widget.NewButton("Send", func() {
	buton := widget.NewButtonWithIcon("Send", theme.NavigateNextIcon(), func() {
		if strings.Trim(entry.Text, " ") == "" {
			uyg.SendNotification(fyne.NewNotification("Codeksiyon", "entry is empty :)"))
		} else {
			uyg.SendNotification(fyne.NewNotification("Codeksiyon", entry.Text))
		}
	})

	group := widget.NewGroup("Codeksiyon Notification", entry, buton)
	win.SetContent(group)
	win.ShowAndRun()
}
