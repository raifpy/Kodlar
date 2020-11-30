package main

import (
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	errHandler.Handler(os.Mkdir("testDir", 1)) // string , perm //os.FıleExists hatası dönüyor

	errHandler.HandlerExit(os.MkdirAll("testDir/hey/hey2", 1))
}
