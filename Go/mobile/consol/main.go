package main

import (
	"fmt"
	"runtime"
	"time"

	"fyne.io/fyne/widget"

	"fyne.io/fyne/app"
)

func toMb(abc uint64) uint64 {
	return abc / 1024 / 1024
}

func runtimeMem() {
	var m runtime.MemStats
	for {

		//abc := "Hey"
		//abcd := &abc
		//fmt.Sprint(abc)

		runtime.ReadMemStats(&m)

		fmt.Printf("Alloc : %d\nTotalAlloc : %d\nSYS : %d\nLast GC : %d\n\n\n", toMb(m.Alloc), toMb(m.TotalAlloc), toMb(m.Sys), toMb(m.LastGC))
		time.Sleep(time.Second)

		//fmt.Scanln()
	}
}

func main() {
	go runtimeMem()
	uyg := app.New()

	win := uyg.NewWindow("Merhaba Dünya")

	label := widget.NewLabel("Merhaba Dünya !")
	entry := widget.NewEntry()
	buton := widget.NewButton("Onayla", func() {
		label.SetText(entry.Text)
	})

	grup := widget.NewGroup("Merhaba Grup", label, entry, buton)
	win.SetContent(grup)
	win.ShowAndRun()
}
