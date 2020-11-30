package main

import (
	"fmt"
)

type nesne struct {
	isim, soyisim string
	yil, ay       uint // uint negatif sayı alamaz .
}

func main() {
	abc := nesne{isim: "raifpy", ay: 0}
	fmt.Println(abc)
	abc.soyisim = "Gocuoğulları"
	fmt.Println(abc)

}
