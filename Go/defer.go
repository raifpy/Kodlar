package main

import "fmt"

func main() {
	fmt.Println("Merhaba , BetikSonu")
	defer fmt.Println("Sanada Merhaba Kardeşim")
	fmt.Println("Nasılsın , iyi misin ?")
}

// erteleme yapıyor .. | main'in sonrnda çalışacak ..

// Birden fazle eklenirse , en yeni eklenen ilk yazılır , kullanılır ..

// Ayrıca garantiye almak için de kullanulabilir // defer fileClose()
// Çalışam zamanı hatalarına sebep olabilir | RuntimeError
