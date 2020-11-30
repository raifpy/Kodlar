package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	username, _ := os.Hostname()
	fmt.Println(username)
	group, err := user.LookupGroup(username)
	errHandler.HandlerExit(err)
	fmt.Println(group.Gid)  // 1000 | root olsaydı 0 {os.geteuid()}
	fmt.Println(group.Name) // raif {os.Hostname()}

	groupWithId, err := user.LookupGroupId("10")
	errHandler.HandlerExit(err)
	fmt.Println(groupWithId)
}

// root olsam dahi Hostname() raif çıkıyor ...
