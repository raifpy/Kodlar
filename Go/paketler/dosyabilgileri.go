package main

import (
	"fmt"
	"os"
)

func main() {
	dosya, err := os.Open("dosya")
	defer dosya.Close()
	if err != nil {
		fmt.Println("Hata : ", err.Error())
		return
	}

	bilgi, _ := dosya.Stat()
	fmt.Println("Dosya ismi      : ", bilgi.Name())          // dosya ismi
	fmt.Println("Mod ( izinler ) : ", bilgi.Mode().String()) // dosya izmi ()-rwxrwxr-x
	fmt.Println("Son değişiklik  : ", bilgi.ModTime().String())
	fmt.Println("Klasör mü ?     : ", bilgi.IsDir())
	fmt.Println("Boyut (bayt)    : ", bilgi.Size())

}
