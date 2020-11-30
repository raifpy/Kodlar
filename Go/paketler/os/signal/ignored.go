package main

import (
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	signal.Ignore(syscall.SIGINT) // ctrl - c
	fmt.Println("SIGINT  : ", signal.Ignored(syscall.SIGINT))
	fmt.Println("SIGKILL : ", signal.Ignored(syscall.SIGKILL))

}

// bool
