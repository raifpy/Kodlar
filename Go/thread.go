// @BetikSonu | Go Thread
package main

import (
	"fmt"
	"time"
)

func test(s string) {
	for i := range s {
		fmt.Println(s[:len(s)-i])
		time.Sleep(time.Second)
	}
}
func main() {
	go test("Merhaba Dünya")
	test("BetikSonu")
}
