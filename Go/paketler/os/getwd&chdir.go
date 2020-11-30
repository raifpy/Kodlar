package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Getwd()) // location | error
	os.Chdir("testDir")     // error
	log.Println(os.Getwd()) // location | error

}
