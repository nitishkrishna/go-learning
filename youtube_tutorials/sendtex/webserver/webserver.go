package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Args are the Response Writer and the HTTP Request
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa, Nice!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the About Page")
}

// Vid 9
// Add HTML Tags
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Whoa, Nice!</h1>")
	fmt.Fprint(w, "<h3>Nice Go</h3>")
	fmt.Fprint(w, "<p>This is a paragraph</p>")
	// Note Fprintf for formatting
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
	// Multiline print
	fmt.Fprintf(w, `<h1>Hi!</h1>
<h3>This is also Go</h3>
<p>Multiline</p>`)

}

// Vid 16

// NewsAggPage ...
type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// Build the page
	p := NewsAggPage{Title: "Amazing News Agg Page", News: "some news"}
	t, err := template.ParseFiles("basictemplating.html")

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// Execute the page
	t.Execute(w, p)

}

func main() {

	// Vid 5

	// Similar to bottle request, function to handle path
	http.HandleFunc("/", indexHandler)
	// Another url handler
	http.HandleFunc("/about/", aboutHandler)

	// Vid 9
	http.HandleFunc("/html/", htmlHandler)

	// Vid 16
	http.HandleFunc("/agg/", newsAggHandler)

	// All function handlers should come before this
	// Direct WS to listen on Port, nil handler, DefaultServeMux used
	http.ListenAndServe(":8080", nil)

}
