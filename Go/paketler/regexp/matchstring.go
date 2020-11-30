package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regStr = "p[a-z]+ch"                          // [a-z] tüm alfabeyi , + ise devamlı olabilceğini
	fmt.Println(regexp.MatchString(regStr, "puch"))   // true , nil
	fmt.Println(regexp.MatchString(regStr, "pach"))   // true , nil
	fmt.Println(regexp.MatchString(regStr, "puach"))  // true , nil
	fmt.Println(regexp.MatchString(regStr, "puuach")) // true , nil
	fmt.Println(regexp.MatchString(regStr, "p**ch"))  // false , nil | yıldız alfabeden değil :))
	fmt.Println(regexp.MatchString(regStr, "pch"))    // false , nil

	// bool , error
	//MatchString 'in olayı bu , düzenli ifadeyi ver , normal string ver. Bu buna olabilir mi bi bak
}
