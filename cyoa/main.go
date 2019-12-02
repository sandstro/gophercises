package main

import (
	"gophercises/cyoa/stories"
	"io"
	"log"
	"net/http"
)

func main() {
	data := stories.ParseStoriesFromFile()

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
