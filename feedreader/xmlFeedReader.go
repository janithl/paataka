package feedreader

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/janithl/paataka/entities"
)

type feedItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Content string `xml:"description"`
}

type feed struct {
	XMLName xml.Name   `xml:"rss"`
	Items   []feedItem `xml:"channel>item"`
}

// XMLFeedReader reads RSS feeds
type XMLFeedReader struct {
}

// Read is the read function
func (x XMLFeedReader) Read(url string) map[string]entities.Post {
	posts := map[string]entities.Post{}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return posts
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return posts
	}

	feed := feed{}
	if err := xml.Unmarshal([]byte(body), &feed); err != nil {
		fmt.Println("Error: ", err)
		return posts
	}

	for index, post := range feed.Items {
		posts[string(index)] = entities.Post{Title: post.Title, URL: post.Link}
	}

	return posts
}
