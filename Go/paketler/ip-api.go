package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiurl := "http://ip-api.com/json" // http olacak
	ham, _ := http.Get(apiurl)
	//errHandler.HandlerExit(err)
	ham2, _ := ioutil.ReadAll(ham.Body)
	//errHandler.HandlerExit(err)
	html := string(ham2)
	//errHandler.HandlerExit(err)
	fmt.Println(html)

	var MAP = map[string]string{} // key ve value kısmına string alan map oluşturduk
	json.Unmarshal(ham2, &MAP)
	fmt.Println("\n----------------------------------\n")
	fmt.Println(MAP)

	if deger := MAP["status"]; deger != "success" {
		fmt.Print("\033[31m", MAP["message"], "\033[0m\n")
		return
	}

	iss := MAP["isp"]
	country := MAP["country"]
	ip := MAP["query"]
	fmt.Printf("IP\t=\t%s\nISS\t=\t%s\nCOU\t=\t%s\n", ip, iss, country)

	//

}
