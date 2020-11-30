package main

import "fmt"

// Py'daki *argv yapısı ile ilgileneceğız .
func main() {
	var abc = OrtalamaBul(1, 2, 3, 6, 3, 1, 2, 3, 4, 6, 7, 2, 4)
	fmt.Println(abc)
}

func OrtalamaBul(x ...int) float32 {
	sayi := len(x)
	toplam := 0
	for _, deger := range x {
		toplam += deger
	}
	return float32(toplam) / float32(sayi)
}

// x ...[tür] yapısı ile istediğimiz kadar değişleni alabiliriz (liste olarak alıyoruz muhtemlen)
