package main

import (
	"fmt"
	"reflect"
)

func main() {
	verBanaBirşey([]string{"a", "b", "c"})
}

func verBanaBirşey(abc interface{}) {
	switch tip := abc.(type) {
	default:
		fmt.Println(tip, reflect.TypeOf(tip)) // ne işe yaradı ilginç

	}
}
