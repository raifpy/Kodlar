package main

import (
	"fmt"
	"io/ioutil"
)

// 1 'de os.Open() / Read() 'i kullanmıştık . Şimdi ise io / iotuil 'i kullanacağız .

func main() {
	icerik, err := ioutil.ReadFile("dosya")
	if err != nil {
		fmt.Println("Olamaz , hata : ", err.Error())
		return
	}

	fmt.Println(icerik)
	fmt.Println(string(icerik))

}

// Bu daha basit ..
