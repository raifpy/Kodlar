package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	log.Println("interrupt bekleniyor ...")
	işlem := <-c
	log.Println("\nOlamaz dostum : ", işlem)

}

// Bu eleman ile ctrl-c 'yi yakalayıp , basıldığında işlem yapabiliriz
