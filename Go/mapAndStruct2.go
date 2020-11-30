package main

import "fmt"

type nesne struct {
	list []string
}

func main() {
	abc := nesne{[]string{"hey", "hey1"}}
	fmt.Println(abc)

	map2 := map[string]nesne{}
	strc := nesne{}
	strc.list = []string{"Ouuiw", "Ouiww", "Hey", "Deneme"}
	map2["Deneme"] = strc
	map2["a"] = strc
	map2["b"] = strc

	fmt.Println(map2["Deneme"])
	delete(map2, "b")
	abc, bol := map2["b"] // delete() ile sildik
	fmt.Println(abc, bol) // false çıktı :))
}
