package main

import (
	"fmt"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	siemdi := exec.Command("notify-send", "Dünya", "Merhaba Dünya")
	errHandler.Handler(siemdi.Run())

	echo, err := exec.Command("echo", "merhaba", "Dünya").Output()
	errHandler.HandlerExit(err)

	fmt.Println(string(echo))

}
