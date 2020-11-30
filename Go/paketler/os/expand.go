package main

import (
	"fmt"
	"os"
)

func mapper(name string) string {
	switch name {
	case "name":
		return "Ömer Raif"

	case "username":
		return "@raifpy"

	default:
		return "Default döndü !"
	}
}

func main() {
	fmt.Println(os.Expand("Merhaba $name \tUserName : ${username}", mapper))
	fmt.Println(os.Expand("Merhaba $name \tUserName : ${yok}", mapper))

	// Python'daki .format yapısına benziyor

	// $name olduğunda switch'den name'i çağırıyor
	fmt.Println(os.ExpandEnv("Merhaba Dostum , senin shell =  $SHELL"))
	fmt.Println(os.ExpandEnv("Şuanda olmayan bir env       =  $olmayan"))
	os.Setenv("olmayan", "Artık var")
	fmt.Println(os.ExpandEnv("Şuanda olmayan bir env       =  $olmayan"))

}

// Env'den veri çekmeye yarıyor , güzel bir yapı
