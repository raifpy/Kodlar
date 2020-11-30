package main

import (
	"fmt"
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	proc, err := os.FindProcess(2848) // pid verilir
	errHandler.HandlerExit(err)

	fmt.Println("Proc  : ", *proc)
	fmt.Println("PID   : ", proc.Pid)
	//errHandler.Handler(proc.Signal("SIGKILL")) // Bilemiedim os.Signal 'in ne olabileceğini
	errHandler.Handler(proc.Kill()) // Spotify pid'ine kill attı

	// kill
	// pid
	// release
	// signal
	// wait
}
