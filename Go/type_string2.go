package main

import "fmt"

type str string

func (a str) getHalf() str {
	return a[(len(a) / 2):]
	// Burdaki a elemanı string ama str olarak geçiyor
}

func main() {
	var abc str = "Merhaba Dünya"
	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(abc.getHalf())
}

// Olarak kullanılabilir
