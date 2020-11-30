package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	win := uyg.NewWindow("@raifpy - Buton.go")

	var sayi = 1
	label := widget.NewLabel("Merhaba Ben Orjinal Label")

	buton := widget.NewButton("Ben Buton", func() {
		// butona basıldığında aktif olacak eleman
		label.SetText("Merhaba Ben Butonsal Label {" + strconv.Itoa(sayi) + "}")
		sayi++

	})

	vbox := widget.NewVBox(label, buton)
	win.SetContent(vbox)

	win.CenterOnScreen()
	win.SetOnClosed(func() {
		fmt.Println("Peki , git .")
	})
	win.ShowAndRun()
}
