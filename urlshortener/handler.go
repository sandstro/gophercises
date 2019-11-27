package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		redirect, ok := pathsToUrls[path]
		if ok {
			http.Redirect(w, r, redirect, http.StatusSeeOther)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var yamlStruct []yamlFormat
	err := yaml.Unmarshal(yml, &yamlStruct)

	if err != nil {
		return nil, err
	}

	yamlMap := yamlToMap(&yamlStruct)

	return MapHandler(yamlMap, fallback), nil
}

type yamlFormat struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func yamlToMap(yamls *[]yamlFormat) map[string]string {
	mappedYaml := make(map[string]string)
	for _, yaml := range *yamls {
		mappedYaml[yaml.Path] = yaml.URL
	}
	return mappedYaml
}
