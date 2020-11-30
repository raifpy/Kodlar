package main

import "fmt"

func main() {
	// Değiken değer verilmeden atanırsa ;

	var STR string
	//var INT int
	//var BOOL bool
	//var FLOAT float32

	//fmt.Println(string,int,bool,float32) // Zero Value halindeyken hata verecek
	STR = "MerhabaDünya"
	fmt.Println(STR)
	//fmt.Println(STR="MerhabaDünya2") // Değişken atarken return edemiyoruz

	//fmt.Println(abc:="Merhabalar Efenim") // Aynı şey

	var x, y int = 100, 400 // Multi atama
	x, y = 10, 30

	fmt.Println(x, y)

	var a, b = 100, 400 // Multi atama 2 :D
	fmt.Println(a, b)

	abc, abcd := "Merhaba Dünya", 10 // var yapısı da aynı görevi görebilir
	fmt.Println(abc, abcd)
	//fmt.Scan()

}
