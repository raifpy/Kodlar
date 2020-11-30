package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.NewSyscallError("Olamaz Dostum Bu bir hata", os.ErrClosed)) // Olamaz Dostum Bu bir hata: file already closed

	// Custom * hata vermek için
}
