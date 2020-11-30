package main

import (
	"fmt"
	"os"
	"time"
)

func test(s string) {
	for i := 0; i < len(s); i++ {
		time.Sleep(time.Second)
		fmt.Println(s[:len(s)-i])
	}
}

func exit() {
	fmt.Println("Exit veriliyor bakalım programdan mı thread'den mi çıkacak")
	os.Exit(0)
	fmt.Println("Çıkmadı .")
}
func main() {
	go test("Dünya") // Bu da demek oluyor ki go .. 'dan sonra başka birşey olmazsa exit veriyor.
	// Tıpkı Python threading gibi
	time.Sleep(time.Second * 2)
	go exit()
	test("Devam")

}

// Bu da demek oluyor ki go ... sonra betik biterse exit verir . go .. içinde exit verilirse gine programdan çıkar
