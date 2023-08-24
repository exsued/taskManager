package main

import (
	"net/http"
	"text/template"
)

var (
	tmplPath  = "./public/pages/"
	indextmpl = template.Must(template.ParseFiles(tmplPath + "index/index.html"))
)

func Index(w http.ResponseWriter, r *http.Request) {
	indextmpl, _ = template.ParseFiles(tmplPath + "index/index.html")
	indextmpl.ExecuteTemplate(w, "index.html", Tasks)
}
