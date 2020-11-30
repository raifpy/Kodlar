// atanmış veri türünü bulmak için kullanacağımız kütüphane "reflect"
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var abc = 10
	var abc2 = "merhaba Dünya"
	var abc3 = 10.9
	var abc4 = true

	fmt.Println(abc, abc2, abc3, abc4)

	// abc'nin türünü bulalım :

	fmt.Println("abc 'nin türü : ", reflect.TypeOf(abc))

	// abc2 'nin türünü bulalım :

	fmt.Println("abc2 'nın türü : ", reflect.TypeOf(abc2))

	// abc3 'üm türünü bulalım :) | Bu iş şimdiden hoşuma gidiyor

	fmt.Println("abc3 'nın türü : ", reflect.TypeOf(abc3))

	// 4'ü de bulduk mu tamamdır

	fmt.Println("abc4 'nın türü : ", reflect.TypeOf(abc4))

	// En son da reflect.TypeOf() 'ın çıktısının hangi türden olduğunu , gine reflect.TypeOf() kullanarak bulalım ..

	fmt.Println("abc 'nin reflect çıktısının türü : ", reflect.TypeOf(reflect.TypeOf(abc))) // int değeri verim , bakalım ne çıkacak

	// *reflect.rtype || Çıktı bu oldu :D . string çıkmasını beklerdim hayal kırıklığına uğramadım değil :)

	reflect_oge := reflect.TypeOf(abc)       // Uyarı veriyor , nedenini bilmiyorum :))
	reflect_ama_string := reflect_oge.String // Sonuç String'e dönüşmüş oluyor böylece :)
	//fmt.Println("+ ile String olduğunu doğruluyorum :D " + reflect_ama_string) // Hüü kabul etmedi :D
	fmt.Println(reflect.TypeOf(reflect_ama_string)) // func() string

	// :') Yapamadım :D

}
