package main

import (
	"errHandler"
	"log"
	"os"
)

func main() {
	if _, err := os.Open("dosya"); os.IsNotExist(err) {
		log.Println("Not exists , creating")
		os.Create("dosya")
	}
	log.Println("Removing")
	err := os.Remove("dosya")
	errHandler.Handler(err)
}
