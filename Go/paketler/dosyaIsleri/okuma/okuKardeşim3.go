package main

import (
	"bufio"
	"errHandler"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("dosya")

	errHandler.HandlerExit(err) // Bu errHandler bana özel bliyorsunuz .. | BetikSonu.org'da bulabilirsiniz

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
		fmt.Println("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Bunu sitenin birinden çaldım :')
// satır satır alıyor
// Ginede en mantıklısı ioutil.ReadFile() :))
