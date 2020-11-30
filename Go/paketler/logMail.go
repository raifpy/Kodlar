package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sss, _ := syslog.New(syslog.LOG_NEWS, "Codeksiyon Log")
	log.SetOutput(sss)
	log.SetPrefix("Codeksiyon-Log")
	log.Println("Log with log")
	fmt.Fprintln(sss, "Log With fmt f")

	sss.Warning("Warning with orgi")

}
