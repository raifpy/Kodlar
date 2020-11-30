package main

import "fmt"

func main() {
	raif := "raif"
	fmt.Println(raif)

	arabaLastik(&raif)

	fmt.Println(raif)

}

func arabaLastik(isim *string) {
	*isim = fmt.Sprintf("Merhaba %s", isim)
}
