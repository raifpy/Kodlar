package main

import (
	"encoding/json"
	"errHandler"
	"fmt"
)

func main() {
	// map to json string

	var abc = map[string]string{}
	abc["Dünya"] = "Merhaba Dünya"
	abc["Merhaba"] = "Dünya Merhaba"

	var abc_json, err = json.Marshal(abc) // abc_json []byte
	errHandler.HandlerExit(err)

	abc_json_string := string(abc_json) // Bu bizim json dosyası
	fmt.Println(abc_json_string)
	fmt.Println("---------")
	var cba map[string]string
	errHandler.HandlerExit(json.Unmarshal(abc_json, &cba))
	fmt.Println(cba)
	for key, eleman := range cba {
		fmt.Println("Key : ", key, "\tValue : ", eleman)
	}

}
