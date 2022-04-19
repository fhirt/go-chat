package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var port = flag.String("port", ":8080", "The port of the application.")
	flag.Parse()
	r:= NewRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// get the room going
	go r.Run()
	//start the web server
	log.Println("Starting web server on", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
