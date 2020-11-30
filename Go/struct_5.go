package main

import "fmt"

type insan struct {
	isim string
	nick string
}

func main() {
	raif := insan{"Ömer Raif Tekin", "raifpy"} // Bu şekilde de kullanabiliyoruz :)
	fmt.Println(raif)
	fmt.Println(raif.isim)
	fmt.Println(raif.nick)

}
