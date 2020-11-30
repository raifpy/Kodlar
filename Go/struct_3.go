package main

import "fmt"

func main() {

	// main içinde tanımlayalım birde :)
	abc := struct {
		isim, memleket string
	}{"raif", "konya"}

	//abc.isim = "raif"
	//abc.memleket = "Gonya" // tabiki sonradan da değiştirilebilir

	fmt.Println(abc)
	fmt.Println(abc.isim)
}
