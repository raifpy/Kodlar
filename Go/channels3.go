package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if hata := recover(); hata != nil {
			fmt.Println("Panic Yakalandı ! : ", hata, " main() baştan çağırılıyor .")
		}
	}()

	kanal := make(chan string) //

	run := func() {
		fmt.Println("Merhaba BetikSonu ")
		time.Sleep(time.Second * 5)
		kanal <- "BetikSonu"
		time.Sleep(time.Second * 5)
		kanal <- "@raifpy"

	}

	go run()
	fmt.Println(<-kanal) // Veri dönesiye kadar bekliyor !
	fmt.Println("Veri bekleniyor")
	fmt.Println(<-kanal) // fatal error | Bütün goruntimelar ölü diyor abm
	fmt.Println("Veri bekleniyor")
	//fmt.Println(kanal)
	//fmt.Println(kanal)
	//fmt.Scanln()

}

// Eğer bir chan varsa ve fmt.println(<-chan) tanımlaması var ise
// fmt. Chan'a değer dönesiye kadar bekleyecek . Eğer değer dönecek GoRunTime çalışmıyorsa bam fataaal
