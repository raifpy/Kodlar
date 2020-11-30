package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)                        // nil 'de verebilirdik , ziyanı yok yine de
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) // popup olsun bu da
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	win.SetTitle("Merhaba Dünya !")
	win.SetTooltipText("Merhabna , @codeksiyon tarafından hazırlanan bir betiğe bakıyorsunuz")
	win.Connect("destroy", gtk.MainQuit) // Bu eleman gtk.Main() 'i destroy ediyor

	text := gtk.NewTextView()
	text.SetTooltipText("@codeksiyon")
	text.SetEditable(true)
	text.SetVisible(true)

	buffer := text.GetBuffer()
	//buffer.SetText("ilginç")

	win.Add(text)

	win.Resize(500, 300)

	//win.ShowAll()
	win.Show()
	fmt.Println("1")
	gtk.Main()
	fmt.Println("2")

}
