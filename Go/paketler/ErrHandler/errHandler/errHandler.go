package errHandler

import (
	"log"
	"os"
)

func errHandler(err error) {
	if !(err == nil) {
		log.Println("Error : ", err.Error())
	}
}

func errHandlerExit(err error) {
	if !(err == nil) {
		log.Println("Error : ", err.Error(), "Exiting (1)")
		os.Exit(1)
	}
}
