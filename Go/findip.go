package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	fmt.Println(ip())
}

func ip() string {
	addrs, err := net.Interfaces()
	if errHandler.HandlerBool(err) {
		return ""
	}
	//var AddrList = []string{}
	for _, interfacea := range addrs {
		adr, err := interfacea.Addrs()
		if errHandler.HandlerBool(err) {
			return ""
		}
		adrwithport := adr[0].String()
		slashint := strings.Index(adrwithport, "/")
		addrStr := adrwithport[0:slashint]

		if addrStr != "127.0.0.1" && addrStr != "localhost" {
			return addrStr
		}

	}

	return ""
}
