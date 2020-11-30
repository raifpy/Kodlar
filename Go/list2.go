package main

import "fmt"

func main() {
	abc := [2]string{} // 2 elemanlı bir string liste oluşturduk
	fmt.Println(abc)
	abc[0] = "Hey"
	abc[1] = "Dünya"
	fmt.Println(abc)
	//abc[2] = "Hata Verecek"

}
