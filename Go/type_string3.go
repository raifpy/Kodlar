package main

import "fmt"

type eleman string

func (elem *eleman) chanceEleman(a string) {
	var iki eleman = a
	fmt.Println(iki)
	//*elem =
}

func main() {
	var abc eleman = "Merhaba Dünya"
	abc.chanceEleman("Değiş !")
	fmt.Println(abc)
}
