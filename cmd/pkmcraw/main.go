package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Volume struct {
	ChapLink []Chapter
}

type Chapter struct {
	Title    string
	PageLink []string
}

var (
	PathPkm     = "Pokemon/"
	DomainPkm   = "www.pokemonspecial.com"
	LinkPkmCraw = "https://www.pokemonspecial.com/2013/12/chapter-001.html"
	BaseURL     = "https://www.pokemonspecial.com/2022/10/pokemon-dac-biet.html"
)

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains(DomainPkm),
)

func readHTML() map[string][]string {

	var title string
	var pages []string
	mapChap := make(map[string][]string)
	checkEnd := false

	BaseURL := "https://www.pokemonspecial.com/2014/06/chapter-003.html"
	nextVolume := LinkPkmCraw

	for {
		collectorPKM.OnHTML("h3.post-title", func(element *colly.HTMLElement) {
			title = strings.Trim(element.Text, "\n")
		})

		// collectorPKM.OnHTML(".separator a", func(element *colly.HTMLElement) {
		// 	link := element.Attr("href")
		// 	pages = append(pages, link)
		// 	mapChap[title] = pages
		// })

		collectorPKM.OnHTML(".post-body div a", func(element *colly.HTMLElement) {
			link := element.Attr("href")
			pages = append(pages, link)
			mapChap[title] = pages
		})

		collectorPKM.OnHTML("a#Blog1_blog-pager-newer-link", func(element *colly.HTMLElement) {
			nextVolume = element.Attr("href")
			if nextVolume == BaseURL {
				checkEnd = true
			}
		})

		if err := collectorPKM.Visit(nextVolume); err != nil {
			log.Fatal(err)
		}

		if checkEnd {
			break
		}
	}

	return mapChap
}

func main() {
	mapChap := readHTML()

	for k, v := range mapChap {
		if err := os.MkdirAll(PathPkm+k, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		fmt.Println(k)
		fmt.Println(len(v))
		// for _, x := range v {
		// if err := crawpkm.DownloadFile(x, "", PathPkm+k); err != nil {
		// 	log.Fatal(err)
		// }
		// }
	}
}
