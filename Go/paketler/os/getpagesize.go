package main

import (
	"fmt"
	"os"
)

func main() {
	page := os.Getpagesize()
	fmt.Println(page)
}

//Getpagesize returns the underlying system's memory page size.
//Getpagesize, temel sistemin bellek sayfası boyutunu döndürür.
