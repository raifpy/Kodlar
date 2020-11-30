package main

import "fmt"

func main() {
	var _1 = "Merhaba"
	_2 := "Dünya"
	fmt.Println(_1, _2)  // , yapılarını kullandım
	fmt.Println(_1 + _2) // + yapısını kullandım

	//fmt.Println(_1 * 10) // Hata veriyor .
	//fmt.Println(_1 - _2) // Hata veriyor .

	_3 := 10
	var _4 = 20

	fmt.Println(_3 + _4) // Toplama işlemi yapmış olduk (+)
	fmt.Println(_3, _4)  // , kullandığimız için sadece yazdı

	//fmt.Println(_2+_3) // string and int hata verdi (doğal olarak)
	fmt.Println(_2, _3) //

	// Python yapısı ile aynı |

}
