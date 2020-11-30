package main

import (
	"fmt"
	"time"
)

func main() {
	kanal := make(chan string)
	go func() {
		kanal <- "Merhaba Dğnya"
		fmt.Println("Veri alındı , Kanal kapanıyor !")
		close(kanal)
	}()
	fmt.Println(kanal)
	time.Sleep(time.Second * 3)
	fmt.Println(<-kanal)
	fmt.Println(<-kanal) // kanala başka veri atanmadı neden fatal vermedi :))
	fmt.Println("İşlem bitti .")

	// close() ile kanala başla veri yollanmayacağını belirtmiş olduk .
	// Başka veri olmayacağı için <-kanal "" dönüyor
}
