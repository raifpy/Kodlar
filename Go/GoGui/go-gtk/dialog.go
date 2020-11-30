package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.Connect("destroy", gtk.MainQuit)
	win.SetTitle("NewWindow")

	dialog := gtk.NewDialog()
	//dialog.Connect("destroy", gtk.MainQuit) // Bu exit ab
	dialog.SetTitle("Merhaba Dünya!")
	dialog.SetPosition(gtk.WIN_POS_CENTER)

	dialog.Show()

	win.Resize(600, 400)
	//win.Activate()
	win.Show()
	gtk.Main()
}
