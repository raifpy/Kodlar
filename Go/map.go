package main

import "fmt"

func main() {
	abc := map[string]string{}
	abc["Merhaba"] = "Dünya"
	fmt.Println(abc["Merhaba"])
	deger, boolean := abc["Dünya"]
	fmt.Println(deger, boolean)
	fmt.Println(abc)
}
