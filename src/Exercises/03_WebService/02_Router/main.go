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

func handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		io.WriteString(w, "Create: "+r.URL.Path+"\n")
	case "GET":
		io.WriteString(w, "Read: "+r.URL.Path)
	case "PUT":
		io.WriteString(w, "Update: "+r.URL.Path)
	case "DELETE":
		io.WriteString(w, "Delete: "+r.URL.Path)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

/*
TODO: function "handleQuotesList"

* Verify that the method is "GET"
* Return "Bad Request" otherwise
*/

func main() {
	prefix := "/api/v1/"
	http.HandleFunc(prefix+"quote/", handleQuote)
	// TODO: Register handler
	http.HandleFunc("/", hello)

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
