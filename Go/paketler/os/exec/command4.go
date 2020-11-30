package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("date").String() //    /bin/date
	log.Println(cmd)
}
