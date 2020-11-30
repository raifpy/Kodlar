package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains ile aynı ama int değerinden veri alıyor
	fmt.Println(string(97)) // a
	fmt.Println(string(98)) // b
	str1 := "Go Kardeşim"
	fmt.Println("STR    : ", str1)
	fmt.Println("CoR 97a: ", strings.ContainsRune(str1, 97))
	fmt.Println("CoR 98b: ", strings.ContainsRune(str1, 98))
}
