package main

import "fmt"

func main() {
	liste := [2]string{"Merhaba", "Dünya"}
	for i, eleman , a := range liste {
		fmt.Println(i, ". eleman = ", eleman)
	}

}
