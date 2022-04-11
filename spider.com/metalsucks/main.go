package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://studygolang.com/articles/17638")
	if err != nil {
		log.Fatal("error: ", err)
		return
	}
	// fmt.Println(res)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code {%d} error: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("error: ", err)
	}
	doc.Find("#myeditor").Each(func(i int, s *goquery.Selection) {
		text := s.Find("p").Text()
		text2 := s.Find("ol").Text()
		fmt.Println("Review", text)
		fmt.Println("Review", text2)
	})
}
