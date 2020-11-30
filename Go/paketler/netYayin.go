package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

//GetType ...
func GetType(a interface{}) interface{} {
	return reflect.TypeOf(a)

}

//AnaSayfa ...
func AnaSayfa(cevap http.ResponseWriter, istek *http.Request) {
	//cevap.Header().Set("Content-Type", "text/html")
	io.WriteString(cevap, "Merhaba Dünya")
	//log.Println(cevap)
	//log.Println(istek, "\n\n")
	//log.Println(istek.Context())
	//log.Println(istek.BasicAuth())
	//log.Println(istek.Body)
	log.Println(istek.Form, GetType(istek.Form))
	log.Println(istek.Method, GetType(istek.Method))
	log.Println(istek.Header["User-Agent"], "User-Agenti al kardeşim")

	log.Println(cevap.Header())

}

func main() {
	var port string = ":"
	if len(os.Args) == 1 {
		port += "8080"
	} else {
		port += os.Args[1]
	}
	log.Printf("%s Portunda , 127.0.0.1 adresinde başlatılıyor .\nURL : http://127.0.0.1%s\n\n\n", port, port)
	http.HandleFunc("/", AnaSayfa)
	http.ListenAndServe(port, nil)
	//http.ListenAndServe(":1010", nil)
	fmt.Println("Bekleniyor ..")
	fmt.Scanln()
}
