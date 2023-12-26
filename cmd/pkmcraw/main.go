package main

import (
	"log"
	"os"
	"strings"

	craw_pkm "github.com/YukiHime23/craw-pkm"
	"github.com/gocolly/colly"
)

type Volume struct {
	ChapLink []Chapter
}

type Chapter struct {
	Title    string
	PageLink []string
}

func main() {
	collectorPKM := colly.NewCollector(
		colly.AllowedDomains(DomainPkm),
	)

	var chap Chapter
	checkEnd := false

	// nextVolume := "https://www.pokemonspecial.com/2023/09/swsh-42.html"
	nextVolume := LinkPkmCraw

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

		CrawChapter(&chap)

		if err := collectorPKM.Visit(nextVolume); err != nil {
			log.Fatal(err)
		}

		if checkEnd {
			break
		}
	}
}

var (
	PathPkm     = "Pokemon/"
	DomainPkm   = "www.pokemonspecial.com"
	LinkPkmCraw = "https://www.pokemonspecial.com/2013/12/chapter-001.html"
	BaseURL     = "https://www.pokemonspecial.com/2022/10/pokemon-dac-biet.html"
)

func CrawChapter(chapter *Chapter) (error, string) {
	if err := os.MkdirAll(PathPkm+chapter.Title, os.ModePerm); err != nil {
		return err, err.Error()
	}

	for _, v := range chapter.PageLink {
		if err := craw_pkm.DownloadFile(v, "", PathPkm+chapter.Title); err != nil {
			return err, err.Error()
		}
	}

	return nil, "craw -> " + chapter.Title + " <- done!"
}
