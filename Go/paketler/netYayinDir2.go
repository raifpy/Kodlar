package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("html")))
}

// php gibi davranıyor . Doğrudan var olan dosyalarıüzerinden çalıştırıyor ...
