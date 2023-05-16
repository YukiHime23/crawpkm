package service

import (
	"fmt"
	"github.com/gocolly/colly"
	"goCraw/config"
	"goCraw/model"
	"os"
	"strings"
)

type PokemonSpecialService interface {
	CrawVolume(nextVolume string, volume model.Volume) error
}

var collectorPKM = colly.NewCollector(
	colly.AllowedDomains(config.DomainPkm),
)

func (a AppService) CrawVolume(nextVolume string, vol model.Volume) error {
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
	a.CrawChapter(&c)
	defer a.CrawVolume(nextVolume, vol)

	return nil
}

func (a AppService) CrawChapter(chapter *model.Chapter) (error, string) {
	if err := os.MkdirAll(config.PathPkm+chapter.Title, os.ModePerm); err != nil {
		return err, err.Error()
	}

	for _, v := range chapter.PageLink {
		if err := DownloadFile(v, "", config.PathPkm+chapter.Title); err != nil {
			return err, err.Error()
		}
	}

	return nil, "craw -> " + chapter.Title + " <- done!"
}

func (a AppService) GetPageLink(c *model.Chapter, linkPage string) error {
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
