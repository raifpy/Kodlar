package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Alloc : ", bToMb(m.Alloc), " mb")
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\n\n", m.NumGC)

}

func bToMb(i uint64) uint64 {
	return i / 104 / 104
}
