package rssfilter

import (
	"fmt"
	"net/http"
	"github.com/mmcdole/gofeed"
	"github.com/gorilla/feeds"
)

type Server struct {
	Config *Config
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	var feed *Feed
	for _, f := range server.Config.Feeds {
		if f.Path == path {
			feed = &f
			break
		}
	}
	if feed == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fetched, err := feed.FetchAndFilter()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rss, err := buildResponse(fetched)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	fmt.Fprintf(w, "%s", rss)
}

func buildResponse(feed *gofeed.Feed) (string, error) {
	rss := &feeds.Feed{
		Title: feed.Title,
		Link: &feeds.Link{Href: feed.Link},
		Description: feed.Description,
		Author: &feeds.Author{Name: feed.Author.Name, Email: feed.Author.Email},
		Created: *feed.PublishedParsed,
		Items: []*feeds.Item{},
	}

	for _, item := range feed.Items {
		rss.Items = append(rss.Items, convertItem(item))
	}

	xml, err := rss.ToRss()
	if err != nil {
		return "", err
	}
	return xml, nil
}

func convertItem(item *gofeed.Item) *feeds.Item {
	return &feeds.Item{
		Title: item.Title,
		Link: &feeds.Link{Href: item.Link},
		Description: item.Description,
		Created: *item.PublishedParsed,
		Author: &feeds.Author{Name: item.Author.Name, Email: item.Author.Email},
	}
}

// func convertEnclosures(enclosures []*gofeed.Enclosure) []*feeds.Enclosure {
//
// }
