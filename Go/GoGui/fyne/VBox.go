package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("Şef")
	win.Resize(fyne.NewSize(400, 200))
	win.CenterOnScreen()

	//label := widget.NewLabel("merhaba")
	label := widget.NewLabel("Şef")
	buton := widget.NewButtonWithIcon("Buton", theme.CheckButtonCheckedIcon(), func() {
		//tema := *fyne.Theme{}
		//uyg.Settings().SetTheme(theme.DarkTheme())
		//uyg2 := app.New()
		uyg.Settings().SetTheme(theme.LightTheme())

		win2 := uyg.NewWindow("o ye bum")
		win2.Resize(fyne.NewSize(400, 200))
		win2.Show()
	})
	radio := widget.NewRadio([]string{"linux", "windows", "mac"}, func(string) {})
	radio.SetSelected("linux")

	vbox := widget.NewVBox(label, buton, radio)

	win.SetContent(vbox)

	win.ShowAndRun()
}
