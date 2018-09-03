/*
Package common implements some common structs and utilities that are used throughout the application
*/
package common

import "net/url"
import "time"

// Post is a struct that holds info about blog/site posts
type Post struct {
	id          int
	Title       string   `xml:"title"`
	URL         *url.URL `xml:"link"`
	Content     string   `xml:"description"`
	publishedAt time.Time
	addedAt     time.Time
	updatedAt   time.Time
}
