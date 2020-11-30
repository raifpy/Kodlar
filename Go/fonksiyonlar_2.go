package main

import "fmt"

var deger = 10
var degisti = false

func main() {
	printDeger()
	setDeger(100)
	printDeger()
	setDeger(12.0)
	printDeger()

}

func getDeger() int {
	return int(deger)
}

func setDeger(x int) {
	deger = x
}
func printDeger() {
	fmt.Println(getDeger())
}
