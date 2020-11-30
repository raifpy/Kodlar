package main

import (
	"io"
	"os"
)

const myStr = "Hello GO"

func main() {
	say(myStr)
}
func say(text string) {
	io.WriteString(os.Stdout, text)
}
