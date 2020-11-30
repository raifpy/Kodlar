package main

import "C"

func main() {
}

//export sayHi
func sayHi() *C.char {
	return C.CString("Go'dan Merhaba ... Bep bep bop biip 01010110")

}
