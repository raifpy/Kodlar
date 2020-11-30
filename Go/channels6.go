// Hafızalu kanallar | tamponlu
// goroutine kullanmaya gerek yok
// range ile de kullanalabilir

package main

import "fmt"

func get(tampon chan int) {
	for i := 0; i < 5; i++ { // 6 yaparsak fatal error . Çünkü bizim eleman 5 alabilir
		tampon <- i
	}
}

func main() {
	tampon := make(chan int, 5)
	get(tampon)
	close(tampon)                          // kanalı kapatıyor
	fmt.Println("kanal tanımlama tamam !") // normal kanal gibi beklemedi
	fmt.Println("Len : ", len(tampon))     // 5
	fmt.Println(<-tampon)
	fmt.Println(<-tampon)
	// Sırayla ..
	fmt.Println("Len : ", len(tampon)) // 3  // <- tampon dedikçe eleman sayısı azalıyor
	fmt.Println("CAP : ", cap(tampon))
	// kalanları da range ile yazalım

	for i := range tampon {
		fmt.Println("in range : ", i)
	}

	// Kanal kapandığı için hata vermedi , öbür türlü fatal

}
