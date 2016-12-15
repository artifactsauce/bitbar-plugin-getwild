//usr/bin/env go run $0 $@; exit

package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"net/url"
)

var youtubeBaseUrl = "https://www.youtube.com"
var youtubeSearchPath = "/results"

type Video struct {
	Title string
	Path  string
	Id    string
}

func GetWildAndTough(searchUrl string) {
	doc, err := goquery.NewDocument(searchUrl)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("h3.yt-lockup-title").Each(func(i int, s *goquery.Selection) {
		if i > 5 { return }
		v := Video{}
		v.Title = s.Find("a").Text()
		v.Path, _ = s.Find("a").Attr("href")
		Url := youtubeBaseUrl + v.Path
		fmt.Printf("%s | href=%s\n", v.Title, Url)
	})
}

func GetSearchUrl(search_phrase string) string {
	return youtubeBaseUrl + youtubeSearchPath + "?search_query=" + url.QueryEscape(search_phrase)
}

func main() {
	search_phrase := "get wild"

	fmt.Println("Get Wild")
	fmt.Println("---")
	GetWildAndTough(GetSearchUrl(search_phrase))
	fmt.Println("---")
	fmt.Println("Refresh | refresh=true")
}
