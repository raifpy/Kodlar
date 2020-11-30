package main

import (
	"fmt"
	"strings"
)

func main() {
	var abc = "Merhaba Dünya"
	fmt.Println("Orijinal      : ", abc)

	var abcUpper = strings.ToUpper(abc) // str'ı tamamen büyük yapıyor
	fmt.Println("Upper         : ", abcUpper)

	var abcLower = strings.ToLower(abc) // str'i tamamen küçük yapıyıor
	fmt.Println("Lower         : ", abcLower)

	var abcTitle = strings.Title(abcLower) // kelimlerin baş harflerini büyütüyor
	fmt.Println("Title (lower) : ", abcTitle)

	var abcToTitle = strings.ToTitle(abc) // Hala anlayamadım
	fmt.Println("ToTitle       : ", abcToTitle)

	var abcSplit = strings.Split(abc, "") // Klasik split
	fmt.Println("Split ''      : ", abcSplit)

	var abcJoin = strings.Join(abcSplit, "-") // klasik join
	fmt.Println("SplitJoin -   : ", abcJoin)

	var abcSplitAfter = strings.SplitAfter(abcJoin, "-") // Çok güzel bir arkadaş
	fmt.Println("SplitAfter '' : ", abcSplitAfter)

	var abcReplace = strings.Replace(abc, "a", "c", -1) // orjinal | değişecek str | yerine geçecek str | kaç değişimde durayım |||| Biçeminde olması gerek
	fmt.Println("Replace a - c : ", abcReplace)

	var abcContainsAny = strings.ContainsAny(abc, "Merhaba") // str içinde str aramak için kullanılıyor | contains = içermek
	fmt.Println("ContainsAny   : ", abcContainsAny, " Aranan kelime : Merhaba")

	var abcCount = strings.Count(abc, "a") // str içerindeki , str sözcüğün kaç kez geçtiğini verir
	fmt.Println("Count - a     : ", abcCount)

	var abcHasPrefix = strings.HasPrefix(abc, "M") // str'in başında (ilk harfinde) arama yapıyor
	fmt.Println("HasPrefix - M : ", abcHasPrefix)

	var abcHasSuffix = strings.HasSuffix(abc, "Dünya") // str'in sonunda arama
	fmt.Println("HasSuffix Dnya: ", abcHasSuffix)

	var abcIndex = strings.Index(abc, "ü") // str'i arıyor bulduğu ilk indexi döner . Eğer bulamazsa -1 döner
	fmt.Println("Index - ü     : ", abcIndex)

	var abcIndexAny = strings.IndexAny(abc, "nyü") // nyü str'ini tek tek arar n , y, ü ilk bulduğunu döner
	fmt.Println("IndexAny -nyü : ", abcIndexAny)

	var abcRepeat = strings.Repeat(abc, 10)
	fmt.Println("Repeat 10     : ", abcRepeat)

}
