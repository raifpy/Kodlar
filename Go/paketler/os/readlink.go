package main

import (
	"log"
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	link, err := os.Readlink("testFile")
	errHandler.HandlerExit(err)
	log.Println(link)

	// Ne işe yaradığını bende eanlayamadım

}
