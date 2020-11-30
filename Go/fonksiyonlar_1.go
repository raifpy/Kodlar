package main

import "fmt"

var selam = "Selam"

func main() {
	saySelam()
	setSelam("Merhaba Dünya !")
	saySelam()
}
func saySelam() {
	fmt.Println(selam)
}
func setSelam(msj string) { selam = msj }

// Bunlar basit .)
