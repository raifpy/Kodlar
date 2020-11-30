package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Handler böyle yapmak zorundayaım lint uyarı veriyor , sinir ediyor
func Handler(err error) {
	if !(err == nil) {
		fmt.Println("Olamaz Hata : ", err.Error())
		os.Exit(1)
	}
}

func main() {
	dosya := "dosya"
	icerik, err := ioutil.ReadFile(dosya)
	Handler(err)

	fmt.Println(string(icerik))

}
