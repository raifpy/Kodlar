package main

import "fmt"

func main() {
	var abc string
	fmt.Print("Merhaba adın ne ? ")
	fmt.Scanln(&abc)
	fmt.Printf("Merhaba %s\n", abc)

	// fmt.Scanln(pointer-string)
	// &abc ile abc'nin pointerına ulaşabiliyorduk ..
	// Ama boşluk bırakıp enterlarsak sorun çıkarıyor ( adın ne ? merhaba deneme) // Sadece merhabayı alıyor , denemeyi' bash'e veriyor

	// Birde fmt.Scan(&abc) 'yi deneyelim

	fmt.Printf("Peki soyadın nedir %s ? ", abc)
	var soyad string
	fmt.Scan(&soyad)
	fmt.Printf("Merhaba %s %s", abc, soyad)

	// Hayır değişiklik yok :D
	// Tabiki tek input Scan... değil .
}
