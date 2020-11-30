package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	veri, _ := ioutil.ReadDir(".")
	for index, eleman := range veri {
		fmt.Print(index, "\t", eleman.Name())
		fmt.Println(eleman.Size(), " ", eleman.ModTime().String(), " ", eleman.Mode().String())
	}
}
