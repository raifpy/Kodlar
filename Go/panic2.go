package main

import "fmt"

func run() {
	// defer ErrorHandler() // recover'ı başka fonksiyondan tanımlayamıyoruz . iç içe fonksiyon olması gerekiyor , olsun bakalım
	ErrHndler := func() {
		if hata := recover(); hata != nil {
			fmt.Println("Olamaz Hata ! := ", hata)
		}
	}
	defer ErrHndler()
	panic("Olamaz panikledim")
}

/*
func ErrorHandler() {
	if abc := recover(); recover != nil {
		fmt.Println("Hata yakalandı ! : ", abc)
	}
}
*/

func main() {
	run()
}
