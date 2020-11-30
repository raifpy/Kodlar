package main

import (
	"fmt"
	"time"
)

func main() {
	//time.Sleep(1000000000 * 5) // 5 saniye :D
	var milisecond = time.Microsecond
	fmt.Println(milisecond) // 1µs
	time.Sleep(time.Second) // 1 saniye
	fmt.Println("\t\t----------------------")
	time.Sleep(time.Second * 4) // 4 saniye
}
