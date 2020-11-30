package main

import (
	"errHandler"
	"os"
)

func main() {
	var metin = "Pzt Tem 13 17:15:56 +03 2020 | @raifpy\n"
	file, err := os.Create("dosya")
	errHandler.HandlerExit(err)

	metin_byte := []byte(metin)
	_, err = file.Write(metin_byte)
	errHandler.HandlerExit(err)
	file.Write([]byte("Bu da 2. satır :¿\n"))
}

// os.Create() ile dosya açıldığında dosya baştan yazılır !
