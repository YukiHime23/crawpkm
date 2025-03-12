package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/YukiHime23/crawpkm"
	"github.com/gocolly/colly"
)

var (
	PathPkm     = "Pokemon/"
	DomainPkm   = "www.pokemonspecial.com"
	LinkPkmCraw = "https://www.pokemonspecial.com/2013/12/chapter-001.html"
	BaseURL     = "https://www.pokemonspecial.com/2022/10/pokemon-dac-biet.html"
)

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains(DomainPkm),
)

func crawHTML() {
	var title string

	collectorPKM.OnHTML("a#Blog1_blog-pager-newer-link", func(element *colly.HTMLElement) {
		nextVolume := element.Attr("href")
		fmt.Println(nextVolume)
		//if nextVolume == "https://www.pokemonspecial.com/2014/06/chapter-004.html" {
		//	return
		//}
		collectorPKM.Visit(element.Request.AbsoluteURL(nextVolume))
	})

	collectorPKM.OnHTML("h3.post-title", func(element *colly.HTMLElement) {
		title = strings.Trim(element.Text, "\n")
		if err := os.MkdirAll(PathPkm+title, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	})

	collectorPKM.OnHTML(".separator a", func(element *colly.HTMLElement) {
		//link := element.Attr("href")
		//mapChap[title] = append(mapChap[title], link)
	})

	collectorPKM.OnHTML(".post-body div a", func(element *colly.HTMLElement) {
		fmt.Println(title)

		link := element.Attr("href")
		if err := crawpkm.DownloadFile(link, "", PathPkm+title); err != nil {
			log.Fatal(err)
		}
	})

	if err := collectorPKM.Visit(LinkPkmCraw); err != nil {
		log.Fatal(err)
	}

	// Wait until threads are finished
	collectorPKM.Wait()
}

func main() {
	crawHTML()
}
