package main

import "fmt"

type nesne string

func (n nesne) getNesne() string {
	return "heyt "
}

type eleman interface {
	getNesne() string
}

func getNesne(a eleman) { fmt.Println(a) }

func main() {
	var abc nesne = "Merhaba Dünya"
	getNesne(abc)

}
