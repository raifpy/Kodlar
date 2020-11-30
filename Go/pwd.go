package main

import "fmt"
import "os"

func main(){
	fmt.Println("Başlıyor")
	konum := os.GetPwd(-1)
	fmt.Println(konum)
}
