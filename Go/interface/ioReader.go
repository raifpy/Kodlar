package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

//Ömer Raif Tekin | @raifpy
//Sal 29 Ara 2020 12:14:04 +03

/* Go programlarken neredeyse her zamann karşınıza çıkan bir yapı.
// io.Reader Read([]byte) adlı bir fonksiyon içeren her nesne io.Reader olabilir.
// Read([]byte) int,err | []byte alarak, int ve error return edecek.

bytes.NewReader
strings.NewReader
bufio.NewBuffer
os.File
http.Response.Body
...
*/

type golang string

func (g golang) Read(data []byte) (int, error) {
	lenght := len(data)
	for index, value := range []byte(g) {

		if index == lenght {
			return index, nil
		}
		data[index] = value

	}
	return len([]byte(g)), nil
}

func (g golang) alakasizFunc() {}

func main() {
	var dil golang = "Merhaba Golang"
	tmp := make([]byte, 10)
	yazilan, _ := dil.Read(tmp)
	fmt.Printf("%d kadar veri okundu\n", yazilan)
	fmt.Println(string(tmp))

	//dil dediğimiz eleman io.Reader olarak kullanılabilir;
	ioutil.ReadAll(dil) // Tepkime olmayacak

	var deneme io.Reader
	cevirildi := deneme.(golang) // deneme adlı io.Reader interface nesnesini golang adlı kendi oluşturduğum interface'e çevirdik
	cevirildi.alakasizFunc()
}

// Read([]byte) int,error olan her nesne io.Reader olabilir
// Close() error olan her nesne io.Closer olabilir
// Write([]byte) int,error olan her nesne io.Writer olabilir
// Yukarıdaki Read ve Close 'un ikisini birden içeren her nesne io.ReaderCloser olabilir [...]
// Yukarıdaki Read , Write ve Close'u içeren her nesne io.ReaderWriterCloser olabilir {os.File}
