package main

import (
	"fmt"
	"os/user"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	// returns current user
	user, err := user.Current()
	errHandler.HandlerExit(err)

	fmt.Println(user.Gid)
	fmt.Println(user.GroupIds()) //[]strings | err
	fmt.Println(user.HomeDir)
	fmt.Println(user.Name)
	fmt.Println(user.Uid)
	fmt.Println(user.Username)
}

// os.Hostname() kullanmak yerine user.Current() kullanmak çok daha mantıklı .
