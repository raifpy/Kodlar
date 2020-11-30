package main

import (
	"fmt"
	"os"
)

// fmt ile yazacağız os/fmt

//Handler ...
func Handler(err error) {
	if !(err == nil) {
		fmt.Println("Olamaz , hata : ", err.Error())
		os.Exit(1)
	}
}

func main() {
	icerik, err := os.Create("dosya1_1") // eğer dosya yok ise dosya oluştuğu gibi kullanılabilir || os.Create() ile açtığımız dosyayı sıkıntısızca yazabiliyorken , os.Open() ile açılan dosyada bad file descriptor adlı bir hata alıyoruz
	Handler(err)

	var str = "Merhaba Dünya yazısı\n2 MErhaba"
	_, err = fmt.Fprintln(icerik, str)
	Handler(err)

	icerik.Close()

	// Çok ama çok daha rahat ..

}
