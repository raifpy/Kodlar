package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func howManyTimesn(abc string) int {
	return strings.Count(abc, "\n")
}

func main() {
	fmt.Println("Göndermek için [enter] : ")
	input1, _ := bufio.NewReader(os.Stdin).ReadString('\n') // " yerine ' kullandık ! (byte eleman alması gerekiyordu)
	// ReadString('\n') demek "Enter" 'a basıldığı zaman okunacağı anlamına geliyor

	fmt.Println("Sonuç : ", input1)

	fmt.Println("Cümle içerisinde '.' (nokta) olmak zorunda : ")
	input2, _ := bufio.NewReader(os.Stdin).ReadString('.')
	fmt.Println("\033[31mSonuç : \033[0m\n", input2, howManyTimesn(input2), "Kez \\n geçiyor")
	// input içerisine 'eleman' elemanı var ise gönder demek oluyormuş :))
}
