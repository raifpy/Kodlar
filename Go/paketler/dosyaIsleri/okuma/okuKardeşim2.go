package main

import (
	"errHandler"
	"fmt"
	"io/ioutil"
)

func main() {
	icerik, err := ioutil.ReadFile("dosya")
	errHandler.HandlerExit(err)
	fmt.Println(string(icerik))
}

// os.Open() 'a kıyasla çok daha rahat
