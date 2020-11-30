package main

import "fmt"

func main() {
	var sayi = 10
	if sayi < 10 {
		fmt.Printf("Sayınız 10'dan küçük , \033[31m %d \033[0m\n", sayi) // Printf kullanmamızın sebebi %d ,s f, v gibi elemanların çalışmasını istememeiz
	} else if sayi > 10 {
		fmt.Printf("Sayınız 10'dan büyük , \033[31m %d \033[0m\n", sayi)
	} else {
		fmt.Printf("Sayınız 10 olsa gerek . Else bloğu çalışıyor ..\n") // Printf'de \n oto değil , biz ekliyoruz mühim değil
	}
}
