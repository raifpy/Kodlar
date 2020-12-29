package main

import (
	"errors"
	"fmt"
	"strings"
)

//Ömer Raif Tekin | @raifpy
//Sal 29 Ara 2020 11:57:59 +03

//Interface ile ilgili birkaç örnek. Oldukça yararlı yapılar..

type muz interface {
	Error() string
}

func verBanaMuz(m muz) {
	if m != nil {
		fmt.Println("Olamaz muz çürük! ", m.Error())
	}
}

type text string

func (a text) Error() string {
	tmp := string(a)
	if strings.ToLower(tmp) == "codeksiyon" {
		return "Hayır :)"
	}
	return tmp
}

func main() {
	var err error = nil
	verBanaMuz(err) // Hiçbir şey
	err = errors.New("codeksiyon'a hala katılmamışsın")
	verBanaMuz(err) // Olamaz muz çürük!  codeksiyon'a hala katılmamışsın

	var codeksiyon text = "Codeksiyon"
	verBanaMuz(codeksiyon) // text adlı type'in Error() adında bir fonksiyonu olduğu ve bu fonksiyon String döndürdüğü için kullanabiliriz

}

// error nesnesi aslında bir interface'dir. Error fonksiyonu olup string döndüren her nesne error olarak kullanılabilir.
