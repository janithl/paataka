package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Publication is a struct that holds info about websites and blogs
type Publication struct {
	id         int
	Name       string
	URL        *url.URL
	accessedAt time.Time
	updatedAt  time.Time
	Posts      []Post
}

// SetLink sets the publication's URL
func (p *Publication) SetLink(link string) {
	url, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	p.URL = url
}

// AddPost adds a new post to a publication
func (p *Publication) AddPost(post Post) {
	p.Posts = append(p.Posts, post)
	p.updatedAt = time.Now()
}

// Fetch publication's latest posts
func (p *Publication) Fetch() {
	resp, err := http.Get(p.URL.String())
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	type Feed struct {
		XMLName xml.Name `xml:"rss"`
		Name    string   `xml:"channel>title"`
		Posts   []Post   `xml:"channel>item"`
	}

	feed := Feed{}
	if err := xml.Unmarshal([]byte(body), &feed); err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, post := range feed.Posts {
		p.AddPost(post)
	}

	p.accessedAt = time.Now()
	fmt.Println(p)
}
