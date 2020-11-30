package main

import (
	"log"
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	if errHandler.HandlerBool(os.Link("samefile.go", "testGo.go")) {
		log.Println("Görünüşe göre testGo.go hali hazırda mevcut . Es geçiliyor")
	}

	file1, err := os.Stat("testFile")
	errHandler.HandlerExit(err)
	file2, err := os.Stat("testGo.go")
	errHandler.HandlerExit(err)

	log.Println(os.SameFile(file1, file2)) // false dedi :))
}

/*
SameFile reports whether fi1 and fi2 describe the same file.
For example, on Unix this means that the device and inode fields of the two underlying structures are identical; on other systems the decision may be based on the path names.
SameFile only applies to results returned by this package's Stat.
It returns false in other cases.
*/
