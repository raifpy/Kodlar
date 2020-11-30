package main

import "fmt"

func main() {
	// main fonksiyonu içerinde fonksiyon tanımalayacağız :))
	var sayi = 10

	getSayi := func() int {
		return sayi
	}
	setSayi := func(x int) {
		sayi = x
	}

	fmt.Println(getSayi())
	setSayi(1000)
	fmt.Println(getSayi())
}
