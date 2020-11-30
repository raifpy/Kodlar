package main

import "C"

func main() {}

//export retutf8
func retutf8() *C.char {
	return C.CString("Mehaba Dünya !")
}

//export retascii
func retascii() *C.char {
	return C.CString("Merhaba Dunya")
}
