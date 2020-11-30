package main

import "fmt"

type kisi struct {
	isim string // alttaki gibi de olur ama şimdilik olmasın :)
	//isim string = "raifpy"
	yas      int
	boy      float32
	memleket string
}

func main() {
	raifpy := kisi{}
	raifpy.isim = "Ömer Raif Tekin"
	raifpy.memleket = "Konyam konyam g-güzel konyam .."
	raifpy.yas = 16
	raifpy.boy = 186
	fmt.Println(raifpy)

	// Py 'da nesneye benzere bir eleman bu struct yapısı

	eleman := kisi{}
	eleman.isim = "Eleman"
	eleman.memleket = "raif@neon"

	eleman2 := kisi{}
	eleman2.isim = eleman.isim
	eleman2.yas = eleman.yas
	eleman2.memleket = eleman.memleket

	fmt.Println(eleman) // Çok hoş bir yapısı var

	fmt.Println(eleman.isim == eleman2.isim) // true çıktı . Güzel :
}
