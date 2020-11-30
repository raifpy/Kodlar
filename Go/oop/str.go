package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Buradaki amacım Python'daki str() olayını taklit etmeye çalışmak
/// Acemi bir kodlama olacak .

type str string

func (text str) split(key string) []string {
	return strings.Split(string(text), key)
}

func (text str) strip(key string) str {
	if key == " " {
		return str(strings.TrimSpace(string(text)))
	}
	return str(strings.Trim(string(text), key))

}
func (text str) center(index int) str {
	tmpInt := index - len(text)
	return str(strings.Repeat(" ", tmpInt/2) + string(text) + strings.Repeat(" ", tmpInt/2)) // tek sayı bile olsa çift sayı gibi olacak :D

}

func (text str) join(list []string) str {
	return str(strings.Join(list, string(text)))
}

func (text str) endswith(strl string) bool {
	return strings.HasSuffix(string(text), strl)
}

func (text str) startswith(strl string) bool {
	return strings.HasPrefix(string(text), strl)
}

func (text str) find(strl string) int {
	for index, eleman := range string(text) {
		if string(eleman) == strl {
			return index
		}
	}
	return -1
}

func (text str) format_map(strl map[string]string) str {
	text2 := string(text)
	for key, value := range strl {
		text2 = strings.Replace(text2, "{"+key+"}", value, -1)
	}
	return str(text2)
}

func (text str) format(strl ...string) str {
	text2 := string(text)
	for index, eleman := range strl {
		text2 = strings.Replace(text2, "{}", eleman, index+1)
	}
	return str(text2)
}

func (text str) replace(text1, text2 string) str {
	return str(strings.Replace(string(text), text1, text2, -1))
}

func (text str) upper() str {
	return str(strings.ToUpper(string(text)))
}

func (text str) lower() str {
	return str(strings.ToLower(string(text)))
}

func (text str) title() str {
	return str(strings.ToTitle(string(text)))

}

func (text str) isupper() bool {
	if string(text) == strings.ToUpper(string(text)) {
		return true
	}
	return false

}

func (text str) islower() bool {
	if string(text) == strings.ToLower(string(text)) {
		return true
	}
	return false
}

func (text str) istitle() bool {
	if string(text) == strings.ToTitle(string(text)) {
		return true
	}
	return false

}

func (text str) isdigit() bool {
	if _, err := strconv.Atoi(string(text)); err != nil {
		return false
	}
	return true
}

func (text str) String() string {
	return string(text)
}

func main() {
	abc := str("Merhaba Dünya")
	fmt.Println(abc)
	fmt.Println(abc.split(" "))
	fmt.Println(abc.split("a"))

	fmt.Println(strings.Repeat("-", 20))

	fmt.Println(abc.strip("a"))
	fmt.Println(abc.strip(" "))
	fmt.Println(abc.center(101))
	fmt.Println(str("Merhaba bu bir metin . Bakalım nasıl center yapacak").center(100))
	fmt.Println(str(".").join(strings.Split("Merhaba Dünya Nasılsın ? ?=", " ")))
	fmt.Println(abc.endswith("ab"))
	fmt.Println(abc.find("ü"))

	abcFormatMap := str("Merhaba {username} | {fname}\tNasılsın ?")
	fmt.Println(abcFormatMap.format_map(map[string]string{"username": "raif", "fname": "Ömer Raif Tekin"}))

	abcFormat := str("Merhaba {} | {}\tNasılsın ?")
	fmt.Println(abcFormat.format("raif", "Ömer Raif Tekin", "olmayanformat"))
	fmt.Println(abcFormat.format("kodware", "KodWare Team").center(100))

	fmt.Println(str("Merhaba Dünya").strip("e").split("a"))
	fmt.Println(str("Me{}rhaba Dünya").replace("e", " ").format("Go?").String())
	fmt.Println(str("Me{} Dünya").format(" \\ raif \\ ").String())
	digit := str("500")
	fmt.Println(digit.isdigit())

}
