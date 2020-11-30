package main

import (
	"fmt"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	cmd := exec.Command("date")
	stderr, err := cmd.StderrPipe()
	errHandler.HandlerExit(err)
	abc := make([]byte, 5)
	str, err := stderr.Read(abc)
	errHandler.HandlerExit(err)
	fmt.Println(string(str))
}

// Ne işe yaradı bende anlayamadım
// Üstüne çalışmadı bile :)
