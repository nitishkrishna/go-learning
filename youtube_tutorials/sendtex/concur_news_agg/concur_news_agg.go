package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
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

var newsAggWaitGroup sync.WaitGroup

// Go routine to pull the news objects
func newsRoutine(channelObj chan News, Location string) {
	defer newsAggWaitGroup.Done()
	var newsObj News
	resp, _ := http.Get(strings.TrimSpace(Location))
	bytes, _ := ioutil.ReadAll(resp.Body)
	// Create a news Obj from response Data
	xml.Unmarshal(bytes, &newsObj)
	resp.Body.Close()

	// Fill the News Objects into the Channel
	channelObj <- newsObj
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	// Parse XML
	var siteMapIndexObj SitemapIndex

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &siteMapIndexObj)
	resp.Body.Close()

	newsMap := make(map[string]NewsMap)

	// Channel to push News Objects into
	queue := make(chan News, 30)

	// Call the Go Routines to concurrently pull Info from each XML
	for _, Location := range siteMapIndexObj.Locations {
		newsAggWaitGroup.Add(1)
		go newsRoutine(queue, Location)
	}

	// Wait for channel Buffer to fill, then close it
	newsAggWaitGroup.Wait()
	close(queue)

	// Iterate the Channel to get news Type Objects
	for elem := range queue {
		// Each News Object has a number of Keywords
		for idx := range elem.Keywords {
			newsMap[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
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
