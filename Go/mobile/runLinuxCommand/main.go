package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	uyg := app.New()
	//uyg.Settings().SetTheme(theme.DarkTheme())
	win := uyg.NewWindow("Run Linux Commands")

	text := widget.NewTextGridFromString("Codeksiyon - Mobile")
	textGroup := widget.NewScrollContainer(text)
	textGroup.SetMinSize(fyne.NewSize(300, 300))

	entry := widget.NewEntry()

	buton := widget.NewButtonWithIcon("Run", theme.ConfirmIcon(), func() {
		if strings.Trim(entry.Text, " ") == "" {
			text.SetText("Write some Linux command !")
			return
		}
		go func() {
			textSplit := strings.Split(entry.Text, " ")

			var out, errOut bytes.Buffer

			cmd := exec.Command(textSplit[0], textSplit[1:]...)
			cmd.Stdout = &out
			cmd.Stderr = &errOut
			cmd.Start()

			done := make(chan error)
			go func() {
				done <- cmd.Wait()
			}()

			timer := time.After(5 * time.Second)

			select {
			case <-timer:
				cmd.Process.Kill()
				text.SetText("Timeout ! { 5sec }")
			case err := <-done:
				str := out.String()
				if err != nil && str == "" {
					text.SetText(errOut.String() + "\n" + err.Error())
				} else {
					text.SetText(str)
				}
			}
		}()
	})
	themeButon := widget.NewButton("Theme", func() {
		if uyg.Settings().Theme() == theme.DarkTheme() {
			uyg.Settings().SetTheme(theme.LightTheme())
		} else {
			uyg.Settings().SetTheme(theme.DarkTheme())
		}
	})

	used := widget.NewLabelWithStyle("Calculating", fyne.TextAlignLeading, fyne.TextStyle{false, true, true})
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			alloc := m.Alloc
			free := m.Frees
			used.SetText(fmt.Sprintf("Alloc : %d mb\nFrees : %d mb", alloc/1024/1024, free/1024/1024))

			time.Sleep(time.Second)
			//runtime.GC()
		}
	}()

	win.SetContent(widget.NewVBox(textGroup, entry, buton, themeButon, used))
	win.ShowAndRun()

}
