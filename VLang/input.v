// Ömer Raif Tekin | @raifpy
// Çrş 30 Ara 2020 03:05:32 +03
import os

fn main(){
	start:
	print("Merhaba! Sen Kimsin: ")
	text := os.input("")
	if text.trim(" ") == ""{
		goto start
	}
	println("Merhaba ${text}. Hoşgeldin!")
	
}