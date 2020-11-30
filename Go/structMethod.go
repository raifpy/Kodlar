package main

import "fmt"

type nesne struct {
	isim, tur string
}

func (nes nesne) birlestir() string {
	return nes.tur + " : " + nes.isim
}

func main() {
	abc := nesne{"Bilgisayar", "Teknoloji"}
	abc2 := nesne{tur: "Kimya", isim: "Sabun"}
	fmt.Println(abc)
	fmt.Println(abc.birlestir())
	fmt.Println(abc2.birlestir())
}
