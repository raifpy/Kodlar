package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("Pencere")

	win.CenterOnScreen()
	win.Resize(fyne.NewSize(400, 200))

	bir := widget.NewEntry()
	bir.SetText("email")
	iki := widget.NewPasswordEntry()
	iki.SetText("şifre")

	uc := widget.NewProgressBar()
	uc.SetValue(0.30)
	//uc.Refresh()
	dortStr := "Merhaba Dünya !\nBu bir label .... "
	dort := widget.NewLabel(dortStr + " Orjinal")
	bes := widget.NewRadio([]string{"tamam", "peki", "tım"}, func(text string) {
		dort.SetText(dortStr + text)
		switch text {
		case "tamam":
			uc.SetValue(0.10)
		case "peki":
			uc.SetValue(0.40)
		case "tım":
			uc.SetValue(0.90)
		default:
			uc.SetValue(0)
		}
	})

	alti := widget.NewButton("Kontrol", func() {
		secilen := bes.Selected
		if secilen == "" {
			dort.SetText("Bir eleman seçmek durumundasın şef ")
		} else {
			dort.SetText("o ye " + secilen)
		}
	})

	gwithsroll := widget.NewGroupWithScroller("Merhaba Dünya", bir, iki, uc, dort, bes, alti)
	win.SetContent(gwithsroll)
	win.ShowAndRun()
}
