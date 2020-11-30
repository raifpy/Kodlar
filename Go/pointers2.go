package main

import "fmt"

func main() {
	abc := 10
	yolla(&abc)
}

func yolla(sayi *int) {
	fmt.Println(*sayi)
}
