package main

import "fmt"

func main() {
	abc := []string{"M", "E", "R", "H", "A", "B", "A"}
	for index, eleman := range abc {
		fmt.Println(index, eleman)
	}

	for _, eleman := range abc {
		fmt.Println(eleman) // _ hata veriyor :))
	}

	//for deneme_:= abc{ // range olmadan ne olacak bakalım
	//	fmt.Println(deneme_)
	//
	//} Olmuyormuş böylece anlamış olduk

	// Peki range'i fmt.println() ile yazdırırsak ne olur =

	// fmt.Println(range abc) | Hata veriyormuş :D
}
