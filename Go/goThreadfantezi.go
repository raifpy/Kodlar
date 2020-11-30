package main

import "fmt"

func runMe(a int) {
	for {
		fmt.Printf("Merhaba Dünya %d\n", a)
	}
}

func okRun() {
	for i := 0; i < 100; i++ {
		go runMe(i)
	}
}

func main() {
	fmt.Println("Dikkat , yüksek ihtimal ile işletim sistemi yanıt vermeyecek . Çökecektir .")
	fmt.Scanln()

	for {
		go okRun()
	}

	//fmt.Scanln()
}
