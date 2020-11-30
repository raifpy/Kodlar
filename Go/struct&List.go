package main

import "fmt"

type dagitimlar struct {
	dagitim, pkm, taban string
}

func main() {
	abcd := make([]dagitimlar, 3) // Böyle de yapabilirdik
	//abcd := [3]dagitimlar{}
	abcd[0] = dagitimlar{dagitim: "Debian", pkm: "apt", taban: "Debian"}
	abcd[1] = dagitimlar{"Ubuntu", "apt", "Debian"}
	abcd[2] = dagitimlar{"Arch", "pacman", "Arch"}
	//abcd[3] = dagitimlar{"PiSi", "pisi", "PiSi"}

	fmt.Println(abcd)
	for _, dagitim := range abcd {
		fmt.Println(dagitim.dagitim)
	}
	fmt.Println(cap(abcd))
	abcd = append(abcd, dagitimlar{"PiSi", "pisi", "PiSi"})
	fmt.Println(cap(abcd))
	println("\t---------")
	for _, eleman := range abcd {
		fmt.Println(eleman)
	}

}
