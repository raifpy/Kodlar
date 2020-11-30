package main

import (
	. "fmt"     // fmt.Println() yapmak yerine Println() yapmış olacağız
	_ "os"      // kullanılmayan modülü disable etmek için _ koyuyoruz
	r "reflect" // modüle alias atamış olduk
)

func main() {
	abc := "Merhaba Dünya"
	Println(abc)
	Println(r.TypeOf(abc))

}
