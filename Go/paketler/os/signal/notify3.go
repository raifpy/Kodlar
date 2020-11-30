package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func handler(kanal chan os.Signal) {
	eleman := <-kanal                                                          // kanala öge gelene kadar bu bizi oyalayacak
	fmt.Printf("Olamaz , birşeyler bizi engelliyor (%T) %v\n", eleman, eleman) // %T type(reflect.TypeOf) elemanını veriyor .))
}

func main() {
	kanal := make(chan os.Signal)
	signal.Notify(kanal)
	go handler(kanal)
	for {
		fmt.Println("Şuan burada iş yapıyor !")
		time.Sleep(time.Second * 10)
	}
}

// Bu şekilde kullanım sonucunda signal.Notify(kanal) ne bulursa (tüm sinyalleri) handler elemanına yönlendiriyor :)
// signal.Ignore(................) yapmak yerine bunu kullanmak daha mantıklı olur (kötücül olmayan kötücül yazılımlar için :)
// Mantık olarak da aynı işi yapıyorlar zaten
