package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Vid 17

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

// NewsAggPage ...
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	// Parse XML
	// var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)

	myLocs := []string{
		//"https://www.washingtonpost.com/news-sitemaps/politics.xml",
		//"https://www.washingtonpost.com/news-sitemaps/opinions.xml",
		//"https://www.washingtonpost.com/news-sitemaps/local.xml",
		//"https://www.washingtonpost.com/news-sitemaps/sports.xml",
		//"https://www.washingtonpost.com/news-sitemaps/national.xml",
		//"https://www.washingtonpost.com/news-sitemaps/world.xml",
		//"https://www.washingtonpost.com/news-sitemaps/business.xml",
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
	// Build the page
	p := NewsAggPage{Title: "Amazing News Agg Page", News: newsMap}
	t, err := template.ParseFiles("newsaggtemplate.html")

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// Execute the page
	t.Execute(w, p)

}

func main() {

	// Vid 17
	http.HandleFunc("/agg/", newsAggHandler)

	// All function handlers should come before this
	// Direct WS to listen on Port, nil handler, DefaultServeMux used
	http.ListenAndServe(":8080", nil)

}
