//gerçek zamanlı exec outputu vercek abum

package main

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	cmd := exec.Command("./zamanGecir")
	stdout, err := cmd.StdoutPipe()
	errHandler.HandlerExit(cmd.Start())
	fmt.Println("bir : ", err)

	bf := bufio.NewReader(stdout)

	for {
		line, _, err := bf.ReadLine()
		errHandler.HandlerExit(err)
		fmt.Println(string(line))

	}

}
