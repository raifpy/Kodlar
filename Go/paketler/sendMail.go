package main

import (
	"log"
	"net/smtp"
)

func main() {
	const port = "25"
	const server = "...betiksonu.org"
	const mail = "admin@betiksonu.org"
	const pass = ""
	var to = "...@gmail.com"
	msg := "From: " + mail + "\n" + "To: " + to + "\n" + "Subject: Merhaba Dünya ! Bu bir test maili\n"

	err := smtp.SendMail(server+":"+port, smtp.PlainAuth("", mail, pass, server), mail, []string{to}, []byte(msg))
	log.Println(err)
}
