package main

import (
	"fmt"
	"os/exec"
)

func main() {
	exec_str := "uname"
	com := exec.Command(exec_str, "-a") // Boşluk koyacaksak burayada koymamız gerekiyor . subproject'deki gibi ..
	//fmt.Println(com)
	out, err := com.Output()
	if err != nil { //Hata var ise ;
		fmt.Println("Hata : ", err)
	} else {
		fmt.Println(string(out))
	}
}
