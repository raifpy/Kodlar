package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/raifpy/Go/errHandler"
)

//Web ...
func Web(cevap http.ResponseWriter, istek *http.Request) {
	io.WriteString(cevap, "<html><head><title>Merhaba Dünya</title></head>Merhaba Dost !</hmt>")
}

func main() {
	fmt.Println("Başlatılıyor !")
	http.HandleFunc("/world", Web)
	//http.HandleFunc("/favicon.ico", Web)
	//http.HandleFunc("/gorsel.png", Web)
	//errHandler.HandlerAndExit(http.ListenAndServe(":8080", http.FileServer(http.Dir("/"))))
	errHandler.HandlerBool(http.ErrAbortHandler)
}
