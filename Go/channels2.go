package main

import (
	"fmt"
	"time"
)

func main() {

	kanal := make(chan string)

	go func() { // goruntime ile kullanılmak zorunda
		time.Sleep(time.Second * 3)
		kanal <- "Merhaba Dünya"
	}()

	fmt.Println(<-kanal)
	//time.Sleep(time.Second * 3)
	fmt.Println(<-kanal) // Bu çökecek

	// Çökmeyi de defer - recover () ile aşacağız 0_0
}
