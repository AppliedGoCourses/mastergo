package main

import (
	"io"
	"log"
	"net/http"
)

// hello receives a ResponseWriter (implements io.Writer) and a
// Request object. The Request object contains various information
// about the HTTP request recieved, such as the URL, the HTTP
// headers, or the request body.
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

func main() {
	// Here we advise the HTTP server to pass all requests that start
	// with "/" (in other words, all requests) to the handler function
	// hello().
	http.HandleFunc("/", hello)

	// Now we can start the server. We want it to listen to port 8000
	// but only to local clients. (If we used ":8000" instead, without
	// "localhost", the server would also listen to requests made from
	// outside our local machine--unless the local firewall blocks
	// this behavior.
	// The second parameter to ListenAndServe is a Handler type and
	// can be ignored, as we have already registered our handler
	// function.
	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
