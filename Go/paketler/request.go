package main

import (
	"fmt"
	"net/http"
	"os"
)

//Handler ...
func Handler(err error) {
	if !(err == nil) {
		fmt.Println("Hata : ", err.Error())
		os.Exit(1)
	}
}

func main() {
	içerik, err := http.Get("http://127.0.0.1:8080/olmayan?id=10")
	Handler(err)
	fmt.Println("Body   : ", içerik.Body)
	fmt.Println("Header : ", içerik.Header)
	fmt.Println("Statüs : ", içerik.Status)
	fmt.Println("StatüCo: ", içerik.StatusCode)
	fmt.Println("Cookies: ", içerik.Request.Cookies())
	fmt.Println("REQUEST: ", içerik.Request)
	fmt.Println("Req Url: ", içerik.Request.URL)
	fmt.Println("Req Mer: ", içerik.Request.Method)
	//fmt.Println("BasitAu: ", içerik.Request.BasicAuth())
	fmt.Println("Req UA : ", içerik.Request.UserAgent())
	fmt.Println(içerik.Request.Form)

}
