package main

import (
	"io"
	"log"
	"net/http"
)

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source,omitempty"`
}

// A new type, App, serves for providing objects that are beyond the
// scope of a single request. Especially, we need a single storage instance
// without having to use a global variable.
//
// All functions that need to access the storage are turned into methods
// of this context. This way, they do not need to receive an extra input
// parameter. This is especially important for our handler functions that
// must adhere to a given function signature, but it is useful for other
// functions as well.

type App struct {
	storage map[string]*Quote
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

// NOTE: These two additional functions - createQuote and getQuote -
// help separating data access from request handling.
// Do everything data store related in these functions, and
// let the handler functions only take care of the REST requests.

// createQuote receives a JSON object representing a Quote and stores
// the quote in the app storage.
func (app *App) createQuote(b []byte) error {

	// TODO:
	// * Unmarshal b into a Quote
	// * Check if the quote exists in the storage, return an error if it does
	// * Write the quote to the storage

}

// getQuote receives an author name and returns the corresponding
// quote of the author, or an error if no quote exists for this author.
func (c *App) getQuote(author string) ([]byte, error) {

	// TODO:
	// * Fetch author's quote from the storage
	// * Return an error if the quote does not exist
	// * Marshal the quote into JSON and return the result.
	//
	// Hint: Use
	//    json.MarshalIndent(q, "", "  ")
	// for producing formatted JSON - the test file expects this!

}

// handleQuote is the handler for all operations on a single quote.
func (app *App) handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		// TODO:
		// * Read the complete request body into a byte slice.
		//   (The request body contains only a single quote in JSON format)
		//   Hint: The package ioutil contains a convenience
		//   function for reading everything until EOF.
		// * Pass the byte slice to createQuote()
		//
		// At each step, handle all errors that can occur.
		// Write the HTTP status to the response header,
		// and the error message to the response body.

	case "GET":

		// TODO:
		// * Fetch the author's name from the request.
		//   It is contained in the URL like:
		//   http://localhost:8000/api/v1/quote/Oscar Wilde
		//   (Escape sequences like %20 for space are already removed in r.URL.Path)
		// * Pass the name to getQuote.
		// * Write the result to the response.
		//
		// Again, handle all errors that can occur.

		// NOTE: If you want, implement PUT and DELETE as well.
		// (Not covered in the solution.)

	case "PUT":
		io.WriteString(w, "PUT: "+r.URL.Path)
	case "DELETE":
		io.WriteString(w, "DELETE: "+r.URL.Path)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// handleQuotesList is the handler for the "get quotes" operation.
func (app *App) handleQuotesList(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
	}

	// TODO:
	// * Turn the contents of app.storage map into a slice of Quote
	// * Turn the slice of Quote into JSON - again, using MarshalIndent().
	//
	// As always, check for errors.

}

func main() {
	app := &App{
		storage: map[string]*Quote{},
	}

	// Since processQuotes and listQuotes have retained their
	// function signature, we can still pass them to HandleFunc()
	// even though they are now methods of an App object.

	prefix := "/api/v1/"
	http.HandleFunc(prefix+"quote/", app.handleQuote)
	http.HandleFunc(prefix+"quotes/", app.handleQuotesList)
	http.HandleFunc("/", hello) //    ^--- now a method

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
