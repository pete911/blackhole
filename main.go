package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const (
	maxNumberOfLinks = 100
	minNumberOfLinks = 5
	maxLinkDepth     = 10
	minLinkDepth     = 1
)

func handler(w http.ResponseWriter, _ *http.Request) {

	if err := writePage(w, getLinks()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getLinks() []string {

	var links []string
	for i := 0; i < random(minNumberOfLinks, maxNumberOfLinks); i++ {
		links = append(links, getPath())
	}
	return links
}

func getPath() string {

	var path []string
	for i := 0; i < random(minLinkDepth, maxLinkDepth); i++ {
		path = append(path, fmt.Sprintf("%d", rand.Intn(10000)))
	}
	return "/" + strings.Join(path, "/")
}

func writePage(w io.Writer, links []string) error {

	tmpl, err := template.New("page").Parse(pageTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, links)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var pageTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<title>test page</title>
</head>
<body>
{{ range . }}
<a href="{{ . }}">{{ . }}</a>
{{ end }}
</p>
</body>
</html>`
