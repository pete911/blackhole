package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	pageTemplate *template.Template
)

func init() {

	htmlTemplate := `<!DOCTYPE html><html lang="en">
<head><title>blackhole</title></head>
<body><p>{{ range . }}<a href="{{ . }}">{{ . }}</a><br />{{ end }}</p></body>
</html>`

	tmpl, err := template.New("page").Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("parse template: %v", err)
	}
	pageTemplate = tmpl
}

func main() {

	flags := ParseFlags()
	log.Printf("started blackhole, flags: %+v", flags)
	http.HandleFunc("/", handler(flags))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", flags.Port), nil))
}

func handler(flags Flags) func(w http.ResponseWriter, _ *http.Request) {

	return func(w http.ResponseWriter, _ *http.Request) {
		if err := pageTemplate.Execute(w, GetLinks(flags)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
