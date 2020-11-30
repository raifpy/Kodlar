package main

import (
	"fmt"
	"time"
)

// çoklu goruoutine leri kontrol ederken select-case yapısını kullanıyoruz
func bir(kanal chan string) {
	fmt.Println("Bir başladı .")
	time.Sleep(time.Second * 3)
	kanal <- "Merhaba Bir"
	fmt.Println("Bir \033[31mbitti\033[0m")
	//close(kanal)
}

func iki(kanal chan string) {
	fmt.Println("İki başladı .")
	time.Sleep(time.Second * 5)
	kanal <- "Merhaba İki"
	fmt.Println("İki \033[31mbitti\033[0m")
	//close(kanal)
}

func main() {
	kanal1 := make(chan string)
	kanal2 := make(chan string)
	go bir(kanal1)
	go iki(kanal2)

	//fmt.Println(<-kanal2) // bire erişmek için illaki ikiyi bekleyeceğiz :/
	//fmt.Println(<-kanal1) // bir
	for i := 0; i < /*3*/ 2; i++ { // döngüyü 3 yaptım ama bizden max 2 çıkabilir . Program çökecek | fatal |all goroutines are asleep - deadlock!

		fmt.Printf("\033[32mDöngü\033[0m %d\n", i+1)
		select { // Ama burda beklemiyorz. Hangisi önce çıkarsa o kullanılıyor select{} den çıkılıyor
		case hesap := <-kanal2:
			fmt.Println(hesap)
			//close(kanal2)

		case hesap := <-kanal1:
			fmt.Println(hesap)
			//close(kanal1)
		}
	}
	// Böyle döngüye sokarak da sonuncuyuda elde edebiliyoruz

	fmt.Println("Son")
	fmt.Scanln()

}
