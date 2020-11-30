package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><meta charset="utf8"><title>Merhaba Dünya</title></head>
		<body><h2>Merhaba Codeksiyon! ü</h2></body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
