package main

import (
	"fmt"
	"os"
)

func main() {

	icerik, err := os.OpenFile("dosya", os.O_WRONLY|os.O_APPEND, os.ModePerm) // Evet ilginç ..
	if !(err == nil) {
		fmt.Println("Olamaz , Hata : ", err.Error())
		return
	}
	yazilan, err := icerik.WriteString("Mehaba Go")
	fmt.Println(yazilan, err)
	icerik.Close()
}
