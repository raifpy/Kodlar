package main

// Yapamadım :D

import (
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	os.Setenv("FYNE_THEME", "dark")
	uyg := app.New()
	win := uyg.NewWindow("Theme değişilir")

	degis := func() {
		if os.Getenv("FYNE_THEME") == "dark" {
			os.Setenv("FYNE_THEME", "light")
			win.Close()
			main()
		} else {
			win.SetTitle("dark temaya geçildi")
			os.Setenv("FYNE_THEME", "dark")
			win.Close()
			main()
		}
	}
	buton := widget.NewButton("Değiş", degis)

	grup := widget.NewGroup("       [dark]      Tema     [light]        ", buton)

	win.SetContent(grup)
	win.SetOnClosed(main) // :D
	win.ShowAndRun()

	//fmt.Println("hm")
}
