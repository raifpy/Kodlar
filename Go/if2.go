package main

import "fmt"

func main() {
	// if'de en sondaki yapı kontrol edilir ve blog ona göre devam eder

	// if işlem1 ; işelm2 ; bool{ .... }

	var bol = true
	if abc := "Merhaba"; bol {
		fmt.Println("bol is true ve ", abc)
	}
	//fmt.Println(abc) // Sadece if blogu içinde tanımlanmış olacak ..
}
