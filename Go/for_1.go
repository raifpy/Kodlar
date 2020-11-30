// Go'da sadece for döngüsü var :)
// Mantık olarak C-pp / Java aynı ;

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// C söz diziminde () arasında , int i = 0 olarak tanımlanır , burada azıcık daha kolay (kimine göre daha zor yapısı var)

	// := yerine var i = 0 'da kullanabilirdik ; |- Kullanabiliriz diye düşünmüştüm ama olmuyormuş :D

	for i := 0; i<10; i++ {
		fmt.Println(i)
	}

	// fmt.Println(i) // not definied . Çünkü i , for döngüsü içerisinde tanımlandı . Java'daki kurallara aynı devam ..

	//var i = 0
	//for i; i < 10; i++ {
	//	fmt.Println(i)
	//} // Hata veriyor bu arkadaş :D

	// Hop bu olur
	//i = i + 1
	//fmt.Println(i)
}
