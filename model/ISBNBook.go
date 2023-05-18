package model

import "time"

type ISBNBook struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	ISBN             string    `gorm:"unique" json:"ISBN"`
	BookTitle        string    `json:"book_title"`
	Author           string    `json:"author"`
	Editor           string    `json:"editor"`
	Publisher        string    `json:"publisher"`
	Partner          string    `json:"partner"`
	PlaceOfPrinting  string    `json:"place_of_printing"`
	SubmissionDateLC time.Time `json:"submission_date_lc"`
}
