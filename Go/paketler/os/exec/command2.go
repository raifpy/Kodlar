package main

import (
	"fmt"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	fmt.Println("Başladı")
	cmd := exec.Command("sleep", "10")
	errHandler.HandlerExit(cmd.Start())
	fmt.Println("Yarım")
	errHandler.HandlerExit(cmd.Wait()) // Commadn() ın bitmesini bekliyor
	fmt.Println("Bitti")
}

// .Run() 'da bekletiyor
// . Start() bekletmiyor , direk devam
