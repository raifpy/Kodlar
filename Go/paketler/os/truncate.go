package main

import (
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	// Turncatte changes size
	errHandler.HandlerExit(os.Truncate("testFileSize", 500000))
}

// Ciddi ciddi boyutunu 500000 yaptı | 500 kb
// fallocate ile aynı eleman
