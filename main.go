package main

import (
	"log"
	"net/http"
	"strings"
)

var ()

func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", neuter(fs)))
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
