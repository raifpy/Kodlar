package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Şuanki PID : ", pid)
	fmt.Scanln() // Bekleyelim ki doğrulayalım
}
