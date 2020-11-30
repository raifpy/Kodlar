package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	port := ":" + os.Args[1]
	tcp, err := net.Listen("tcp4", port)
	errHandler.HandlerExit(err)
	fmt.Println("Döngüye giriliyor")
	for {
		c, err := tcp.Accept()
		fmt.Println(c)
		if errHandler.HandlerBool(err) {
			continue
		}
		go func() {
			for {
				veri, err := bufio.NewReader(c).ReadString('\n')
				if errHandler.HandlerBool(err) {
					continue
				}

				fmt.Println("VERİ  : ", string(veri))
				c.Write([]byte("codeksiyon"))
			}
		}()

	}
}
