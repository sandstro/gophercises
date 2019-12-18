package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"log"
	"net/http"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file for stories")
	flag.Parse()
	data := cyoa.ParseStoriesFromFile(*file)

	handler := cyoa.TemplateHandler(data)
	fmt.Println("Server running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", handler))
}
