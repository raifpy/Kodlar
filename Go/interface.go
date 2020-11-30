// Henüz ne işe yaradığını çok iyi kavrayamadım
package main

import "fmt"

type nesne struct {
	isim    string
	soyisim string
}

func (eeeyy nesne) Birlestir() string {
	return eeeyy.isim + " : " + eeeyy.soyisim
}

type deneyelim interface {
	Birlestir() string
}

func getDeneyelim(a deneyelim) {
	fmt.Println(a.Birlestir())
}

func main() {
	var abc = nesne{"Hey", "BetikSonu"}
	getDeneyelim(abc)
}
