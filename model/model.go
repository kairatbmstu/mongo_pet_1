package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LinkStatus string

type Link struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Url    string             `bson:"url"`
	Status LinkStatus         `bson:"status"`
}

const (
	PageInitial    = "initial"
	PageProcessing = "processing"
	PageDownloaded = "downloaded"
	PageToUpdate   = "to-update"
)

type PageStatus string

type Page struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Url      string             `bson:"url"`
	Created  time.Time          `bson:"created"`
	Updated  time.Time          `bson:"updated"`
	Status   PageStatus         `bson:"status"`
	HTML     string             `bson:"html"`
	Links    []Link             `bson:"links"`
	PageRank float64            `bson:"pageRank"`
}

func PageToLink(page Page) Link {
	var link = Link{}

	return link
}

func LinkToPage(link Link) Page {
	var page = Page{}
	return page
}

func PagesToLinks(pages []Page) []Link {
	var links []Link

	for _, page := range pages {
		var link = Link{
			ID:     page.ID,
			Url:    page.Url,
			Status: "initial",
		}
		links = append(links, link)
	}

	return links
}
