package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "LinUx ÇoK gÜvEnli yİğEnİm VirÜs MiRüs YkO"
	fmt.Println(str)
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str)) // ASCII standartına göre büyük oldu , i İ olması gerekirken I oldu

	tr := unicode.TurkishCase

	fmt.Println(strings.ToUpperSpecial(tr, str)) // Düzeldi :)
}
