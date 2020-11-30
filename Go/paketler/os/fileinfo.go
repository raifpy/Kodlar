//fileinfo
package main

import (
	"fmt"
	"os"
)

func main() {
	FileInfo, _ := os.Stat("fileinfo.go")                // FileInfo , err
	fmt.Println("IsDır    : ", FileInfo.IsDir())         // dir mi [bool]
	fmt.Println("ModTime  : ", FileInfo.ModTime())       // Yapılan en son değişiklik zamaı
	fmt.Println("Perm     : ", FileInfo.Mode().String()) // unix iznleri
	fmt.Println("Name     : ", FileInfo.Name())          // isim
	fmt.Println("Size     : ", FileInfo.Size())          // boyutu ( byte cinsizinde)
}
