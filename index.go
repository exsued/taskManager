package main

import (
	"net/http"
	"text/template"
)

var (
	tmplPath = "./static/"
	tmpl     = template.Must(template.ParseFiles(tmplPath + "index.html"))
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
