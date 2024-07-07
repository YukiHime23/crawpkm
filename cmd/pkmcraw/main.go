package main

import (
	"fmt"
	"log"
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

func main() {
	collectorPKM := colly.NewCollector(
		colly.AllowedDomains(DomainPkm),
	)

	var chap Chapter

	mapChap := make(map[string][]string)
	checkEnd := false

	BaseURL := "https://www.pokemonspecial.com/2014/06/chapter-003.html"
	nextVolume := LinkPkmCraw

	for {
		collectorPKM.OnHTML("h3.post-title.entry-title", func(e *colly.HTMLElement) {
			chap.Title = strings.Trim(e.Text, "\n")
		})

		collectorPKM.OnHTML(".separator a", func(element *colly.HTMLElement) {
			link := element.Attr("href")
			chap.PageLink = append(chap.PageLink, link)

			mapChap = map[string][]string{
				chap.Title: chap.PageLink,
			}
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

	for k, v := range mapChap {
		// if err := os.MkdirAll(PathPkm+k, os.ModePerm); err != nil {
		// 	log.Fatal(err)
		// }

		// for _, x := range v {
		// 	if err := crawpkm.DownloadFile(x, "", PathPkm+k); err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Println("craw -> " + k + " <- done!")
		// }

		fmt.Println(k)

		for _, x := range v {
			fmt.Println(x)
		}
	}
}
