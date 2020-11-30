package main

import (
	"fmt"
	"os"
	"strings"
)

//Handler ...
func Handler(err error) {
	if !(err == nil) {
		fmt.Println("Olamaz \033[31mhata\033[0m ,", err.Error())
		fmt.Scanln()
		os.Exit(1)
	}
}

func main() {
	os.Create("dosya")
	icerik, err := os.Create("dosya") // Open'dan Create evirdik
	Handler(err)

	metin := "Merhaba Dünya"

	metinArray := strings.Split(metin, "")
	byteArray := make([]byte, len(metinArray))
	for sayi, metin := range metinArray {
		copy(byteArray[sayi:], metin)
		fmt.Println(byteArray)
	}

	fmt.Println(string(byteArray))

	//_, err = icerik.Write(byteArray)
	//Handler(err)
	_, err = icerik.WriteString("Merhaba Dünya")
	Handler(err)

	fmt.Scanln()
}

// Evet ilginç birşey oldu :D
// Bu da bozuldu Hmmm
// İkiside çalışmıyor :D

// En sonunda fark ettim ki os.Open() kullanarak dosyaya birşeyler yazamıyoruz . os.Create() yapısı ile yazabilriz ama ..
