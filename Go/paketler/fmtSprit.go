package main

import (
	"fmt"
	"reflect"
)

func print(text interface{}) {
	fmt.Println(text)
}
func main() {
	print("Merhaba")
	print(reflect.TypeOf("Merhaba"))
	print("i-i-i-i-i-i-i")
	var deneme = fmt.Sprintln("Merhaba Sprint")
	print(deneme)
	var deneme2 = fmt.Sprintln(reflect.TypeOf("Dünya"))
	print(deneme2)
	print(reflect.TypeOf(deneme2))

	// Bu arkadaşın olayı ekrana yazılan nesneyi (her ne olursa) string olarak vermek

	var kemiksizkaka = fmt.Println // Pointer olacak
	var kemiksizkaka_string = fmt.Sprint(kemiksizkaka)
	fmt.Println(kemiksizkaka_string, reflect.TypeOf(kemiksizkaka_string)) // pointer_kodu string

	// Çok güzel bir eleman

}
