package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "Go Kardeşim"
	fmt.Println("STR    : ", str1)
	fmt.Println("CoA a  : ", strings.ContainsAny(str1, "a"))
	fmt.Println("CoA ae : ", strings.ContainsAny(str1, "ae"))
	fmt.Println("CoA '' : ", strings.ContainsAny(str1, ""))
	fmt.Println("CoA ' ': ", strings.ContainsAny(str1, " "))
	fmt.Println("CoA'  ': ", strings.ContainsAny(str1, "  "))

}

// str1 , "abcd" | Eğer abcd hartflerininden birisi str1 'in içerisinde var ise true dönecek

// Contains kullanmak daha doğru olur ama Any 'de işe yarayacak ..
