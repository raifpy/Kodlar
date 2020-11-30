package main

import (
	"errHandler"
	"fmt"
	"os"
)

func main() {
	dosya, err := os.Create("dosya")
	errHandler.HandlerExit(err)

	//dosya.Seek()
	_, err = fmt.Fprintln(dosya, "Merhaba Dünya")
	errHandler.HandlerExit(err)

	// Py : open("dosya","w").write("Merhaba Dünya")
}
