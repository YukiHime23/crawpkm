package service

import (
	"fmt"
	"goCraw/model"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type BookVnService interface {
	CrawBookVn() any
}

func (a AppService) CrawBookVn() any {

	// Tạo một Collector mới
	c := colly.NewCollector()

	// Biến lưu trữ thông tin sách
	var books []model.ISBNBook

	c.OnHTML("#list_data_return tbody tr", func(e *colly.HTMLElement) {
		var book model.ISBNBook

		isbn := strings.TrimSpace(e.ChildText("td:nth-child(2)"))
		if isbn == "" {
			return
		}
		book.ISBN = isbn
		book.BookTitle = strings.TrimSpace(e.ChildText("td:nth-child(3)"))
		book.Author = strings.TrimSpace(e.ChildText("td:nth-child(4)"))
		book.Editor = strings.TrimSpace(e.ChildText("td:nth-child(5)"))
		book.Publisher = strings.TrimSpace(e.ChildText("td:nth-child(6)"))
		book.Partner = strings.TrimSpace(e.ChildText("td:nth-child(7)"))
		book.PlaceOfPrinting = strings.TrimSpace(e.ChildText("td:nth-child(8)"))
		submissionDateLC, err := parseDateStringDDMMYYYY(strings.TrimSpace(e.ChildText("td:nth-child(9)")))
		if err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
		book.SubmissionDateLC = submissionDateLC

		books = append(books, book)
	})

	// Biến lưu trữ số lượng trang
	var totalPages int

	c.OnHTML(".pagination li:not(.disabled)", func(e *colly.HTMLElement) {
		totalPages = e.DOM.Length()
		fmt.Println(strconv.Atoi(e.DOM.Text()))
	})

	urlVisit := createUrl("22/05/2023", "01/06/2023")
	err := c.Visit(urlVisit)
	if err != nil {
		log.Fatal(err)
	}

	// Điều hướng đến các trang phân trang trừ 2 trang cuối cùng
	for i := 1; i < totalPages-2; i++ {
		link := fmt.Sprintf("%s&p=%v", urlVisit, fmt.Sprint(i))
		err := c.Visit(link)
		if err != nil {
			log.Println(err)
		}
	}

	// Lưu thông tin sách
	a.db.Create(books)
	return books
}

func createUrl(start string, end string) string {
	baseUrl := "https://ppdvn.gov.vn"
	tclc := "web/guest/tra-cuu-luu-chieu"
	startStr := url.QueryEscape(start)
	endStr := url.QueryEscape(end)
	urlStr := fmt.Sprintf("%s/%s?query=&id_nxb=-1&bat_dau=%s&ket_thuc=%s", baseUrl, tclc, startStr, endStr)
	return urlStr
}
