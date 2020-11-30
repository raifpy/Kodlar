package main

import (
	"errHandler"
	"fmt"
	"os"
)

func main() {
	var dosya = "dosya"
	var yazilacak = []string{"Merhaba Dünya !", "Merhaba Github", "Birileri Go için PDF yazsa çok hoş olmaz mı"}
	file, err := os.Create(dosya) // Create ile açmaz isek hata verecek :')

	errHandler.HandlerExit(err) // *Bana özel kütüphane ~> BetikSonu.org | Go Kütüphaneleri konusu

	for _, eleman := range yazilacak {
		yazilanInt, err := file.WriteString(eleman)
		if !(errHandler.HandlerBool(err)) { // Heyt
			fmt.Println(eleman, "\t", yazilanInt, "Kadar yazıldı !")
		}
	}
	file.Close()
	fmt.Println(dosya, "\033[33miçeriği okunuyor ;\033[0m")
	file, err = os.Open(dosya)
	errHandler.HandlerExit(err)
	var byteInt int
	{
		info, err := file.Stat()
		errHandler.HandlerExit(err)
		byteInt = int(info.Size())
	} // info değişkeni {} içerisinde kalmış oldu .

	icerik := make([]byte, byteInt)
	_, err = file.Read(icerik)
	errHandler.HandlerExit(err)
	fmt.Println(string(icerik)) // Böylece os.WriteString() ile yazarken sonuna \n koymayı unutmuyoruz

}

// Bilmiyorum sen her kimsin ama bu betiği okuyorsan senden olacak demektir
// Go hala yeni bir oluşum . Bu oluşum henüz küçük iken içerisinde olmak çok ama çok önemli
// Bilmiyorum bunu okuyan kişi kaç yaşında , ne düşünüyorum ama lise3 öğrencisinden bir tavsiye , bu yazılım dilinin geleceği var
// mobil , web , gömülü sistemler , desktop
// Tabiki mobilde hala stabil değil . SDK 'yı ayarlayamadığım için derleme yapamadım bile
// Ginede mobil için dil arıyorsanız kesinlikle Flutter(Dart) öğrenin ..

// @raifpy - Sal Tem 14 00:48:07 +03 2020 (bash - date)
