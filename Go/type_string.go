package main

import (
	"fmt"
	"reflect"
)

type str string

func main() {
	var a str = "Merhaba Dünya"
	fmt.Println(a)                 // Merhaba Dünya
	fmt.Println(reflect.TypeOf(a)) // main.str
}
