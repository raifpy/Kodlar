package main

import (
	"errHandler"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(file *os.File) string {
	// ioutil ile de yapabiliriz ama []byte ile yapacağım
	info, err := file.Stat()
	errHandler.HandlerExit(err)
	icerik := make([]byte, info.Size())
	_, err = file.Read(icerik)
	errHandler.HandlerExit(err)
	return string(icerik)

} // Hmmm , bu eleman aynen şu hatayı döndürdü : ~/M/K/G/p/d/yazma ❯❯❯ go run yazKardeşim3.go 2020/07/14 01:06:03 Error :  EOF Exiting (1)

func ReadFile2(fileName string) string {
	// ioutil ile açacağız :D
	icerik, err := ioutil.ReadFile(fileName)
	errHandler.HandlerExit(err)
	return string(icerik)

}

func main() {
	var dosya = "dosya"
	yazilacak := []string{"Merhaba Dünya", "Merhaba BetikSonu", "Merhaba Github", "Richard'a duyduğum saygının aynısını sana duyuyorum Fırat hocam. Eline sağlık .. \t Python3 istihza", "Merhaba C v2  `Go`"}

	file, err := os.Create(dosya) // Open() hata veriyor

	errHandler.HandlerExit(err) // HandlerAndExit anlamında :')
	// yazKardeşim2 'den öğrendiğimiz ile aynı , sonuna \n koymaz isek yan yana yazacak

	for _, eleman := range yazilacak {
		fmt.Fprintln(file, eleman) // println 'daki ln line anlamında , kendisi \n'leri otomatik verecektir
	}
	// Printf 'de format anlamında olsa gerek

	fmt.Println("Yazıldı\fİçerik okunuyor : \n")
	fmt.Println(ReadFile2(dosya))

}
