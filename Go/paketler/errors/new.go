package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	customError := errors.New("Olamaz dostum bu dir hata ...")
	fmt.Println(reflect.TypeOf(customError))
	fmt.Println(customError.Error())
}
