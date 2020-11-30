package main

import (
	"fmt"
	"time"
)

func main() {
	chantime := time.NewTimer(time.Second * 5) // 5 saniyelik bir timer oluştu (chantime)
	fmt.Println("Timer beklenecek")
	fmt.Println(<-chantime.C)
	fmt.Println("Bittu -> ", !chantime.Stop())

}
