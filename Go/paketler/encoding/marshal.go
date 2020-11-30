package main

import (
	"encoding/json"
	"fmt"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	jsonJ := map[string]string{
		"ok": "ok",
	}
	jsonStr, err := json.Marshal(jsonJ)
	errHandler.HandlerExit(err)

	fmt.Println(string(jsonStr))

	// Marshal , interface ögesini json biçimine çevirir..

}
