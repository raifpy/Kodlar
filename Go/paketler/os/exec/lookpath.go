package main

import (
	"fmt"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	path, err := exec.LookPath("go")
	errHandler.HandlerExit(err)

	fmt.Println("Çıktı :  ", path)
}

// $PATH 'in içerisinden uygulamaların konumu buluyor .

//which go | gibi davranıyor
