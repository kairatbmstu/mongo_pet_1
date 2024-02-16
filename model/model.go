package model

import "time"

type Link struct {
	Id     string
	Url    string
	Status string
}

const (
	PageInitial    = "initial"
	PageProcessing = "processing"
	PageDownloaded = "downloaded"
	PageToUpdate   = "to-update"
)

type PageStatus string

type Page struct {
	Id       string
	Url      string
	Created  time.Time
	Updated  time.Time
	Status   PageStatus
	HTML     string
	Links    []Link
	PageRank float64
}

func PageToLink(page Page) Link {
	var link = Link{}

	return link
}

func LinkToPage(link Link) Page {
	var page = Page{}
	return page
}
