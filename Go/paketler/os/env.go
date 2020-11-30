package main

import (
	"fmt"
	"os"
)

func main() {
	for _, eleman := range os.Environ() {
		fmt.Println(eleman)
	}

	os.Clearenv() // Tüm env'i sildi .

	fmt.Println("\n--\n", os.Environ(), "Yok")

	os.Setenv("raifpy", "Burada .") // Ekleme yapabiliyoruz

	fmt.Println(os.Environ())

	fmt.Println(os.Getenv("raifpy"))

	os.Unsetenv("raifpy") // unset edildi ( temizlendi )

	fmt.Println(os.Getenv("raifpy"))
}
