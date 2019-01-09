package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
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
	var siteMapIndexObj SitemapIndex
	var newsObj News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &siteMapIndexObj)
	resp.Body.Close()

	newsMap := make(map[string]NewsMap)

	for _, Location := range siteMapIndexObj.Locations {
		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &newsObj)
		resp.Body.Close()
		// Put Data into NewsMap
		for idx := range newsObj.Titles {
			newsMap[newsObj.Titles[idx]] = NewsMap{newsObj.Keywords[idx], newsObj.Locations[idx]}
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
