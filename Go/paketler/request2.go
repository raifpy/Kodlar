package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

//Handler ...
func Handler(err error) {
	if !(err == nil) {
		fmt.Println("Hata : ", err.Error())
		os.Exit(1)
	}
}
func main() {
	icerik, err := http.Get("http://127.0.0.1:8080/")
	Handler(err)
	html, err := ioutil.ReadAll(icerik.Body)
	Handler(err)
	fmt.Println("ICRK | TypeOf : ", reflect.TypeOf(icerik)) // *http.Response
	fmt.Println("HTML | TypeOf : ", reflect.TypeOf(html))   // Klasik []byte | C ile aynı :')
	fmt.Println(string(html))                               // String source code
}
