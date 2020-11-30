package main

import (
	"fmt"
	"os"
)

func main() {

	// cacheDir'i dönüyor

	fmt.Println(os.UserCacheDir()) // dir , err
}
