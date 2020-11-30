package main

import (
	"fmt"
	"os/user"
)

func main() {
	user1, _ := user.Lookup("olmayanUser") // user, error

	fmt.Println(user1) // user.Current ile aynı elemanları dönüyor

	user.LookupId("id") // Aynı işler
}
