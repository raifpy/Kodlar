package main

import "fmt"

func main() {
	var abc = map[string]int{}
	abc["a"] = 1
	fmt.Println(abc["a"])
	fmt.Println(abc["b"]) // int'in null hali 0
	var deger, boolean = abc["b"]
	fmt.Println(deger, boolean) // Aynı biçimde

	for deneme, deneme2 := range abc {
		fmt.Println(deneme, deneme2) // keyleri dönüyor , güzel |   | key | value
	}

}
