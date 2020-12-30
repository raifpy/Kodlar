// Ömer Raif Tekin | @raifpy
// Çrş 30 Ara 2020 02:44:47 +03

import net.http


fn main(){
	
	response := http.get("http5://ip-api.com/json") or {
		println("Hata alındı: "+err.str()) 
		return
	}
	println("StatusCode: "+response.status_code.str())
	println(response.text)

}