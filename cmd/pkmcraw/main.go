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
	pathPkm   = "Pokemon/"
	DomainPkm = "www.pokemonspecial.com"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2013/12/chapter-001.html"
	LinkPkmCraw = "https://www.pokemonspecial.com/2014/06/chapter-048.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2014/06/chapter-089.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2014/06/chapter-245.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2014/06/chapter-246.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2014/06/chapter-376.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2021/09/chapter-566.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2021/09/chapter-568.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2023/08/swsh-37.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2024/02/sv-02.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2024/12/sv-14.html"
	// LinkPkmCraw = "https://www.pokemonspecial.com/2024/11/sv-15.html"

	BaseURL = "https://www.pokemonspecial.com/2022/10/pokemon-dac-biet.html"
)

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains("www.pokemonspecial.com"),
	colly.Async(true),
	// Attach a debugger to the collector
	colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
)

func crawHTML() {
	var title string
	var i int
	// Rotate two socks5 proxies
	collectorPKM.Limit(&colly.LimitRule{
		DomainGlob:  "*pokemonspecial.*",
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	collectorPKM.OnHTML("a#Blog1_blog-pager-newer-link", func(element *colly.HTMLElement) {
		nextVolume := element.Attr("href")
		if nextVolume == "https://www.pokemonspecial.com/2014/06/chapter-053.html" {
			return
		}
		// error vol 4 chap 48 while craw from chap 41 to 53
		//vol 7 chap 90, vol 20 chap 245, vol 34 chap 376, chap 566, 568, swsh-chap 37->43, Scarlet Violet: Chapter 2, Scarlet Violet - Chapter 14
		//
		i = 0
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
		title = strings.ReplaceAll(title, "?", "~")

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
		link := element.Attr("href")
		i++
		if err := crawpkm.DownloadFile(link, fmt.Sprintf("%d", i), pathPkm+title); err != nil {
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
