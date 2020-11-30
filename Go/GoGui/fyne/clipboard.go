package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	defer func() {
		if abc := recover(); abc != nil {
			fmt.Println("Bilinmeyen bir sebepten ötürü clipboard adlı program çöktü !\n\n", abc)
		}
	}()
	uyg := app.New()
	win := uyg.NewWindow("take ur clipboard")

	label := widget.NewLabel(strings.Repeat(" ", 100))

	go func() {
		for {
			if clip := win.Clipboard(); clip != nil {
				label.SetText(strings.Replace(clip.Content(), "\t", "    ", -1))
			}
			time.Sleep(time.Second)
		}

	}()

	win.SetContent(label)

	win.Show()
	uyg.Run()
}
