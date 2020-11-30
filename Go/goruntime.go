package main

import (
	"fmt"
	"time"
)

// thread yapısı gibi
func test(str string) {
	for i := 0; true; i++ {
		time.Sleep(1000 * time.Microsecond * 1000) // 1 saniye
		fmt.Printf("%s | %d\n", str, i)
	}
}
func main() {
	//test("ilk")
	//test("iki")
	//Bu dizim normal zaten . test("ilk") hiç bitmediği için ikinci test() asla başlmayacak

	go test("go ile")
	test("NORMAL")

	// go dediğimiz eleman go runtime oluyor . Arka planda devam ediyor ..

}
