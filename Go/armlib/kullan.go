package main

import (
	"fmt"
	"os"

	"github.com/raifpy/Go/errHandler"
	"github.com/rainycape/dl"
)

import "C"

func main() {
	file, err := dl.Open(os.Args[1], 1)
	errHandler.HandlerExit(err)

	var fonk func() *C.char
	file.Sym("sayHi", &fonk)
	str := C.GoString(fonk())
	fmt.Println(str)
}
