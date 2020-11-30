// Bu hata bu mu diyor özetle

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	myError := errors.New("eşek bilmediği otu yerse karnı ağrırmış")

	fmt.Println(errors.Is(myError, os.ErrExist)) // false
}
