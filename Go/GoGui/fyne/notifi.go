package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	uyg := app.New()
	uyg.SendNotification(fyne.NewNotification("Notifi der ki", "osuruk kemiksiz kakadır yiğen"))

	uyg.Run()

}
