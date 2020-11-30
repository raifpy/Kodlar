// goruntime ögeleri ( eş zamanlı çalışan elemanlar `go dewam()`) birbirleri ile iletişime geçmek için kanalları ( chhan ) kullanır . Bu da bunun ile ilgili

package main

import "fmt"

func main() {
	var abc chan string = make(chan string)

	// yada

	abcd := make(chan string)

	fmt.Println(abc, abcd) // ikisde pointers dönecek
}
