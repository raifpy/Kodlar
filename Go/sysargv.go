package main

import (
	"fmt"
	"os"
)

func main() {
	abc := os.Args
	if len(abc) > 1 {
		fmt.Println(abc)
	} else {
		fmt.Println("Yanlış kullanım ! Lütfen arguman ekle ...")
	}

}
