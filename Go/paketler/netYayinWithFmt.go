package main

import (
	"fmt"
	"net/http"
	"time"
)

func Yayin(cevap http.ResponseWriter, istek *http.Request) {
	fmt.Fprintf(cevap, "Merhaba Dünya %s\n", time.Now())
}

func main() {
	http.HandleFunc("/", Yayin)
	http.ListenAndServe(":8080", nil)
}
