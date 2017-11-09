package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/appliedgocourses/quotes"
	"github.com/pkg/errors"
)

// We move the Quote struct to our new quotes package, as the
// quotes package needs to use this struct.

type App struct {
	db quotes.DB // This replaces our map storage
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

func (app *App) createQuote(b []byte) error {
	q := &quotes.Quote{} // Referring to quotes.Quote now
	err := json.Unmarshal(b, q)
	if err != nil {
		return errors.Wrap(err, "createQuote: cannot unmarshal the JSON data")
	}

	err = app.db.Create(q) // Replaced map with DB
	if err != nil {
		return errors.Wrapf(err, "Cannot insert quote '%s'", q.Text)
	}
	return nil
}

func (app *App) getQuote(author string) (*quotes.Quote, error) {
	q, err := app.db.Get(author) // replaced map with DB
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot get quote from %s", author)
	}
	return q, nil
}

func (app *App) processQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot read request body. "+err.Error())
			return
		}

		err = app.createQuote(b)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot unmarshal quote. "+err.Error())
			return
		}

	case "GET":
		author := r.URL.Path[len("/api/v1/quote/"):]
		q, err := app.getQuote(author)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot unmarshal quote. "+err.Error())
			return
		}

		quoteJSON, err := json.MarshalIndent(q, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error: Cannot marshal quote. "+err.Error())
			return
		}
		io.WriteString(w, string(quoteJSON))

	case "PUT":
		io.WriteString(w, "PUT: "+r.URL.Path)
	case "DELETE":
		io.WriteString(w, "DELETE: "+r.URL.Path)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (app *App) listQuotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
	}
	quotes, err := app.db.List() // replaced map with DB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error: Cannot list quotes. "+err.Error())
		return
	}
	quotesJSON, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error: Cannot marshal quotes. "+err.Error())
		return
	}
	io.WriteString(w, string(quotesJSON))
}

func main() {
	// Here we initialize the quotes database and pass the
	// db instance to a new App object.

	db, err := quotes.Open("quotesdb")
	if err != nil {
		log.Fatalln("Cannot open quotesdb:", err)
	}

	// Ensure to close the DB even if errors occur
	// later in the code.
	defer db.Close()

	app := &App{db: *db}

	prefix := "/api/v1/"
	http.HandleFunc(prefix+"quote/", app.processQuote)
	http.HandleFunc(prefix+"quotes/", app.listQuotes)
	http.HandleFunc("/", hello)

	err = http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
