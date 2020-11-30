package main

import "fmt"

func main() {
	var ilk = []string{"merhaba", "dünya"}
	var son = make([]string, len(ilk)) // ilk'in elemanları kadar make() ile liste oluşturdum

	fmt.Println(copy(son, ilk)) //
	fmt.Println(ilk)            //
	fmt.Println(son)            //
}

// listeleri kopyalamak için || dict'de muıhtemlnl
