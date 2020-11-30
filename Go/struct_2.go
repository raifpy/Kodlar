package main

import "fmt"

type kisi struct {
	isim string
}

func main() {
	abc := kisi{"Raifpy"} // Doğrudan tanımalamada yapabiliyoruz
	fmt.Println(abc.isim)
}
