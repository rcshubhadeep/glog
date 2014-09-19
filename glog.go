package glog

import (
	_ "fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("static/tmpl/*"))

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "indexPage", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
