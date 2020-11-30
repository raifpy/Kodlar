package main

import (
	"fmt"
	"os"
)

func main() {
	dosya, err := os.Open("dosya")
	if !(err == nil) {
		fmt.Println("Olamaz , hata : ", err.Error())
		return
	}

	bilgi, _ := dosya.Stat()
	byteArray := make([]byte, bilgi.Size())
	icerikByte, err := dosya.Read(byteArray)
	if !(err == nil) {
		fmt.Println("Olmaaz , hata (okuma esnasında ) : ", err.Error())
		return
	}
	icerik := string(byteArray)

	fmt.Println("Toplam byte ", icerikByte, "İçerik : \n", icerik)
}
