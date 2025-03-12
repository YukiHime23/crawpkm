package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/YukiHime23/crawpkm"
	"github.com/gocolly/colly"
)

var (
	pathPkm     = "Pokemon/"
	DomainPkm   = "www.pokemonspecial.com"
	LinkPkmCraw = "https://www.pokemonspecial.com/2013/12/chapter-001.html"
	BaseURL     = "https://www.pokemonspecial.com/2022/10/pokemon-dac-biet.html"
)

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains("www.pokemonspecial.com"),
	colly.Async(true),
	// Attach a debugger to the collector
	colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
)

func crawHTML() {
	var title string
	// Rotate two socks5 proxies
	collectorPKM.Limit(&colly.LimitRule{
		DomainGlob:  "*pokemonspecial.*",
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	collectorPKM.OnHTML("a#Blog1_blog-pager-newer-link", func(element *colly.HTMLElement) {
		nextVolume := element.Attr("href")
		//if nextVolume == "https://www.pokemonspecial.com/2014/06/chapter-004.html" {
		//	return
		//}
		// error
		// vol 4 chap 41->51
		//
		collectorPKM.Visit(element.Request.AbsoluteURL(nextVolume))
	})

	absDir, err := filepath.Abs(pathPkm)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	collectorPKM.OnHTML("h3.post-title", func(element *colly.HTMLElement) {
		title = strings.Trim(element.Text, "\n")
		title = strings.ReplaceAll(title, ":", " -")

		newpath := filepath.Join(absDir, title)
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
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
		if err := crawpkm.DownloadFile(link, "", pathPkm+title); err != nil {
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
