package main

import (
	"strings"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("inputVre")

	label := widget.NewLabel("Merhaba , aşağıya birkaç birşeyler yazdıktan sonra\n\"bas\" adlı butona basabilirsiniz")
	input := widget.NewEntry()

	//input.TypedKey(&fyne.KeyEvent{fyne.KeyEnter}) // Bu olacak gibi ama dur bakalım :D

	buton := widget.NewButton("Bas", func() {
		text := strings.Trim(input.Text, " ")
		if text == "" {
			label.SetText("Boş değer verme kardeşim :)")
		} else {
			label.SetText("Değer : " + input.Text)
		}
	})

	vbox := widget.NewVBox(label, input, buton)
	win.SetContent(vbox)

	win.ShowAndRun()

}
