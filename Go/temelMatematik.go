package main
import "fmt"
func main(){
	// Kurallar her dilde olduğu gibi işliyor ..
	sayi1 := 10
	sayi2 := 20
	var sayi3 int
	sayi3 = sayi1 * sayi2
	fmt.Println(sayi3)
	sayi3 *= 2
	fmt.Println(sayi3)
	sayi3 += 1
	fmt.Println(sayi3)
	sayi3++
	fmt.Println(sayi3)
	//++sayi3
	//fmt.Println(sayi3)

	// java' ile tamamen aynı ** ++sayi3 yapısı dışında | C syntaxı . (Matematik işlemleri aynı)

}
