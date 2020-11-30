package main

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	html, err := ioutil.ReadFile("buton_template/index.html")
	if err != nil {
		panic("index.html yok şef")
	}
	ui, err := lorca.New("data:text/html,"+url.PathEscape(string(html)), "", 500, 200)
	defer ui.Close()

	ui.Bind("kontrol", func(text string) {
		fmt.Println(text)
		if text == "" {
			ui.Eval("alert('input boş kardeşim .');")
			return
		}
		ui.Eval(`$("p").hide(1000).show(1000).text("` + text + `")`)
	})

	<-ui.Done()
}
