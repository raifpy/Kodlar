package main

import (
	"fmt"
	"runtime"
	"time"
)

func gc() {
	runtime.GC()
}

func getUsage() {
	var abc runtime.MemStats
	runtime.ReadMemStats(&abc)
	fmt.Println(time.Now().String()[11:22], abc.Alloc/1024/1024, "mb")
}

func main() {
	getUsage()
	slice := make([]string, 9999999)
	fmt.Println(cap(slice))
	getUsage()
	fmt.Println(cap(slice))
	inta := 0
	for ; inta < 100; inta++ {
		slice = append(slice, "10")
	}
	getUsage()
	fmt.Println(cap(slice))
	gc()
	getUsage()
	fmt.Println(len(slice))
}
