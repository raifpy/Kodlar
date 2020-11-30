package main

import (
	"fmt"
	"net"

	"github.com/raifpy/Go/errHandler"
)

//const ip = "240.24.205.35.bc.googleusercontent.com"
//const ip = "35.205.24.240"

func main() {
	var ip string
	fmt.Print("Site : ")
	fmt.Scanln(&ip)
	addr, err := net.LookupIP(ip)
	errHandler.HandlerExit(err)
	fmt.Println(addr[0])
	fmt.Println(net.LookupAddr(addr[0].String()))
	fmt.Println(net.LookupCNAME(addr[0].String()))
}
