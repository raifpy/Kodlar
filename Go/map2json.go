package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var abc = make(map[string]string)
	abc["Merhaba"] = "Dünya"
	fmt.Println(abc)
	abc_json, _ := json.MarshalIndent(abc, "", "")
	fmt.Println(abc_json) // Bu olmadı neden bilmiyorum :D
}
