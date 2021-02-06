package handlers

import (
	"log"
	"net/http"

	"github.com/fuzzylimes/goma/pkg/templateloader"
)

var templates = templateloader.GetTemplates()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// err := testTemplate.ExecuteTemplate(w, "base.html", vd)
	template, err := templates.GetTemplate("index.html")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = template.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	template, err := templates.GetTemplate("about.html")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = template.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
