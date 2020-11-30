package main

import (
	"fmt"
	"os"
)

//Handler Bilader Hata ayıklıyor işte :D
func Handler(err error) {
	if !(err == nil) {
		fmt.Println("Olamaz \033[31mhata\033[0m yakalandı :", err.Error())
		os.Exit(1)
	}
}

//main main fonksiyon
func main() {
	dosya, err := os.Open("dosya")
	Handler(err)
	stat, err := dosya.Stat()

	Handler(err)
	hamVeri := make([]byte, stat.Size())
	//hamVeri := make([]byte, 1)
	_, err = dosya.Read(hamVeri)
	Handler(err)
	fmt.Println(hamVeri)
	fmt.Println(string(hamVeri))

}
