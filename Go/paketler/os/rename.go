package main

import (
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	errHandler.Handler(os.Rename("testFile", "testFile2"))
	errHandler.Handler(os.Rename("testFile2", "testFile"))
}
