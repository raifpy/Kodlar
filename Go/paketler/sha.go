package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

const text = "Merhaba Dünya"

func main() {
	fmt.Println("Sha256 : ", sha256.Sum256([]byte(text)))
	fmt.Println("Sha1   : ", sha1.Sum([]byte(text)))
	fmt.Println("md5    : ", md5.Sum([]byte(text)))

	//for file;

	hash := sha256.New()
	file, _ := os.Open("sha.go")
	io.Copy(hash, file)

	fileSha256 := hash.Sum(nil)
	fmt.Println(string(fileSha256), "\n\n")

	read := strings.NewReader("merhaba")
	hash = sha256.New()

	io.Copy(hash, read)
	hash.Write([]byte("muzlusut"))
	fmt.Println(string(hash.Sum(nil)))

}
