package main

// Hata ayıklarken kullanılıyormuş ama hala ne işe yaradığını pek anlamadım
import "fmt"

func main() {
	run()
}
func run() {
	fmt.Println("Merhaba Dünya")
	panic("Olamaz , panikledim")
	abc := 10 * 10
	fmt.Println(abc)
}

/*
Merhaba Dünya
panic: Olamaz , panikledim

goroutine 1 [running]:
main.run()
        /home/raif/Masaüstü/Kodlar/Go/panic.go:10 +0x79
main.main()
        /home/raif/Masaüstü/Kodlar/Go/panic.go:6 +0x20
exit status 2
*/

// raise yapısı gibi kullanılabilir diye düşünüyorum

// panic hata yapısı imiş
