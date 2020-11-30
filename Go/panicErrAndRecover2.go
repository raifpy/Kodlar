package main

import (
	"fmt"
	"os"
)

func getPanic(text string) {
	panic(text)
}
func run(panic bool) {
	defer func() {
		if hata := recover(); hata != nil {
			fmt.Println("\033[31mHata yagaladum\033[0m ", hata)
		}
	}()

	fmt.Println("Merhaba Dünya")
	if panic {
		getPanic("panic bool değeri true , bu bile panik yapmama yeter :(")
	} else {
		fmt.Println("\033[32mDone .\033[0m")
	}
}

func RunWithArgs() {
	args := os.Args
	if len(args) < 2 {
		run(true)
	} else if args[1] == "false" {
		run(false)
	} else {
		run(true)
	}
}

func main() {
	RunWithArgs()
}
