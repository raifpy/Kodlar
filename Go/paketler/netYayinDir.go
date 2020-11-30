package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.Handle("/hu", http.FileServer(http.Dir("html")))
	http.ListenAndServe(":8080", nil)

}

// Belirtilen isimde /devam / gibi , dosya üzerinden yayın yapıyor
