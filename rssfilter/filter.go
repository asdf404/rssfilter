package rssfilter

import (
	"github.com/mmcdole/gofeed"
)

func (feed *Feed) FetchAndFilter() (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	fetched, err := fp.ParseURL(feed.Url)
	if err != nil {
		return nil, err
	}

	items := []*gofeed.Item{}
filtering:
	for _, item := range fetched.Items {
		for _, rule := range feed.Keep {
			if rule.IsMatch(item) {
				items = append(items, item)
				continue filtering
			}
		}

		for _, rule := range feed.Skip {
			if rule.IsMatch(item) {
				continue filtering
			}
		}

		items = append(items, item)
	}

	fetched.Items = items
	return fetched, nil
}

func (rule *rule) IsMatch(item *gofeed.Item) bool {
	authorMatch := true
	if len(rule.Author) > 0 {
		authorMatch = false
		for _, pattern := range rule.Author {
			same, err := match(pattern, item.Author.Name)
			if err != nil {
				panic(err)
			}
			if same {
				authorMatch = true
				break
			}
		}
	}

	titleMatch := true
	if len(rule.Title) > 0 {
		titleMatch = false
		for _, pattern := range rule.Title {
			same, err := match(pattern, item.Title)
			if err != nil {
				panic(err)
			}
			if same {
				authorMatch = true
				break
			}
		}
	}

	categoriesMatch := true
	if len(rule.Category) > 0 {
		for _, category := range item.Categories {
			categoryMatch := false
			for _, pattern := range rule.Category {
				same, err := match(pattern, category)
				if err != nil {
					panic(err)
				}
				if same {
					categoryMatch = true
					break
				}
			}

			if categoryMatch == false {
				categoriesMatch = false
			}
		}
	}

	return authorMatch && titleMatch && categoriesMatch
}
