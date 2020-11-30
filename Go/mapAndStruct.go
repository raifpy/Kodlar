package main

import "fmt"

type nesne struct {
	araba string
	agac  string
	world int
}

func main() {
	__map__ := map[string]nesne{}
	__map__["go"] = nesne{"Toyota", "Çam", 2020}

	fmt.Println(__map__)
	fmt.Println(__map__["go"])
	fmt.Println(__map__["go"].world)

	__map__["py"] = nesne{}

	fmt.Println(__map__["py"])
	//__map__["py"].araba = "ARABABV" // Sonradan eklemeye yapamıyormuşuz :D
	//__map__["py"].world = 202020
	//fmt.Println(__map__["py"])

	var konya = nesne{}
	konya.agac = "Konya"
	konya.world = 10101010
	konya.araba = "VİYUUUUUUUUUU"

	__map__["_/\\_"] = konya

	fmt.Println(__map__["_/\\_"])

	// Böylede olabilir ...

	MAP2 := map[string]nesne{
		"tlc": nesne{"asd", "dsa", 1},
		"atv": nesne{"ÇÖP", "DİZİLER", 0},
	}
	fmt.Println(MAP2["tlc"])
	fmt.Println(MAP2["atv"])
}
