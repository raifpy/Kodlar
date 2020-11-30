package main

import (
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	errHandler.HandlerExit(os.Remove("olmayan_dosya"))
	errHandler.HandlerExit(os.RemoveAll("olmayan_path")) // çok vahşi bu :)
}
