package main

import (
	"fmt"
	"log"
	"time"

	"github.com/raifpy/Go/errHandler"

	"github.com/raifpy/Go/raiFile"
)

// Modül olmuş mu diye test ediyorum 0_0

func main() {
	//raifile.writeFile("dosya", "Merhaba Dünya")
	icerik, err := raiFile.ReadFile("log")
	errHandler.HandlerExit(err)
	log.Println("İÇERİK : \n\n", icerik)
	raiFile.WriteFileLines("log", "Son Güncelleme : "+fmt.Sprint(time.Now()))
	log.Println("Son değişiklik en altına yazıldı !")

}

//Hojdur
