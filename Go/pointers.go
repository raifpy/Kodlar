package main

import "fmt"

func main() {

	var str = "Merhaba Dünya"
	//var str2 = &str
	//var str3 = *str

	fmt.Println(str, &str, *&str)

	// & adresi , * de adresi değere çevirir
}
