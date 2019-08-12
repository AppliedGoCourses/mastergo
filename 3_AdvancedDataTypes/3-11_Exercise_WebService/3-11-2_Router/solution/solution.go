package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

// handleQuote handles all requests that begin with "/quote/".
// CRUD (Create, Read, Update, Delete) operations are
// determined by the corresponding HTTP method:
// HTTP POST   = Create (usually, but not always)
// HTTP GET    = Read
// HTTP PUT    = Update (usually, but not always)
// HTTP DELETE = Delete
// Note that POST can also be used for updating resources,
// and PUT can also be used for creating resources.
// Rule of thumb: Use PUT if you know the exact resource
// location already, otherwise use POST.
func handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		io.WriteString(w, "Create: "+r.URL.Path+"\n")
	case "GET":
		io.WriteString(w, "Read: "+r.URL.Path+"\n")
	case "PUT":
		io.WriteString(w, "Update: "+r.URL.Path+"\n")
	case "DELETE":
		io.WriteString(w, "Delete: "+r.URL.Path+"\n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// handleQuotesList delivers a list of all quotes.
func handleQuotesList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
	}
	io.WriteString(w, "List: "+r.URL.Path)
}

func main() {
	// Adding handlers for our REST API. All requests shall start
	// with "/api". A good practice is to bake a version number
	// into the path.
	prefix := "/api/v1/"
	http.HandleFunc(prefix+"quote/", handleQuote)
	http.HandleFunc(prefix+"quotes/", handleQuotesList)
	http.HandleFunc("/", hello)

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
