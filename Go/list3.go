package main

import "fmt"

func main() {
	//var abc = [2]string{}                                            // 2 elemanlı string liste
	var abcd = []string{"a", "b", "c", "d", "e", "f", "g", "h", "j"} // eleman sayısının belirtmedik , {}'dan sonrasına göre kendisi ayarlayacak

	//var abcde = append(abc, abcd...) // abc'ye ekleyemedi , hata verdi . abc'nin slice olmasını istiyor (sayı belirtme diyor yani)

	var b = []string{}              // slice liste atadım
	fmt.Println(append(b, abcd...)) // ... listenin içeriğini b'ye ekle demek
	fmt.Println(b)
	fmt.Println(abcd)
	/*
		[a b c d e f g h j]
		[]
		[a b c d e f g h j]
	*/

	// Bu da demek oluyor ki append ile return oluyor !
	// Doğrudan eklemek için copy() var .

	sayi := copy(b, abcd) // Bu kaç elemanın kopyalayındığını return eder , çok mühim değil

	fmt.Println(sayi) // 0 gözüküyor , yani hiç kopyalanmamış
	fmt.Println(b)
	fmt.Println(abcd)

	fmt.Println(cap(b)) // cap() kapasiteyi verir // Bunun kapasite 0 :D
	len_b := len(abcd)
	bb := [len_b]string{}
	fmt.Println(bb)

	// fmt.Println(abcd[len(abcd)+1]) // out of range verdi (çöktü)

}
