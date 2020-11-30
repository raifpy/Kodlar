package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "Go Kardeşim"
	fmt.Println("STR   : ", str1)
	fmt.Println("Co a  : ", strings.Contains(str1, "a"))
	fmt.Println("Co ae : ", strings.Contains(str1, "ae"))
	fmt.Println("Co '' : ", strings.Contains(str1, ""))
	fmt.Println("Co ' ': ", strings.Contains(str1, " "))
	fmt.Println("Co'  ': ", strings.Contains(str1, "  "))

}

// if "abc" in "abcd" ifadesinin Go 'cası
// ContainsAny ile arasındaki fark "str1" ,"stsfsdf" 'ın true çıkması .
// Yani Any'de aranacak kelime kelime olarak değil , harf olarak aranıyor

// Contains = içermek
