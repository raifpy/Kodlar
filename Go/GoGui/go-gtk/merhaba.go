package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	win.SetTitle("Merhaba GTK")
	win.Resize(300, 100)
	win.ShowAll()
	gtk.Main()
}
