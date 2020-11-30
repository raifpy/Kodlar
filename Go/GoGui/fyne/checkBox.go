package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("checkBox")
	win.CenterOnScreen()

	label := widget.NewLabel("       ")
	var sayi = 0
	check := widget.NewCheck("Bas Bakim", func(doru bool) {
		if doru {
			label.SetText("Onayladın 0_0    " + strconv.Itoa(sayi))
			sayi++
		}
	})

	buton := widget.NewButton("Bas bakim o_o", func() {
		if check.Checked {
			win2 := uyg.NewWindow("Bastın")
			label := widget.NewLabel("Onayladığınız için teşekürler gibisinden cart curt")
			win2.SetContent(label)
			win2.CenterOnScreen()
			win2.Show()
		}
	})

	grup := widget.NewGroup("NewCheck", label, check, buton)
	win.SetContent(grup)
	win.ShowAndRun()
}
