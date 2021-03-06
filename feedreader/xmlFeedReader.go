// Package feedreader contains implementations of the FeedReader
// interface, and belongs to the infrastructure layer
package feedreader

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/janithl/paataka/entities"
)

type feedItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	GUID    string `xml:"guid"`
	Content string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

type feed struct {
	XMLName xml.Name   `xml:"rss"`
	Items   []feedItem `xml:"channel>item"`
}

// XMLFeedReader reads RSS feeds
type XMLFeedReader struct {
}

// Read is the read function, which reads an RSS feed and returns
// a slice of Post entities
func (x XMLFeedReader) Read(url string) ([]entities.Post, error) {
	posts := []entities.Post{}

	resp, err := http.Get(url)
	if err != nil {
		return posts, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return posts, err
	}

	feed := feed{}
	if err := xml.Unmarshal([]byte(body), &feed); err != nil {
		return posts, err
	}

	for _, post := range feed.Items {
		postEntity := entities.Post{Title: post.Title, URL: post.Link}

		if len(post.Link) == 0 {
			postEntity.URL = post.GUID
		}

		if pubDate, err := time.Parse(time.RFC1123, post.PubDate); err == nil {
			postEntity.CreatedAt = pubDate
		} else if pubDate, err = time.Parse(time.RFC1123Z, post.PubDate); err == nil {
			postEntity.CreatedAt = pubDate
		} else {
			fmt.Println("Warning:", err)
			postEntity.CreatedAt = time.Now()
		}

		posts = append(posts, postEntity)
	}

	return posts, nil
}
