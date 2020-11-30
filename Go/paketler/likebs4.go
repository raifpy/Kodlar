package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/raifpy/Go/errHandler"
)

func main() {
	aranacak := strings.Join(os.Args[1:], " ")
	encode := url.QueryEscape(aranacak)
	fmt.Println(aranacak)
	search, err := http.Get("https://google.com/search?q=" + encode)
	errHandler.HandlerExit(err)
	/*fmt.Println(search)
	fmt.Println("\n\n")
	fmt.Println(search.Body)
	fmt.Println(search.ContentLength)
	fmt.Println(search.Cookies())
	fmt.Println("\n\n")
	fmt.Println(search.Header)
	fmt.Println(search.Status)
	fmt.Println(search.StatusCode)
	fmt.Println("\n\n\n\n")*/
	fmt.Println("StatusCode : ", search.StatusCode)

	//source, _ := ioutil.ReadAll(search.Body)
	doc, err := goquery.NewDocumentFromResponse(search)
	errHandler.HandlerExit(err)
	/*doc.Find("a").Each(func(i int, s *goquery.Selection) {
		key, bol := s.Attr("href")
		//fmt.Println(bol, key, s.Text())
		fmt.Println(bol, key, s.Text())

	})*/

	fmt.Println(doc.Find("h3").Text()) // Bu eleman birleştirip veriyor

}
