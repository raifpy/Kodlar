package main

import "fmt"

// işlem geçerli oldukça for'a devam edecek
func main() {
	sayı := 1 // Bu arada Go UTF-8 destekliyor . Yani değişkenleri türkçe yapabilirsiniz .)
	fmt.Println("Şuanda sayı : ", sayı)
	for sayı < 10 { // Sayı 10'dan küçük olduğu süre boyunca koşul doğrulanacak
		fmt.Println(sayı, " 10'dan küçük , bir arttırılıyor ! ")
		sayı++
	}
	fmt.Println("Döngü sona erdi : ", sayı)

}

// Python'ca da ;

//sayı = 1
//while sayi<10:
//	print("Sayı 10 'dan küçük , bir arttırılıyor")
//	sayı += sayı
//print("Bitti")

// Java'ca da ;

//int sayi = 1
//while (sayi < 10){
//	System.out.println("Sayı 10'dan küçük , bir arttırılıyor")
//	sayi++
//}
//System.out.println("Bitti")

// @raifpy
