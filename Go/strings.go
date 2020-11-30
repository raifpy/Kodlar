package main

import "fmt"

func main() {
	var str = "Merhaba Dünya !"
	fmt.Println(str[1])  // 1
	fmt.Println(str[5:]) // ba Dünya !
	fmt.Println(str[:5]) // Merha

	fmt.Println("Güzel")

	var str2 = str[1]
	fmt.Println(str2)         // Sadece 1 elemanı alırsak , o elemanın utf8'deki sayi hali geliyor // 101
	fmt.Println(string(str2)) // string() ile 'de string yapıyoruz // e
}
