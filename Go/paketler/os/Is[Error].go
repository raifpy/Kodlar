package main

import (
	"fmt"
	"os"
)

// err ile kullanıyor , eğer hata isexist ise true dönecek

func main() {
	fmt.Println(os.IsExist(os.ErrExist)) // this şit already exists

	// Keza IsNotExists'da onun tersi

	fmt.Println(os.IsNotExist(os.ErrNotExist)) // this şit not avaible

	fmt.Println(os.IsPermission(os.ErrPermission)) // Bu izin hatası mı ? evet . Tamam true dönüyorum .

	//gibi gibi devamı var
}
