package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(max-min+1) + min)
}

func main() {
	randomInt(1, 50)
}
