package main

import (
	"gophercises/cyoa/stories"
	"html/template"
	"log"
	"net/http"
)

func getTemplate() string {
	return `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{ .Title }}</title>
		</head>
		<body>
			<header><h2>{{ .Title }}</h2></header>

			<section>{{ range .Story }}<div>{{ . }}</div>{{ else }}<div><strong>No stories left :(</strong></div>{{ end }}</section>

			<footer>
				<div>{{ range .Options }}<a href="/{{ .Arc }}"><button>{{ .Text }}</button></a>{{ else }}<div>The end!</div>{{ end }}</div>
				<div><a href="/"><button>Back to beginning</button></a></div>
			</footer>
		</body>
	</html>`
}

func main() {
	data := stories.ParseStoriesFromFile()

	t, err := template.New("webpage").Parse(getTemplate())
	if err != nil {
		log.Fatal(err)
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		_ = path
		t.Execute(w, data.Intro)
	}

	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
