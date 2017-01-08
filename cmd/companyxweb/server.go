package main

import (
	"flag"
	"fmt"
	"net/http"
	"companyxchallenge"
)

var (
	port int64
)

func init() {
	flag.Int64Var(&port, "port", 7163, "Port to serve on")
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	joke, err := companyxchallenge.GetRandomJoke()

	w.Header().Set("Content-Type", "text/plain")

	if err != nil {
		fmt.Printf("Error handling request: %v\n", err)
		http.Error(w, "Error making joke", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Sent joke %v to %v\n", joke, r.Host)

	fmt.Fprint(w, joke)
}

func main() {
	flag.Parse()

	serveString := fmt.Sprintf(":%v", port)

	fmt.Printf("Serving HTTP on %v\n", serveString)

	http.HandleFunc("/", jokeHandler)
	http.ListenAndServe(serveString, nil)
}
