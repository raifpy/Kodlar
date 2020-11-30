package main

import "fmt"

func main() {
	var liste [1]string
	fmt.Println(liste)
	liste[0] = "Mewrhaba"
	fmt.Println(liste)

	var liste2 = [3]string{"Merhaba", "Dünya", "Hey"}
	fmt.Println(liste2)

	liste2[2] = "Ulan my\f\x1B[32mql\x1B[0m"
	fmt.Println(liste2)
}
