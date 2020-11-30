package main

import (
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	// ilk dosya içeriğini olduğu gibi ikinci dosyaya kayıt ediyor
	// copy ile aynı
	errHandler.Handler(os.Link("oldFile", "newFile"))
}
