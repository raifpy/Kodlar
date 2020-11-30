package main

import (
	"fmt"
	"reflect"
	"time"
)

func VerBanaMuz(kanal chan string) {
	fmt.Println("Yakala kardeşim muz cumhuriyeti")
	kanal <- "Bu ülkeye bu mecralar yakışmıyor yaauv "
	fmt.Println("Hayır barolar dışarıda bekleyecek . Avukatlar meclise giremez saçmalama .d")
	time.Sleep(time.Second)
	kanal <- "Fetöcu botu dislike attı ama teyip trendlere çıktı xd - yeniakil.com.tr"
	kanal <- "Z kUşAğıNı iLgİleNdiReN HaBeR bilmemkimamaRockçu ekperiye DeSkTop VerDi "
}

func main() {
	kanal := make(chan string)
	go VerBanaMuz(kanal)
	fmt.Println(<-kanal)
	son := <-kanal
	fmt.Println("Haber : ", son, " | Tür ", reflect.TypeOf(son))
	fmt.Println("Haber : ", <-kanal, "| Tür bas ışığa mınagoyim")
	fmt.Println("Bitti kardşeim nasıl beğendim mi ülkemi :))")
	fmt.Scanln()
}
