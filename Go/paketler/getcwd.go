package main

import (
	"errHandler"
	"fmt"
	"os"
)

func main() {
	konum, err := os.Getwd()
	errHandler.Handler(err)
	fmt.Println("ur loc : ", konum)

}
