package main

import "fmt"

// panic() ile sisteme custom hatalar veriyoruz .
// Sistem panic hatası aldığı anda defer hariç bütün fonksiyonları kapatır .
// deferler çalışmaya devam eder

// İşte bu yüzden recover() 'ı defer ile kullanıyoruz . recover() (hatırla) sistemin panic alıp almadığını , aldıysa hatanın ne olduğunu veriyor ;

func main() {
	defer func() {
		if hata := recover(); hata != nil {
			fmt.Println("ErrorHandled  :  ", hata)
		} else {
			fmt.Println("Herhangi bir hata yakalanmadı !")
		}
	}()
	// bir değişkene atamak yerine doğrudan defer ile ve { .... }() ile çalıştırdık
	// Bu deyim programda hata olsa da olmasada , bütün fonkisyonlar bittiğinde çalışacak

	fmt.Println("Merhaba Dünya")
	panic("panikledim yav")
	fmt.Println("Proje sona erdi .")  //
	fmt.Println("Hoşça,kal Dünya ..") // Bu ve üstteki satırlar asla çalışmayacak . Go'da panic() döndüğü gibi defer harici tüm fonksiyonlar sona erdirilir .
}

// Yani hatayı recover() yakalıyor . Ama recover() 'ın çalışması için defer ile birlikte kullanılması şart
