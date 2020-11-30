// C'deki ile tamamen aynı || break dışında

package main

import "fmt"

func main() {
	abc := 10
	switch abc {
	case 1:
		fmt.Println("Merhaba , Bu sayı 1'dir")
	case 2:
		fmt.Println("Merhaba , bu sayı 2 'dir")
	case 3:
		fmt.Println("Bu 3'dür ..")
	default:
		fmt.Printf("%v sayısı , switch'de bulunamadı !\n", abc)

	}
}
