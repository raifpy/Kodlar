package main

import (
	"fmt"
	"os"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	group, err := os.Getgroups()
	errHandler.HandlerExit(err)
	fmt.Println(group)

}

// Ne işe yaradıklarını anlamış değilim
// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
// Getgroups arayan kişiye ait olduğu grupların sayısal kimliklerinin bir listesini verir.
