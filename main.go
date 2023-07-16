package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://yahoo.co.jp"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("faild request: %s", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("faild body: %s", err)
	}

	doc.Find("#tabpanelTopics1 div li").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		val, exists := a.Attr("href")
		if exists {
			fmt.Printf("TITLE: %s, URL: %s\n", s.Text(), val)
		}
	})
}
