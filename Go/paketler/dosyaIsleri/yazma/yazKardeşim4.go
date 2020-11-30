// Bu elemanın kodlarını açıklayamam . Sağolsun t.me/golangtr 'deki bir abi yardımcı oldu

// Bu arkadaşın olayı ise os.Open() ile yazması. Bu da demek oluyor ki içerisindeki veriler silinmeden yazacak

package main

import (
	"errHandler"
	"fmt"
	"io/ioutil"
	"os"
)

/*func (file *os.File) DosyayıOkuYav() {

}*/ // Böyle olmuyormuş :D

func DosyayıOkuYav(file string) string {
	icerik, err := ioutil.ReadFile(file)
	errHandler.HandlerExit(err) // main'deki metine bak ..
	return string(icerik)
}

func main() {
	icerik, err := os.OpenFile("dosya", os.O_WRONLY|os.O_APPEND, os.ModePerm) // Evet ilginç ..
	errHandler.HandlerExit(err)                                               // Bana özel kütüphane ~> BetikSonu.org , Go Kütüphane oluşturma
	fmt.Println("Şuanki dosya içeriği \n", DosyayıOkuYav("dosya"))            // İşler ne kadar da kolaylaştı
	icerik.WriteString("\nKoca Yaşlı Şişko Dünya\n")                          // Ulan ya internetim olsa da dinlesem
	fmt.Println("\nYazıldıktan sonra : ", DosyayıOkuYav("dosya"))
	err = icerik.Close() // Belki hata verir bu da olamaz mı :)
	errHandler.HandlerExit(err)
	// Özet olarak open ile açtığımız için içeriği sıfırlanmamış oldu

}

// Hop author , date , program
// @raifpy - Sal Tem 14 01:29:03 +03 2020 ( bash - date )
