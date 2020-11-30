// []uint8 dediğimiz eleman []byte .
// C'de string yerine []byte olduğunu unutmayalım

package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "Merhaba Dünya"
	fmt.Println(str, " | ", reflect.TypeOf(str)) // Merhaba Dünya | string
	uint8_ := []uint8(str)
	fmt.Println(uint8_, " | ", reflect.TypeOf(uint8_)) // []byte elemanlar | uint8
	// uint8 to string
	fmt.Println(string(uint8_)) // oldukça kolay

	fmt.Println([]byte(str)) // Buda []byte hali , aralarında fark yok

}
