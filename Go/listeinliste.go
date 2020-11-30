package main

import "fmt"

func main() {
	abc := [2][1]string{{"Deneme"}, {"Hey"}}
	fmt.Println(abc)

	abcd := [2][1]string{}

	abcd[0][0] = "Merhaba"
	abcd[1][0] = "Dünya"

	fmt.Println(abcd)

}
