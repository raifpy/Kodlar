package main

import (
	"fmt"
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	errHandler.HandlerExit(os.Symlink("testFile", "testFile2"))
	fmt.Println("testFile2 (kısayol) oluşturuldu")

}

// Kısayol oluşturuyor .
