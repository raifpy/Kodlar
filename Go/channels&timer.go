package main

import (
	"fmt"
	"time"

	"github.com/mndrix/rand"
)

var sec = 1000000000

func TakeWork(kanal chan string) {
	uyu := time.Second * time.Duration(rand.Intn(5))
	fmt.Println("İş başladı : ", uyu)
	time.Sleep(uyu)
	time.Sleep(time.Duration(uyu))
	kanal <- "İş Bitti"
}

func CheckTimeOut(kanal chan string) {

	uyu := time.Second * time.Duration(rand.Intn(5))
	fmt.Println("Check başladı : ", uyu)
	time.Sleep(uyu)
	kanal <- "İş Belirtilen sürede tamamlanamadı : " + uyu.String()

}

func main() {
	kanal1, kanal2 := make(chan string), make(chan string)
	go TakeWork(kanal1)
	go CheckTimeOut(kanal2)
	//fmt.Println(<-kanal1) // checki almak için mecburen işin bitmesini bekleyeceksin .
	//fmt.Println(<-kanal2) // Bunun bir anlamı yok

	select {
	case eleman := <-kanal1:
		fmt.Println("\033[32mİş bitti\033[0m , log : ", eleman)

	case eleman := <-kanal2:
		fmt.Println("\033[31mTimeout !\033[0m : ", eleman)
	}

	time.Sleep(time.Second * 2)
	//fmt.Println("\033[2J")
	main()

}
