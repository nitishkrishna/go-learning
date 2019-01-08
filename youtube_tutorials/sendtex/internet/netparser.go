package main

import (

	// Read HTTP Response
	"fmt"
	"io/ioutil"
	"net/http"

	// Parse XML
	"encoding/xml"
)

// SitemapIndex ...
type SitemapIndex struct {
	// Must capitalize these to export
	Locations []string `xml:"sitemap>loc"`
}

// News ...
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMap Key of the Map is the title ...
type NewsMap struct {
	Keyword  string
	Location string
}

// Overload String to convert
//func (l Location) String() string {
//	return fmt.Sprint(l.Loc)
//}

func main() {

	// Vid 10
	// Fetch a URL

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// stringBody := string(bytes)
	resp.Body.Close()
	// fmt.Println(stringBody)

	// Vid 11
	// Parse XML
	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)

	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)

	myLocs := []string{
		"https://www.washingtonpost.com/news-sitemaps/politics.xml",
		"https://www.washingtonpost.com/news-sitemaps/opinions.xml",
		"https://www.washingtonpost.com/news-sitemaps/local.xml",
		"https://www.washingtonpost.com/news-sitemaps/sports.xml",
		"https://www.washingtonpost.com/news-sitemaps/national.xml",
		"https://www.washingtonpost.com/news-sitemaps/world.xml",
		"https://www.washingtonpost.com/news-sitemaps/business.xml",
		"https://www.washingtonpost.com/news-sitemaps/technology.xml"}

	for _, Location := range myLocs {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		// Put Data into NewsMap
		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	// NewsMap contains all the data we want
	for title, data := range newsMap {
		fmt.Println("\n\n\nTitle: ", title)
		fmt.Println("\nKeywords: ", data.Keyword)
		fmt.Println("\nLocation: ", data.Location)
	}

}
