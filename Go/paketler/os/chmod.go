package main

import (
	"os"

	"github.com/raifpy/Go/errHandler" // Hata ver ise ekrana yazıyor , benim için büyük kolaylık , kullanmasanızda olur
)

func main() {
	errHandler.HandlerExit(os.Chmod("testFile", 0)) // err | 0 tüm izinelri yok ediyor
}
