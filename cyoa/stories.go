package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
	"unicode/utf8"
)

type Story map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

func ParseStoriesFromFile(filename string) Story {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := Story{}

	json.Unmarshal([]byte(file), &data)

	return data
}

func storyTemplate() string {
	return `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{ .Title }}</title>
		</head>
		<body>
			<header><h2>{{ .Title }}</h2></header>

			<section>{{ range .Paragraphs }}<div>{{ . }}</div>{{ else }}<div><strong>No stories left :(</strong></div>{{ end }}</section>

			<footer>
				<div>{{ range .Options }}<a href="/{{ .Arc }}"><button>{{ .Text }}</button></a>{{ else }}<div>The end!</div>{{ end }}</div>
				<div><a href="/"><button>Back to beginning</button></a></div>
			</footer>
		</body>
	</html>`
}

func TemplateHandler(story Story) http.Handler {
	return handler{story}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("webpage").Parse(storyTemplate()))
	path := r.URL.Path

	data, ok := h.s[trimFirstRune(path)]
	if ok {
		t.Execute(w, data)
	} else {
		t.Execute(w, h.s["intro"])
	}
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
