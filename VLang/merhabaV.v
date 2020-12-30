import os
import time

//Ömer Raif Tekin | @raifpy
//Çrş 30 Ara 2020 00:26:46 +03

fn main(){
	println("Merhaba V!") // Golang -> fmt.Println()
	time.sleep(3) // Golang -> time.Sleep(time.Second * 3)
	os.exec("echo 'Merhaba Dünya!'") // Golang -> exec.Command()
	
}