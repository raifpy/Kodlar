package main

import (
	"fmt"
	"time"
)

func main() {
	abc := time.NewTicker(time.Second)
	iint := 0
	for range abc.C {
		fmt.Println("Merhaba Dünya", iint)
		iint++
	}
}
