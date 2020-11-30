package raiFile

import (
	"io/ioutil"
	"os"
)

//ReadFile dosyayı okur string döner
func ReadFile(fileName string) (string, error) { // CleanCode kurallarına uyalım dost
	icerik, err := ioutil.ReadFile(fileName)
	if !(err == nil) {
		return "", err
	}
	return string(icerik), nil
}

//WriteFile Dosyaya yazar ama sıfırlayarak .
func WriteFile(fileName string, text string) error {
	dosya, err := os.Create(fileName)
	defer dosya.Close()
	if !(err == nil) {
		return err
	}
	_, err = dosya.WriteString(text)
	if !(err == nil) {
		return err
	}
	return nil
}

//WriteFileLines Dosyaya yazıyor ama sıfırlamadan .
func WriteFileLines(fileName, text string) error {
	icerik, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm) // dosya adı - string , flag - int , perm - os.FileMode
	defer icerik.Close()
	if !(err == nil) {
		return err
	}
	_, err = icerik.WriteString(text + "\n") // Alt satıra da biz indirelim
	if !(err == nil) {
		return err
	}
	return nil
}
