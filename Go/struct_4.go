package main

import "fmt"

// miras alma işlemi yapacağız

type canlilar struct {
	memeliler   string
	sürüngenler string
	böcükler    string
}

type böcük struct {
	canlilar // canlilar dememiz yeterli oluyor | ilk başa    || miras almış olduk
	tür      string
	oyun     string
}

func main() {
	böcek := böcük{}
	//böcek.asdadas = "Tım"// Olmayan eleman olduğu için çöktü
	böcek.memeliler = "Buraya erişimim var"
	böcek.oyun = "Keza buraya da erişimim var"
	fmt.Println(böcek)
	// Miras almış olduk :)

}
