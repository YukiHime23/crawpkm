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

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains(DomainPkm),
)

func readHTML() map[string][]string {
	var title string
	mapChap := make(map[string][]string)

	collectorPKM.OnHTML("a#Blog1_blog-pager-newer-link", func(element *colly.HTMLElement) {
		nextVolume := element.Attr("href")
		fmt.Println(nextVolume)
		collectorPKM.Visit(element.Request.AbsoluteURL(nextVolume))
	})

	collectorPKM.OnHTML("h3.post-title", func(element *colly.HTMLElement) {
		title = strings.Trim(element.Text, "\n")
	})

	collectorPKM.OnHTML(".separator a", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		mapChap[title] = append(mapChap[title], link)
	})

	// collectorPKM.OnHTML(".post-body div a", func(element *colly.HTMLElement) {
	// 	link := element.Attr("href")
	// 	pages = append(pages, link)
	// 	mapChap[title] = pages
	// })

	if err := collectorPKM.Visit(LinkPkmCraw); err != nil {
		log.Fatal(err)
	}

	// Wait until threads are finished
	collectorPKM.Wait()

	return mapChap
}

func main() {
	mapChap := readHTML()

	for k, v := range mapChap {
		if err := os.MkdirAll(PathPkm+k, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		for _, x := range v {
			if err := crawpkm.DownloadFile(x, "", PathPkm+k); err != nil {
				log.Fatal(err)
			}
		}
	}
}
