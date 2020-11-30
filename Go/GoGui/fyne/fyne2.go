package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uygulama := app.New()
	window := uygulama.NewWindow("window")
	window.SetTitle("Merhaba Fyne !")

	label := widget.NewLabel("Merhaba Dünya ...")
	//label.Hide()
	//label.Refresh() // işe yarar

	label.Resize(fyne.Size{200, 50}) // Değişmedi ama bilmiyorum :D
	label.Refresh()

	window.SetContent(label)
	window.CenterOnScreen()

	go func() {
		for {
			for index, eleman := range strings.Split(strings.Join(os.Args, " "), "") {
				if index == 0 {
					// window.FixedSize() // etkisi olmadı
					label.SetText(eleman)
					window.SetTitle("@raifpy - ")
				} else {
					label.SetText(label.Text + eleman)
					window.SetTitle(window.Title() + eleman)
				}
				time.Sleep(time.Second / 4)
			}
		}
	}()
	window.SetOnClosed(func() {
		fmt.Println("Olamaz program kapatıldı")
	})
	window.ShowAndRun()
}
