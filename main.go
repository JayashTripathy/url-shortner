package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	port := ":8080"
	mux := defaultMux()

	defaultPaths := map[string]string{
		"/portfolio":    "https://jayash.space",
	}


	rawJSON := `[
    {"path": "/pretty-count", "url": "https://prettycount.jayash.space"}
	]`
	mapHandler := jsonHandler([]byte(rawJSON), mux, defaultPaths)

	fmt.Println("Starting the server on :8080")
	servErr := http.ListenAndServe(port, mapHandler)
	if servErr != nil {
		log.Fatalf("server failed to start: %v", servErr)
	}
}
