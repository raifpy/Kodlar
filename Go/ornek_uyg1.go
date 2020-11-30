package main

import "fmt"

type eleman struct {
	isim     string
	memleket string
	yas      int
}

func (a *eleman) SetYas() bool {
	if a.isim == "" {
		return false
	}
	var yas int
	fmt.Printf("Merhaba %s Adın ne acep ? ", a.isim)
	fmt.Scanln(&yas) // Şimdilik str girişimini göz ardı edelim
	a.yas = yas
	return true
}

func main() {
	var raif = eleman{"raif", "konya", 1}
	raif.SetYas()
	fmt.Println(raif)

}
