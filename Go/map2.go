package main

import "fmt"

func main() {
	var abc = map[string]string{}
	abc["hey"] = "Merhaba Dünya"

	fmt.Println(abc)
	fmt.Println(abc[""])

	var abc2 = make(map[string]string, 100) // 100 elemanlı map

	fmt.Println(abc2)

	abc2["a"] = "B"
	abc2["b"] = "A"

	fmt.Println(abc2)
	fmt.Println(abc2["a"])
	fmt.Println(abc2["b"])
	fmt.Println(abc2["OLMAYAN DEGER"])

	var sonuc, boolean = abc2["OLMAYAN DEGER"]
	fmt.Println(sonuc, boolean)
}
