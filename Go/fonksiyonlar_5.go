package main

import "fmt"

func main() {
	// Gine o *argv olaynı yapacağım ama bu sefer iki veri türünü birden kullanacağız
	test()
}

func test(sayi ...int /*, str ...string*/) {
	//fmt.Println(len(sayi))
	//fmt.Println(len(str))
	abc := 0
	for _, deger := range sayi {
		abc += deger
	}
	fmt.Printf("Toplam %d elemanlı , toplam değeri %d çıktı ..\n", len(sayi), abc)
}

// Olmuyormuş :D
// Ama böylece gördük ki değer verilmeden de run olabiliyormuş fonksiyonlar

// Variadic deniyormuş bunlara
