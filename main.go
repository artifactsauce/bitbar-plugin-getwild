//usr/bin/env go run $0 $@; exit

package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"net/url"
)

type Config struct {
	ListItemNumber int
	SearchPhrase   string
}

type Provider struct {
	Name       string
	BaseUrl    string
	SearchPath string
}

type Video struct {
	Title string
	Path  string
	Id    string
}

func GetWildAndTough(p Provider, c Config) {
	doc, err := goquery.NewDocument(GetSearchUrl(p, c))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("h3.yt-lockup-title").Each(func(i int, s *goquery.Selection) {
		if i > c.ListItemNumber {
			return
		}
		v := Video{}
		v.Title = s.Find("a").Text()
		v.Path, _ = s.Find("a").Attr("href")
		Url := p.BaseUrl + v.Path
		fmt.Printf("%s | href=%s\n", v.Title, Url)
	})
}

func GetSearchUrl(p Provider, c Config) string {
	return p.BaseUrl + p.SearchPath + "?search_query=" + url.QueryEscape(c.SearchPhrase)
}

func GetConfig() Config {
	c := Config{}
	c.ListItemNumber = 5
	c.SearchPhrase = "Get Wild"
	return c
}

func GetProvider() Provider {
	p := Provider{}
	p.Name = "YouTube"
	p.BaseUrl = "https://www.youtube.com"
	p.SearchPath = "/results"
	return p
}

func main() {
	p := GetProvider()
	c := GetConfig()

	fmt.Println(":gun:")
	fmt.Println("---")
	GetWildAndTough(p, c)
	fmt.Println("---")
	fmt.Println("Refresh | refresh=true color=#C0C0C0")
}
