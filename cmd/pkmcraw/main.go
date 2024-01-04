package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/YukiHime23/crawpkm"
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
	checkEnd := false

	nextVolume := "https://www.pokemonspecial.com/2014/06/chapter-021.html"
	// nextVolume := LinkPkmCraw

	for {
		collectorPKM.OnHTML("h3.post-title.entry-title", func(e *colly.HTMLElement) {
			chap.Title = strings.Trim(e.Text, "\n")
		})

		collectorPKM.OnHTML(".separator a", func(element *colly.HTMLElement) {
			link := element.Attr("href")
			chap.PageLink = append(chap.PageLink, link)
		})

		collectorPKM.OnHTML("a#Blog1_blog-pager-newer-link.blog-pager-newer-link", func(element *colly.HTMLElement) {
			nextVolume = element.Attr("href")
			if nextVolume == BaseURL {
				checkEnd = true
			}
		})

		if err := os.MkdirAll(PathPkm+chap.Title, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		for _, v := range chap.PageLink {
			if err := crawpkm.DownloadFile(v, "", PathPkm+chap.Title); err != nil {
				log.Fatal(err)
			}
			fmt.Println("craw -> " + chap.Title + " <- done!")
		}

		if err := collectorPKM.Visit(nextVolume); err != nil {
			log.Fatal(err)
		}

		if checkEnd {
			break
		}
	}
}
