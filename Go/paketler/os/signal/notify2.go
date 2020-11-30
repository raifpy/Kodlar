package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

func exit() { // exit() çağıralarak kullanıcıdan
	fmt.Println("\nGerçekten çıkmak istiyor musun ? (\033[31mE\033[0m/\033[32mH\033[0m)") // are u sure metini
	var cevap string                                                                      // scanln için değişken oluşturduk
	fmt.Scanln(&cevap)                                                                    //kullanıcıdan veri aldık
	cevap = strings.ToLower(cevap)                                                        // cevap ne olursa olsun küçülttük
	if cevap == "e" {                                                                     //eğer ki cevap e ise (eleman çıkmak istiyor ise)
		os.Exit(0) // Exit verdik , os.Exit(0) hatasız çıkış olduğunu belirtir
	} else { // değilse (eleman kalmal istiyor ise)
		log.Println("Devam ediliyor !") // Devam ediliyor yazdırdık ...
	}

}

func interrupt(uyari chan os.Signal) { // goroutine ile çalıştırmak için interrupt adında içerisinde os.Signal channel ögesini alacak fonkisyon tanımlıyoruz
	for { // Infinity Loop || Uygulama içerisinde tekrar ve tekrar ctrl-c'ye basım olabilir
		<-uyari // ctrl-c yakalanasıya kadar burada bekliyor olacağız
		exit()  // ctrl-c yakalandıktan sonra exit çalışacak
	}

}

func main() { // Herşey main'de başlar

	uyari := make(chan os.Signal) //ctrl-c 'yi yakalamak için os.Signal chan'İ oluşturmuş olduk
	//var uyari chan os.Signal      // BU yapı da kullanabilir
	signal.Notify(uyari, os.Interrupt) // os.Interrupt geldiğinde uyari kanalaına mesaj gidecek
	go interrupt(uyari)                // goruotine ile çalıştırdık
	for {                              // Bu ve aşağısındaki elemanlar işi temsil ediyor .
		fmt.Println("işlem yapılıyor")
		time.Sleep(time.Second * 5)
	}

}

//Hala çok kullanışlı değil ama idare edecek bizi :)
