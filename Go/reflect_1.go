package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Bu go'nun ide'si gerçekten kaliteli")
	var deneme = reflect.TypeOf("Hmn")
	fmt.Println(deneme)
	fmt.Println(reflect.TypeOf(deneme))
	//var deneme_str = string(deneme)
	//fmt.Println(deneme_str)

	//reflect.Ty
}
