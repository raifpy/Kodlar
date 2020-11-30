package main

import (
	"errHandler"
	"fmt"
	"os"
)

func main() {
	dosya, err := os.Open("dosya") // os.Create() dosyayı sıfırlıyor
	errHandler.HandlerExit(err)    // Hata olup olmadığını denetlediğim kütüphanem . URL ?

	info, err := dosya.Stat()
	errHandler.HandlerExit(err)
	str := make([]byte, info.Size())

	_, err = dosya.Read(str)
	errHandler.HandlerExit(err)
	fmt.Println(str)
	fmt.Println(string(str))

}
