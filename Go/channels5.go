package main

import (
	"fmt"
	"time"
)

func WhileChan(kanal chan string) {
	var a = 0
	for {
		fmt.Println(a)
		kanal <- "Şimdi : " + time.Now().String()
		a++
	}

}

func main() {
	zamanChan := make(chan string)
	go WhileChan(zamanChan)
	for {
		fmt.Println(<-zamanChan)

		time.Sleep(time.Second)

	}
}

// Bu da demek oluyor ki WhileChan() 'da Sayı çok hızlı artması gerekirken yavaş yavaş artıyor . Yani main() 'den <- zamanCahn 'ın alınması bekleniyor
//
