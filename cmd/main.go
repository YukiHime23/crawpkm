package main

import (
	"fmt"
	"os"
	"strings"

	craw_pkm "github.com/YukiHime23/craw-pkm"
	"github.com/YukiHime23/craw-pkm/model"
	"github.com/gocolly/colly"
)

func main() {

}

var (
	PathPkm     = "Pokemon/"
	DomainPkm   = "www.pokemonspecial.com"
	LinkPkmCraw = "https://www.pokemonspecial.com/2013/12/chapter-001.html"
)

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains(DomainPkm),
)

func CrawVolume(nextVolume string, vol model.Volume) error {
	c := model.Chapter{}
	collectorPKM.OnHTML("#Blog1_blog-pager-newer-link", func(element *colly.HTMLElement) {
		nextVolume = element.Attr("href")
	})
	collectorPKM.OnHTML("h3.post-title", func(element *colly.HTMLElement) {
		t := strings.Trim(element.Text, "\n")
		c.Title = t
	})
	//collector.OnHTML(".separator a", func(element *colly.HTMLElement) {
	//	link := element.Attr("href")
	//	c.PageLink = append(c.PageLink, link)
	//})
	//collector.OnHTML(".separator img", func(element *colly.HTMLElement) {
	//	link := element.Attr("src")
	//	c.PageLink = append(c.PageLink, link)
	//})
	collectorPKM.OnHTML(".post-body div a", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		c.PageLink = append(c.PageLink, link)
	})

	if err := collectorPKM.Visit(nextVolume); err != nil {
		return err
	}

	fmt.Println(c.Title)
	vol.ChapLink = append(vol.ChapLink, c)
	CrawChapter(&c)
	defer CrawVolume(nextVolume, vol)

	return nil
}

func CrawChapter(chapter *model.Chapter) (error, string) {
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

func GetPageLink(c *model.Chapter, linkPage string) error {
	collectorPKM.OnHTML("h3.post-title", func(element *colly.HTMLElement) {
		t := strings.Trim(element.Text, "\n")
		c.Title = t
	})
	collectorPKM.OnHTML(".separator a", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		c.PageLink = append(c.PageLink, link)
	})

	if err := collectorPKM.Visit(linkPage); err != nil {
		return err
	}

	return nil
}
