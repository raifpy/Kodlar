//gerçek zamanlı exec outputu vercek abum

package main

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	cmd := exec.Command("youtube-dl", "https://www.youtube.com/watch?v=Dcn4fU9ZI2w&list=PLz-8cAsN0KAH6KvRKlyiNlT49-zk1nNkA&index=6")
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()

	errHandler.HandlerExit(cmd.Start())
	fmt.Println("bir : ", err)

	bf := bufio.NewReader(stdout)
	bferr := bufio.NewReader(stderr)
	for {
		line, _, err := bf.ReadLine()
		if errHandler.HandlerBool(err) {
			break
		}
		fmt.Println(string(line))
	}

	for {
		line, _, err := bferr.ReadLine()
		if errHandler.HandlerBool(err) {
			break
		}
		fmt.Println("ÇIKTI : ", string(line))
	}

}
