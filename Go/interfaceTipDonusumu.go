package main

import "fmt"

func verBanaMap(abc interface{}) {
	fmt.Printf("%T \n", abc)
	veri, ok := abc.(map[string]interface{})
	fmt.Println(veri, ok, "\n")
}

func main() {
	verBanaMap("merhaba dünya")                // false
	verBanaMap(map[uint]string{10: "Merhaba"}) // false
	verBanaMap(map[string]uint{"Merhaba": 10})
	verBanaMap(map[string]string{"Merhaba": "Merhaba"})
	verBanaMap(map[string]interface{}{"Merhaba": "Merhaba"})
}
