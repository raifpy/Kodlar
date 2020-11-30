package main

import (
	"strings"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	w := uyg.NewWindow("Merhaba")
	w.SetContent(widget.NewLabel("Merhaba Dünya !"))
	w.CenterOnScreen()

	go func() {
		for {
			for index, eleman := range strings.Split("Merhaba Dünya !", "") {
				if index == 0 {
					w.SetTitle(eleman)
				} else {
					w.SetTitle(w.Title() + eleman)
				}
				time.Sleep(time.Second / 4)
			}
		}
	}()

	w.ShowAndRun()

}

// Bu kütüphane iyi lakin gtk yerine bunu neden kullanayım düşünmek istiyorum
// qt windows içinde derlenebiliyor ama qt qtbox adında bir eleman indiriyor .
// Bu eleman windows için de derlenebiliyor ama windows bilgisayarda derlenmesi gerek
