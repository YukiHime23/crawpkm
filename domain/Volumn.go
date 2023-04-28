package domain

type Volume struct {
	ChapLink []Chapter
}

type Chapter struct {
	Title    string
	PageLink []string
}
