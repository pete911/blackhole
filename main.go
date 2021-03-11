package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	Version      = "dev"
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

	flags, err := ParseFlags()
	if err != nil {
		failFlags(err)
	}

	log.Printf("started blackhole version: %s flags: %+v", Version, flags)
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

func failFlags(err error) {

	flag.CommandLine.SetOutput(os.Stderr)
	fmt.Fprintln(flag.CommandLine.Output(), err.Error())
	flag.Usage()
	os.Exit(1)
}
