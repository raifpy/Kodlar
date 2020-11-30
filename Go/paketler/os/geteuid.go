package main

import (
	"fmt"
	"os"
)

func main() {
	uid := os.Geteuid()
	if uid != 0 {
		fmt.Println("Olamaz dostum root değilsin !")
	} else {
		fmt.Println("Merhaba root")
	}
}
