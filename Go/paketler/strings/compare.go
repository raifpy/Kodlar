package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "Merhaba Dünya"
	var str2 = "Dünya"

	sonuc := strings.Compare(str1, str2)
	fmt.Println("Sonuç  : ", sonuc)
}

/*
Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

Compare is included only for symmetry with package bytes. It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.
*/
// Anlayabilmiş değilim ne işe yaradığını
