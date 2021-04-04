package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"
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

	serverWG := &sync.WaitGroup{}
	log.Printf("starting blackhole version: %s flags: %+v", Version, flags)

	serverWG.Add(1)
	serverShutdown := startServer(serverWG, "blackhole", flags.Port, handler(flags))

	serverProfileShutdown := func() {}
	if flags.ProfilePort != 0 {
		serverWG.Add(1)
		// pprof registers its handler with default server mux
		serverProfileShutdown = startServer(serverWG, "pprof", flags.ProfilePort, http.DefaultServeMux)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	serverShutdown()
	serverProfileShutdown()

	serverWG.Wait()
}

func startServer(wg *sync.WaitGroup, name string, port int, handler http.Handler) (shutdown func()) {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			log.Printf("[ERROR] listen and serve: %v", err)
		}
	}()
	log.Printf("started %s server on port %d", name, port)

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("[ERROR] server shutdown: %v", err)
		}
		log.Printf("shutdown %s server", name)
	}
}

func handler(flags Flags) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if err := pageTemplate.Execute(w, GetLinks(flags)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

func failFlags(err error) {

	flag.CommandLine.SetOutput(os.Stderr)
	if _, e := fmt.Fprintln(flag.CommandLine.Output(), err.Error()); e != nil {
		fmt.Printf("cannot print %v: %v", err, e)
	}
	flag.Usage()
	os.Exit(1)
}
