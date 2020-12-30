
//Ömer Raif Tekin | @raifpy
//Çrş 30 Ara 2020 01:28:12 +03

fn main(){
	abc := "Bellekte yer kaplayan ama değişemeyen değişken ${"immutability"}"
	// abc = "Değişti ki" // Üzgünüm.
	//Ayrıca Go'daki gibi bu değişkeni kullanmak zorundayız;
	println(abc)

	mut abc2 := "Bellekte yer kaplayan ve değiştirilebilen değişken | mut xx:=xx"
	abc2 = "Evet."
	println(abc2)

	// V'de mut <tip> := <deger> ile tanımlanmayan her değişken immutability; Sonradan değiştirilemiyor.

	//int varsayılan olarak 32bit
	sayi := 10
	println("Sayi: "+sayi.str())  // <integer>.str() -> string

	sayi8 := i8(10) // 8bitlik int
	println("Sayi8: "+sayi8.str()) 

	liste := ["Merhaba","","","Dünya",""]
	//liste := []string{"Merhaba","","","Dünya",""} // Yanlış kullanım
	println(liste.join(" - "))
	println(liste[0..4])
	println(liste[0..4].join("0"))
	//print(liste.sort_by_len()) // Çalışabilmesi mut ile tanımlanması gerekiyor;

	//mut listeMut := "Merhaba Dünya ! , * -".split(" ") // Ayrıca değişkenler büyük harf içeremez?
	mut liste_mut := "Merhaba Dünya ! , * -".split(" ")

	println("liste_mut: "+liste_mut.str())
	println(liste_mut.eq(liste)) // false ~ Karşılaştırma
	println("")
	for value in liste_mut{
		println(value)
	}
	
	// liste << "Eleman" // immiutable
	liste_mut << "Eleman" 
	//liste_mut << false // liste dynamic değil

	println("")
	for _,value in liste_mut{ // _ 'nın hiç bir anlamı yok
		println(value) 
	}

	println(public_function())
	println(private_function())
	
}

pub fn public_function() string{ // "değişken adı küçük harf olmalı" kuralı geçerli.
	return "public return"
}// public tanımladığım için derleyici uyarı gösterdi.

fn private_function() string{
	return "private return"
}