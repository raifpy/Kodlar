package main

import (
	"fmt"
	"os"
)

func main() {
	var dosya = "dosya2"
	if _, err := os.Open(dosya); os.IsNotExist(err) {
		fmt.Println("Hop , dosya yok be kardeşim , senin için hemen oluşturuyorum")
		dosya, err := os.Create(dosya)
		fmt.Println(dosya, err)
		//fmt.Println("Dosya oluşturuldu , main tekrar çağırılıyor !")
		//main()
	} else {
		fmt.Println("Dosya zaten var , siliniyor ..")
		os.Remove(dosya)
	}

}

// os.Create() oluşturmak , os.Remove() silmek için .
// Ayrıca os.Create() ile oluşan dosya , değere atanıp oluşturulduğu gibi kullanılabilir
