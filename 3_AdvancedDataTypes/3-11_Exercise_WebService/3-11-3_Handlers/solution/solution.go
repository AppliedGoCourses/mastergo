package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const (
	api        = "/api/v1/"
	quotepath  = api + "quote/"
	quotespath = api + "quotes/"
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
	// NOTE: A map is not safe for concurrent access!
	// You will learn the basics of safe shared access
	// to resources in the Concurrency lectures.
	// For this exercise, the final solution will be to
	// use a concurrency-safe key-value store
	// in place of the map.
	storage map[string]*Quote
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

// All functions now become methods of the App object.

// createQuote receives a JSON object representing a Quote and stores
// the quote in the app storage.
func (app *App) createQuote(b []byte) error {
	q := &Quote{}
	err := json.Unmarshal(b, q)
	if err != nil {
		return errors.Wrap(err, "createQuote: cannot unmarshal the JSON data")
	}

	// Now we can create the quote in the storage, if it does not exist already.
	_, ok := app.storage[q.Author]
	if !ok {
		app.storage[q.Author] = q
		return nil
	}
	return errors.Errorf("author %s already exists", q.Author)
}

// getQuote receives an author name and returns the corresponding
// quote of the author, or an error if no quote exists for this author.
func (app *App) getQuote(author string) ([]byte, error) {
	q, ok := app.storage[author]
	if !ok {
		return nil, errors.Errorf("Cannot get quote from %s", author)
	}
	quoteJSON, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot create JSON from quote %v", q)
	}
	return quoteJSON, nil
}

// handleQuote is the handler for all operations on a single quote.
func (app *App) handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot read request body. "+err.Error())
			return
		}

		// Here we can call our new method createQuote.
		err = app.createQuote(b)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot unmarshal quote. "+err.Error())
			return
		}

	case "GET":
		// First, we need to fetch the author's name from the request.
		// It is encoded in the URL like:
		// http://localhost:8000/api/v1/quote/Oscar Wilde

		author := r.URL.Path[len(quotepath):]
		quoteJSON, err := app.getQuote(author)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot get quote from author '"+author+"'\n: "+err.Error())
			return
		}

		// Instead of writing plain strings, we return proper JSON data.

		io.WriteString(w, string(quoteJSON))

	case "PUT":
		io.WriteString(w, "PUT: "+r.URL.Path)
	case "DELETE":
		io.WriteString(w, "DELETE: "+r.URL.Path)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (app *App) getQuotesList() ([]byte, error) {
	quotes := []Quote{}
	// Get the entire list of quotes
	for _, q := range app.storage {
		quotes = append(quotes, *q)
	}
	// Turn the list into JSON.
	// MarshalIndent returns ([]byte, error),
	// which we can directly return to the caller.
	return json.MarshalIndent(quotes, "", "  ")
}

// handleQuotesList is the handler for the "get quotes" operation.
func (app *App) handleQuotesList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request: "+r.Method+" not allowed for "+r.URL.Path)
		return
	}

	quotesJSON, err := app.getQuotesList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error: Cannot marshal quotes. "+err.Error())
		return
	}
	io.WriteString(w, string(quotesJSON))
}

func main() {
	app := &App{
		storage: map[string]*Quote{},
	}

	// Since processQuotes and listQuotes have retained their
	// function signature, we can still pass them to HandleFunc()
	// even though they are now methods of an App object.

	http.HandleFunc(quotepath, app.handleQuote)
	http.HandleFunc(quotespath, app.handleQuotesList)
	http.HandleFunc("/", hello) //    ^--- now a method

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
