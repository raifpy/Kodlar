package main

import "fmt"

func main() {
	buyukmu := isBigger(101, 100)
	if !buyukmu {
		fmt.Println("Olmadı usta ..")
	} else {
		fmt.Println("Heyt usta ")
	}

}
func isBigger(x, y int) bool {
	if x > y {
		return true
	} else {
		return false
	}
}

func whichBig(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
